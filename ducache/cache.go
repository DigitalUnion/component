package ducache

import (
	"errors"
	"github.com/coocood/freecache"
)

var api *freecache.Cache

func NewCache(cacheSize int) {
	api = freecache.NewCache(cacheSize)
}
func GetString(k string) string {
	if api == nil {
		panic("Please call NewCache()")
	}
	key := String2Bytes(k)
	v, _ := api.Get(key)
	val := Bytes2String(v)
	return val
}
func SetString(k, v string, expireSeconds int) error {
	if api == nil {
		return errors.New("CacheApi is nil")
	}
	key := String2Bytes(k)
	val := String2Bytes(v)
	return api.Set(key, val, expireSeconds)
}
func GetApi() *freecache.Cache {
	return api
}
func EntryCount() int64 {
	return api.EntryCount()
}
