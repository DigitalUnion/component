package dudecoder

import (
	"crypto/md5"
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//---------------------------------------
//Anonymous algorithm
//---------------------------------------
var xorKey = []byte{0x1f, 0x6f, 0x8b, 0xa3}

//xorEncrypt 进行异或加密
func xorEncrypt(raw, key []byte) []byte {
	kl := len(key)
	enc := make([]byte, len(raw))
	for i := 0; i < len(raw); i++ {
		enc[i] = raw[i] ^ key[i%kl]
	}
	return enc
}

//xorDecrypt 进行异或解密
func xorDecrypt(raw, key []byte) []byte {
	return xorEncrypt(raw, key)
}

//offsetAndroid 安卓 偏移量算法：
func offsetAndroid(timeKey int) int {
	return (timeKey%100)%57 + 1
}

func offsetIOS(timeKey int) int {
	return (timeKey%100)%50 + 1
}

//offsetMethodIdx  解码解密时，获取乱序方法索引数 mtdIdx 用于选择乱序排序方法
func offsetMethodIdx(timeKey int) int {
	return (timeKey / 100) % 4
}

//decrypt0 解密方法0
//ODD: 35241 -> 3 52 41 -> 3 54 21 ->3 21 54 ->21 3 54 -> 12 3 45
//EVEN: 9786 -> 97 86 -> 98 76 -> 76 98-> 67 89
func decrypt0(msg []byte) ([]byte, error) {
	size := len(msg)
	if size == 0 {
		return nil, errors.New("Decrypt_0 empty message")
	}
	if size == 1 {
		return msg, nil
	}
	var oddChar *byte
	
	rSize := size
	sSize := size / 2
	rMsg := msg
	//(1) 判断奇偶,如果是奇数 把首位拿出来
	if size%2 == 1 {
		oddChar = &msg[0]
		rSize = size - 1
		rMsg = msg[1:]
	}
	//(2)插花逆操作 把字段分为两个部分part1,part2
	part1 := make([]byte, sSize)
	part2 := make([]byte, sSize)
	for m, n := 0, rSize-1; n >= 0; m, n = m+1, n-2 {
		part1[m] = rMsg[n]
		part2[m] = rMsg[n-1]
	}
	target := make([]byte, size)
	if oddChar != nil {
		copy(target, part1) // part1在前 正序-->
		target[sSize] = *oddChar
		copy(target[sSize+1:], part2)
	} else {
		copy(target, part1)
		copy(target[sSize:], part2)
	}
	return target, nil
}

//decrypt1 解密方法1
//ODD: 51423 -> 51 42 3 -> 54 12 3 -> 12 54 3 -> 12 3 54 -> 12 3 45
//EVEN: 9687 -> 96 87 -> 98 67 -> 67 98-> 67 89

func decrypt1(msg []byte) ([]byte, error) {
	size := len(msg)
	if size == 0 {
		return nil, errors.New("Decrypt_1 empty message")
	}
	if size == 1 {
		return msg, nil
	}
	var oddChar *byte
	
	rSize := size
	sSize := size / 2
	rMsg := msg
	//(1) 判断奇偶,如果是奇数 把首位拿出来
	if size%2 == 1 {
		oddChar = &msg[size-1]
		rSize = size - 1
		rMsg = msg[0 : size-1]
	}
	//(2)插花逆操作 把字段分为两个部分part1,part2
	part1 := make([]byte, sSize)
	part2 := make([]byte, sSize)
	for m, n := 0, 0; n < rSize; m, n = m+1, n+2 {
		part1[m] = rMsg[n]
		part2[m] = rMsg[n+1]
	}
	//(3)第一部分字符串取反
	for i, j := 0, sSize-1; i < j; i, j = i+1, j-1 {
		part1[i], part1[j] = part1[j], part1[i]
	}
	target := make([]byte, size)
	if oddChar != nil {
		copy(target, part2)
		target[sSize] = *oddChar
		copy(target[sSize+1:], part1)
	} else {
		copy(target, part2)
		copy(target[sSize:], part1)
	}
	return target, nil
}

//decrypt2 解密方法2
//ODD: 53142 -> 531 42 -> 54 32 1 > 12345
//EVEN: 9786 -> 97 86 -> 98 76-> 6789
func decrypt2(msg []byte) ([]byte, error) {
	size := len(msg)
	if size == 0 {
		return nil, errors.New("decrypt_2 empty msg")
	}
	if size == 1 {
		return msg, nil
	}
	remainder := size % 2
	
	sSize := size / 2
	//var oddChar *byte
	//part1 := make([]byte, sSize+remainder)
	//part2 := make([]byte, sSize)
	part1 := msg[0 : sSize+remainder]
	part2 := msg[sSize+remainder:]
	target := make([]byte, 0)
	var i int
	for i = 0; i < sSize; i++ {
		target = append(target, part1[i], part2[i])
	}
	if size%2 == 1 {
		target = append(target, part1[i])
	}
	var res []byte
	for i := size - 1; i >= 0; i-- {
		res = append(res, target[i])
	}
	return res, nil
}

//decrypt3 解密方法3
//ODD:  13524 -> 135 24 -> 12345
//EVEN: 6879 -> 68 79 -> 6789
func decrypt3(msg []byte) ([]byte, error) {
	size := len(msg)
	if size == 0 {
		return nil, errors.New("decrypt_2 empty msg")
	}
	if size == 1 {
		return msg, nil
	}
	remainder := size % 2
	
	sSize := size / 2
	part1 := msg[0 : sSize+remainder]
	part2 := msg[sSize+remainder:]
	target := make([]byte, size)
	
	var i, j int
	for i, j = 0, 0; i < sSize; i, j = i+1, j+2 {
		target[j] = part1[i]
		target[j+1] = part2[i]
	}
	if size%2 == 1 {
		target[size-1] = part1[sSize+remainder-1]
	}
	return target, nil
}

func timeKey(curtime string) int {
	if len(curtime) < 3 {
		return 0
	}
	//123
	timeKey := make([]byte, 3)
	timeByte := []byte(curtime)
	i := 2
	for tlen := len(timeByte) - 1; tlen > 0; tlen-- {
		if timeByte[tlen] >= '0' && timeByte[tlen] <= '9' {
			timeKey[i] = timeByte[tlen]
			i--
		}
		if i < 0 {
			break
		}
	}
	val, _ := strconv.Atoi(string(timeKey))
	return val
}

func offsetDelete(timeKey int) int {
	return timeKey % 10
}

func offsetSum(timeKey int) int {
	return (timeKey / 10) % 10
}

func offsetDecode(msg []byte, offset int) []byte {
	res := make([]byte, len(msg))
	for i := 0; i < len(msg); i++ {
		t := int8(msg[i] - byte(offset))
		if t < int8(32) {
			t = int8(t - 31 + 126)
		}
		res[i] = byte(t)
	}
	return res
}

//insert
func insertByteAtEnd(msg []byte, char byte, offset_end int) ([]byte, bool) {
	size := len(msg)
	if size == 0 {
		return nil, false
	}
	if offset_end > size {
		return nil, false
	}
	t := make([]byte, size+1)
	offset := size - offset_end
	copy(t, msg[0:offset])
	t[offset] = char
	copy(t[offset+1:], msg[offset:])
	return t, true
}

func asciiSum(msg []byte) int {
	sum := 0
	for i := 0; i < len(msg); i++ {
		sum += int(msg[i])
	}
	return sum
}

const (
	asciiBegin = 32
	asciiEnd   = 126
	md5Size    = 2 << 14
)

var md5SumCache map[string]int

func findMd5(md5Sum string) (int, bool) {
	if val, ok := md5SumCache[md5Sum]; ok {
		return val, ok
	}
	return 0, false
}

func containNonAscii(msg []byte) bool {
	for i := 0; i < len(msg); i++ {
		if msg[i] < asciiBegin || msg[i] > asciiEnd {
			return true
		}
	}
	return false
}

// 字段值解密方案一
func decryptPlan1(msg []byte) ([]byte, error) {
	if len(msg) == 0 {
		return nil, errors.New("decrypter empty msg")
	}
	
	base64EncodeLen := base64.StdEncoding.DecodedLen(len(msg))
	base64Decoded := make([]byte, base64EncodeLen)
	n, _ := base64.StdEncoding.Decode(base64Decoded, msg)
	decryptedBytes := xorEncrypt(base64Decoded[0:n], xorKey)
	return decryptedBytes, nil
}

// 字段值解密方案二 不支持匿名化
func decryptPlan2(msg []byte, timeKey int, platform int) ([]byte, error) {
	if len(msg) == 0 {
		return nil, errors.New("decrypter empty msg")
	}
	//如果包含 非 32~126的字符 不进行匿名化操作
	if containNonAscii(msg) {
		return msg, nil
	}
	return plan2Decrypt(msg, timeKey, platform)
}

// 字段值解密方案二 支持匿名化
func decryptPlan2Anon(msg []byte, md5 string, timeKey int, platform int) ([]byte, error) {
	if len(msg) == 0 {
		return nil, errors.New("anonymous msg is empty")
	}
	if len(md5) != 7 {
		return nil, errors.New("anonymous md5 is invalid size")
	}
	
	//如果包含 非 32~126的字符 不进行匿名化操作
	if containNonAscii(msg) {
		return msg, nil
	}
	decrypted, err := plan2Decrypt(msg, timeKey, platform)
	if err != nil {
		return nil, err
	}
	md5Upper := strings.ToUpper(md5)
	md5Sum, ok := findMd5(md5Upper)
	if !ok {
		return nil, errors.New("md5 key not found")
	}
	if md5Sum == 0 {
		return decrypted, nil
	}
	
	decryptedSum := asciiSum(decrypted)
	offsetSum := offsetSum(timeKey)
	originalSum := md5Sum - offsetSum
	if originalSum == 0 {
		return decrypted, nil
	}
	deletedChar := byte(originalSum - decryptedSum)
	offsetDelete := offsetDelete(timeKey)
	var target []byte
	if originalSum > 0 && len(decrypted) <= 8 {
		target, ok = insertByteAtEnd(decrypted, deletedChar, 0)
	} else {
		target, ok = insertByteAtEnd(decrypted, deletedChar, offsetDelete)
	}
	if !ok {
		return nil, errors.New("insert deleted char error")
	}
	return target, nil
	
}

func getMd5Key(key string) string {
	if len(key) != 3 {
		return ""
	}
	return "s" + key[0:1] + key[2:] + key[1:2]
}

func plan2Decrypt(msg []byte, timeKey, platform int) ([]byte, error) {
	var decrypted []byte
	var err error
	offset_method_index := offsetMethodIdx(timeKey)
	if offset_method_index > 3 {
		return nil, errors.New("invalid offset method index")
	}
	decryptFuncs := []func(msg []byte) ([]byte, error){decrypt0, decrypt1, decrypt2, decrypt3}
	decrypted, err = decryptFuncs[offset_method_index](msg)
	if err != nil {
		return nil, err
	}
	
	//计算 offset
	var offset int
	switch platform {
	case platformAndroid:
		offset = offsetAndroid(timeKey)
	case platformIos:
		offset = offsetIOS(timeKey)
	default:
		return nil, errors.New("Unknown platform")
	}
	//offset 偏移解码
	decoded := offsetDecode(decrypted, offset)
	return decoded, nil
}

func init() {
	md5SumCache = make(map[string]int, md5Size)
	for i := 0; i <= md5Size; i++ {
		bs := []byte(strconv.Itoa(i))
		md5Sum := md5.Sum(bs)
		md5Byte := []byte(fmt.Sprintf("%X", md5Sum))
		md5SumCache[string(md5Byte[0:7])] = i
	}
}
