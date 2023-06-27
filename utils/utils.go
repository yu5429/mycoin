package utils

import (
	"bytes"
	"encoding/gob"
	"log"
)

func HandleErr(err error) { //에러 예외처리
	if err != nil {
		log.Panic(err)
	}
}

func ToBytes(i interface{}) []byte { // byte 변환 함수
	var aBuffer bytes.Buffer
	encoder := gob.NewEncoder(&aBuffer)
	HandleErr(encoder.Encode(i))
	return aBuffer.Bytes()
}

func FromBytes(i interface{}, data []byte) { //다시 byte 변환
	encoder := gob.NewDecoder(bytes.NewReader(data))
	HandleErr(encoder.Decode(i))
}
