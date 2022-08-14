package gpk

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"strings"
)

//GetFiles 返回目录下的所有文件
func (g *Object) GetFiles(folder string) (files []fileInfo, err error) {
	var fs []fs.FileInfo
	if fs, err = ioutil.ReadDir(folder); err != nil {
		return
	}
	for _, fi := range fs {
		fullname := folder + fi.Name()
		if fi.IsDir() {
			var temp []fileInfo
			if temp, err = g.GetFiles(fullname + "/"); err != nil {
				return
			}
			files = append(files, temp...)
		} else {
			files = append(files, fileInfo{
				fullname,
				strings.Replace(fullname, folder, "", 1),
				fi.Size(),
			})
		}
	}
	return
}

//CreateMutiDir 创建多级目录
func (g *Object) CreateMutiDir(fullPath string, isFile ...bool) (err error) {
	var path string
	if len(isFile) > 0 && isFile[0] {
		path = getDir(fullPath)
	} else {
		path = fullPath
	}
	if !isExist(path) {
		if err = os.MkdirAll(path, os.ModePerm); err != nil {
			return
		}
	}
	return nil
}

//Create 新建空文件
func (g *Object) Create(target, key string) (err error) {
	g.Close()
	if err = create(target); err != nil {
		return
	}
	if g.file, err = os.OpenFile(target, os.O_RDWR, 0777); err != nil {
		return
	}
	//移到读写位置
	if _, err = g.file.Seek(0, io.SeekEnd); err != nil {
		return
	}
	g.key = key
	g.list = map[string]splitInfo{}
	return
}

//CreateFromFolder 从文件夹打包
func (g *Object) CreateFromFolder(folder, target string, key string, cfun Process) (err error) {
	if err = g.Create(target, key); err != nil {
		return
	}
	if folder[len(folder)-1:] != "/" {
		folder += "/"
	}
	var files []fileInfo
	if files, err = g.GetFiles(folder); err != nil {
		return
	}
	count := len(files)
	for k, v := range files {
		v.ShortName = strings.Replace(v.FullName, folder, "", 1)
		cfun(v.ShortName, k+1, count)
		//读取文件
		var data []byte
		if data, err = ioutil.ReadFile(v.FullName); err != nil {
			return
		}
		var split splitInfo
		if split, err = writeData(g.file, data); err != nil {
			return
		}
		g.list[v.ShortName] = split
	}
	err = g.Save()
	return
}

//Save 保存操作结果
func (g *Object) Save() (err error) {
	err = writeHead(g.file, g.list, g.key)
	return
}

//Close 销毁
func (g *Object) Close() {
	if g.file != nil {
		g.file.Close()
	}
	g.key = ""
	g.list = make(map[string]splitInfo)
}

//Load 载入文件
func (g *Object) Load(file, key string) (err error) {
	g.Close()
	if g.file, err = os.OpenFile(file, os.O_RDWR, 0777); err != nil {
		return
	}
	//读取标识
	tagData := make([]byte, 3)
	if _, err = g.file.Read(tagData); err != nil {
		return
	}
	if string(tagData) != fileTag {
		err = fmt.Errorf("file type is incorrect")
		return
	}
	//读取版本
	verData := make([]byte, 1)
	if _, err = g.file.Read(verData); err != nil {
		return
	}
	var verInt int8
	if verInt, err = bytesToInt8(verData); err != nil {
		return
	}
	if verInt != version {
		err = fmt.Errorf("file version is incorrect")
		return
	}
	//文件头大小
	lenData := make([]byte, 4)
	if _, err = g.file.Read(lenData); err != nil {
		return
	}
	var lenInt int32
	if lenInt, err = bytesToInt32(lenData); err != nil {
		return
	}
	lenInt ^= int32(encNum)
	//文件偏移
	offsetData := make([]byte, 8)
	if _, err = g.file.Read(offsetData); err != nil {
		return
	}
	var offsetInt int64
	if offsetInt, err = bytesToInt64(offsetData); err != nil {
		return
	}
	offsetInt ^= int64(encNum)
	//读入文件头
	if _, err = g.file.Seek(offsetInt, io.SeekStart); err != nil {
		return
	}
	headData := make([]byte, lenInt)
	if _, err = g.file.Read(headData); err != nil {
		return
	}
	if headData, err = aesDecryptCFB(headData, key); err != nil {
		return
	}
	if headData, err = zlibUncompress(headData); err != nil {
		return
	}
	g.key = key
	err = json.Unmarshal(headData, &g.list)
	//fmt.Println(g.list)
	return
}

//GetList 返回文件列表
func (g *Object) GetList() (list []string) {
	for k := range g.list {
		list = append(list, k)
	}
	return
}

//GetList 返回文件数量
func (g *Object) GetTotal() int {
	return len(g.list)
}

//Load 读取文件
func (g *Object) GetData(name string) (data []byte, err error) {
	split := g.list[name]
	if split.Offset == 0 {
		err = fmt.Errorf("file does not exist")
		return
	}
	if _, err = g.file.Seek(split.Offset, io.SeekStart); err != nil {
		return
	}
	data = make([]byte, split.Len)
	if _, err = g.file.Read(data); err != nil {
		return
	}
	var md5 string
	if md5, err = getMd5(data); err != nil {
		return
	}
	if md5 != split.Md5 {
		err = fmt.Errorf("data verification fails")
		return
	}
	//解压
	if data, err = zlibUncompress(data); err != nil {
		return
	}
	return
}

//Delete 逻辑删除文件
func (g *Object) Delete(name string) {
	delete(g.list, name)
}

//AddOrUpdate 逻辑添加或更新文件
func (g *Object) AddOrUpdate(name string, data []byte) (err error) {
	//移到文件尾部
	var offset int64
	if offset, err = g.file.Seek(0, io.SeekEnd); err != nil {
		return
	}
	//写出数据
	var split splitInfo
	if split, err = writeData(g.file, data, offset); err != nil {
		return
	}
	g.list[name] = split
	return
}

//NewFile 存储为新的文件(同时清理逻辑操作的残留文件)
func (g *Object) NewFile(target string, cfun Process) (err error) {
	if err = create(target); err != nil {
		return
	}
	var newFile *os.File
	if newFile, err = os.OpenFile(target, os.O_RDWR, 0777); err != nil {
		return
	}
	defer newFile.Close()
	if _, err = newFile.Seek(0, io.SeekEnd); err != nil {
		return
	}
	headList := make(map[string]splitInfo)
	index := 0
	count := len(g.list)
	for k, v := range g.list {
		index++
		cfun(target, index, count)
		if _, err = g.file.Seek(v.Offset, io.SeekStart); err != nil {
			return
		}
		data := make([]byte, v.Len)
		if _, err = g.file.Read(data); err != nil {
			return
		}
		//写出数据
		var position int64
		if position, err = newFile.Seek(0, io.SeekCurrent); err != nil {
			return
		}
		if _, err = newFile.Write(data); err != nil {
			return
		}
		split := splitInfo{
			Offset: position,
			Len:    v.Len,
			Md5:    v.Md5,
		}
		headList[k] = split
	}
	err = writeHead(newFile, headList, g.key)
	return
}

//Merge 合并自另外一个文件
func (g *Object) Merge(target, key string, cfun Process) (err error) {
	//打开文件
	var src Object
	if err = src.Load(target, key); err != nil {
		return
	}
	defer src.Close()
	for k, v := range src.GetList() {
		cfun(v, k+1, src.GetTotal())
		var data []byte
		if data, err = src.GetData(v); err != nil {
			return
		}
		if err = g.AddOrUpdate(v, data); err != nil {
			return
		}
	}
	err = g.Save()
	return
}
