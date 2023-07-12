package dureflect

import (
	"reflect"
	"sync"
)

var fieldIndexCache map[string]int
var fieldIndexCacheLock sync.RWMutex

func init() {
	fieldIndexCache = make(map[string]int)
}

// FieldByName : 根据name获取field
// 实现方式：
// 如果索引缓存中没有，则用FieldByName查询field，并将field的索引加入缓存
func FieldByName(value reflect.Value, name string) reflect.Value {
	key := value.Type().String() + ":" + name
	fieldIndexCacheLock.RLock()
	index, ok := fieldIndexCache[key]
	fieldIndexCacheLock.RUnlock()
	if ok {
		return value.Field(index)
	}

	if f, ok := value.Type().FieldByName(name); ok {
		i := f.Index[0]
		fieldIndexCacheLock.Lock()
		fieldIndexCache[key] = i
		fieldIndexCacheLock.Unlock()
		return value.Field(i)
	}
	return reflect.Value{}
}
