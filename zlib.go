package gpk

import (
	"bytes"
	"compress/zlib"
	"io"
)

func zlibCompressor(data []byte) (res []byte, err error) {
	var in bytes.Buffer
	var w *zlib.Writer
	if w, err = zlib.NewWriterLevel(&in, zlib.BestCompression); err != nil {
		return
	}
	if _, err = w.Write(data); err != nil {
		w.Close()
		return
	} else {
		w.Close()
		//去掉 78 DA 防止遍历压缩包
		res = in.Bytes()[2:]
	}
	return
}

func zlibUncompress(data []byte) (res []byte, err error) {
	//加上 78 DA
	data = bytesCombine([]byte{120, 218}, data)
	var out bytes.Buffer
	in := bytes.NewBuffer(data)
	var r io.ReadCloser
	if r, err = zlib.NewReader(in); err != nil {
		return
	}
	if _, err = io.Copy(&out, r); err != nil {
		r.Close()
		return
	} else {
		r.Close()
		res = out.Bytes()
	}
	return
}
