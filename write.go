package gpk

import (
	"io"
	"os"
)

func writeData(file *os.File, data []byte, offset ...int64) (split splitInfo, err error) {
	//压缩数据
	if data, err = zlibCompressor(data); err != nil {
		return
	}
	//位置偏移
	var position int64
	if len(offset) > 0 {
		position = offset[0]
		if _, err = file.Seek(position, io.SeekStart); err != nil {
			return
		}
	} else {
		if position, err = file.Seek(0, io.SeekCurrent); err != nil {
			return
		}
	}
	//信息内容
	var md5 string
	if md5, err = getMd5(data); err != nil {
		return
	}
	split = splitInfo{
		Offset: position,
		Len:    int32(len(data)),
		Md5:    md5,
	}
	//写出数据
	if _, err = file.Write(data); err != nil {
		return
	}
	return
}
