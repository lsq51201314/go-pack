package gpk

import "os"

const fileTag string = "GPK"
const version int8 = 1
const encNum int = 703093033 //位异或

type Object struct {
	file *os.File
	list map[string]splitInfo
	key  string
}

type fileInfo struct {
	FullName  string `json:"full_name"`
	ShortName string `json:"short_name"`
	Size      int64  `json:"size"`
}

type splitInfo struct {
	Offset int64  `json:"offset"`
	Len    int32  `json:"len"`
	Md5    string `json:"md5"`
}

type Process func(file string, current, count int)
