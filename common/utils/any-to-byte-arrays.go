package utils

import (
	"bytes"
	"encoding/binary"
)

func GetByteArray(num interface{}) ([]byte, error) {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}
