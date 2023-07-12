package mns

import (
	"encoding/base64"
)

func Base64Decode(content string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(content)
	return string(decoded), err
}
