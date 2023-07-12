/**
 * @Author: chenwei
 * @Description:
 * @Date: 2022/11/04 15:46
 */

package dugindb

import (
	"reflect"
	"unsafe"
)

// GetSliceLMR: 获取两个切片交集的左部分，重合部分，以及右部分
func GetSliceLMR(sliceL, sliceR []string) ([]string, []string, []string) {
	if len(sliceL) == 0 && len(sliceR) == 0 {
		return nil, nil, nil
	}
	if len(sliceL) == 1 && len(sliceR) == 1 {
		if sliceL[0] == sliceR[0] {
			return nil, sliceL, nil
		} else {
			return sliceL, nil, sliceR
		}
	}
	tm := make(map[string]int)
	for _, e := range sliceL {
		if _, ok := tm[e]; ok {
			tm[e] += 1
		} else {
			tm[e] = 1
		}
	}
	for _, e := range sliceR {
		if _, ok := tm[e]; ok {
			tm[e] += 2
		} else {
			tm[e] = 2
		}
	}
	var l, m, r []string
	for k, v := range tm {
		switch v {
		case 1:
			l = append(l, k)
		case 2:
			r = append(r, k)
		case 3:
			m = append(m, k)
		}
	}
	return l, m, r
}
func String2Bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}
func Strings2Bytes(s []string) [][]byte {
	bss := make([][]byte, len(s))
	for i, e := range s {
		bss[i] = String2Bytes(e)
	}
	return bss
}

func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
func SliceContains(s []string, key string) bool {
	for _, e := range s {
		if e == key {
			return true
		}
	}
	return false
}
