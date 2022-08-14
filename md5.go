package gpk

import (
	"crypto/md5"
	"encoding/hex"
)

func getMd5(data []byte) (res string, err error) {
	md5 := md5.New()
	if _, err = md5.Write(data); err != nil {
		return
	}
	res = hex.EncodeToString(md5.Sum([]byte("")))
	return
}
