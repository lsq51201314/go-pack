package gpk

import (
	"encoding/json"
	"io"
	"os"
)

func writeHead(file *os.File, data map[string]splitInfo, key string, offset ...int64) (err error) {
	//构建头部
	var headData []byte
	if headData, err = json.Marshal(data); err != nil {
		return
	}
	if headData, err = zlibCompressor(headData); err != nil {
		return
	}
	if headData, err = aesEncryptCFB(headData, key); err != nil {
		return
	}
	//文件头偏移
	var position int64
	if len(offset) > 0 {
		position = offset[0]
		if _, err = file.Seek(position, io.SeekStart); err != nil {
			return
		}
	} else {
		if position, err = file.Seek(0, io.SeekEnd); err != nil {
			return
		}
	}
	//写出文件头
	if _, err = file.Write(headData); err != nil {
		return
	}
	var headSize, headOffset []byte
	//写出大小
	if headSize, err = int32ToBytes(int32(len(headData)) ^ int32(encNum)); err != nil {
		return
	}
	if _, err = file.WriteAt(headSize, 4); err != nil {
		return
	}
	//写出偏移
	if headOffset, err = int64ToBytes(position ^ int64(encNum)); err != nil {
		return
	}
	if _, err = file.WriteAt(headOffset, 8); err != nil {
		return
	}
	return
}
