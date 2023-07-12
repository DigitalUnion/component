package tools

import (
	"git.du.com/cloud/du_component/ducachecenter/consts"
)

func JointKV(key []byte, value []byte) []byte {
	tmp := make([]byte, 0)
	tmp = append(tmp, key...)
	tmp = append(tmp, consts.Separator)
	tmp = append(tmp, value...)
	return tmp
}
