package gpk

import "os"

//Create 新建空文件
func create(target string) (err error) {
	var file *os.File
	if file, err = os.Create(target); err != nil {
		return
	}
	//文件标识
	if _, err = file.Write([]byte(fileTag)); err != nil {
		return
	}
	var ver, size, position []byte
	//文件版本
	if ver, err = int8ToBytes(version); err != nil {
		return
	}
	if _, err = file.Write(ver); err != nil {
		return
	}
	//文件头大小
	if size, err = int32ToBytes(int32(0 ^ encNum)); err != nil {
		return
	}
	if _, err = file.Write(size); err != nil {
		return
	}
	//文件头偏移
	if position, err = int64ToBytes(int64(16 ^ encNum)); err != nil {
		return
	}
	if _, err = file.Write(position); err != nil {
		return
	}
	return
}
