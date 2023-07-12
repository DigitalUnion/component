package local_cache

import (
	"errors"
	"github.com/coocood/freecache"
	"sync"
)

var LocalCache *freecache.Cache
var CacheMutex sync.RWMutex

func GetNewCache() *freecache.Cache {
	return freecache.NewCache(100 * 1024 * 1024)
}

func UpdateCache(cache *freecache.Cache) {
	CacheMutex.Lock()
	defer CacheMutex.Unlock()
	LocalCache = cache
}

func GetLocalCacheValue(key []byte) []byte {
	CacheMutex.RLock()
	defer CacheMutex.RUnlock()
	if LocalCache == nil {
		return nil
	}
	v, err := LocalCache.Get(key)
	if err != nil {
		return nil
	}
	return v
}

func SetCacheValueNoLock(key []byte, value []byte, cache *freecache.Cache) error {
	if cache == nil {
		return errors.New("cache is nil")
	}
	return cache.Set(key, value, -1)
}
