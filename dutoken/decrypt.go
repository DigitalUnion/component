package dutoken

import (
	"crypto/aes"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var LengthError = errors.New("token length error")

type Token struct {
	DID   string
	Stamp string
}

func TokenEncrypt(did string) string {
	randBytes := make([]byte, 8)
	for i := 0; i < 8; i++ {
		randBytes = append(randBytes, byte(rand.Intn(255)))
	}
	randBytes = []byte{49, 141, 213, 157, 240, 58, 201, 62}
	con := "5B89FDD463CD395BF7F9B8F85A306859"
	randStr := hex.EncodeToString(randBytes)
	stamp := strconv.Itoa(int(time.Now().Unix()))
	stamp = "1679641685"
	data := did + ";" + stamp + ";" + con
	ret := aesEncryptECB([]byte(data), []byte(randStr))
	ret = append(ret, randBytes...)
	return base64.StdEncoding.EncodeToString(ret)
}

func AndDecrypt(raw string) (*Token, error) {

	mid, err := base64.StdEncoding.DecodeString(raw)
	if err != nil {
		return nil, err
	}
	if len(mid) <= 8 {
		return nil, LengthError
	}
	cryptData := make([]byte, len(mid)-8)
	keyBytes := make([]byte, 8)
	copy(cryptData, mid[:len(mid)-8])
	copy(keyBytes, mid[len(mid)-8:])
	key := hex.EncodeToString(keyBytes)
	ret, err := AesDecode(cryptData, []byte(key))
	if err != nil {
		return nil, err
	}
	tokenList := strings.Split(string(ret), ";")
	if len(tokenList) < 2 {
		return nil, LengthError
	}
	return &Token{
		DID:   tokenList[0],
		Stamp: tokenList[1],
	}, nil
}

func AesDecode(data, key []byte) ([]byte, error) {
	var err error
	entrypted := make([]byte, len(data))
	aesDecryptECB(entrypted, data, key)
	if err != nil {
		return nil, err
	}
	return entrypted, nil
}

func aesDecryptECB(decrypted, encrypted, key []byte) {
	cipher, _ := aes.NewCipher(generateKey(key))
	//decrypted = make([]byte, len(encrypted))
	//
	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}
	decrypted = decrypted[:trim]
	return
}

func aesEncryptECB(origData, key []byte) (encrypted []byte) {
	cipher, _ := aes.NewCipher(generateKey(key))
	length := (len(origData) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, origData)
	pad := byte(len(plain) - len(origData))
	for i := len(origData); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted = make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, cipher.BlockSize(); bs <= len(origData); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return encrypted
}

func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}
