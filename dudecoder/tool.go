package dudecoder

import (
	"log"
	"reflect"
	"unsafe"
)

func String2Bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func interfaceToBytes(v interface{}) []byte {
	if v == nil {
		return nil
	}
	switch v.(type) {
	case string:
		return String2Bytes(v.(string))
	case []byte:
		return v.([]byte)
	default:
		log.Println("interfaceToBytes unhandled type:", reflect.TypeOf(v))
		return nil
	}
}
