package gpk

import (
	"bytes"
	"encoding/binary"
)

func int8ToBytes(n int8) (res []byte, err error) {
	bytesBuffer := bytes.NewBuffer([]byte{})
	if err = binary.Write(bytesBuffer, binary.LittleEndian, n); err != nil {
		return
	}
	res = bytesBuffer.Bytes()
	return
}

func bytesToInt8(data []byte) (res int8, err error) {
	bytesBuffer := bytes.NewBuffer(data)
	if err = binary.Read(bytesBuffer, binary.LittleEndian, &res); err != nil {
		return
	}
	return
}

func int32ToBytes(n int32) (res []byte, err error) {
	bytesBuffer := bytes.NewBuffer([]byte{})
	if err = binary.Write(bytesBuffer, binary.LittleEndian, n); err != nil {
		return
	}
	res = bytesBuffer.Bytes()
	return
}

func bytesToInt32(data []byte) (res int32, err error) {
	bytesBuffer := bytes.NewBuffer(data)
	if err = binary.Read(bytesBuffer, binary.LittleEndian, &res); err != nil {
		return
	}
	return
}

func int64ToBytes(n int64) (res []byte, err error) {
	bytesBuffer := bytes.NewBuffer([]byte{})
	if err = binary.Write(bytesBuffer, binary.LittleEndian, n); err != nil {
		return
	}
	res = bytesBuffer.Bytes()
	return
}

func bytesToInt64(data []byte) (res int64, err error) {
	bytesBuffer := bytes.NewBuffer(data)
	if err = binary.Read(bytesBuffer, binary.LittleEndian, &res); err != nil {
		return
	}
	return
}

func bytesCombine(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}
