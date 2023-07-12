package duapache_hbase

import (
	"context"
	"errors"
	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/hrpc"
)

// NewApacheHbaseClient zk: zookeeper's   ip:port
func NewApacheHbaseClient(zk string) gohbase.Client {
	return gohbase.NewClient(zk)
}

// PutOne 单个写入
func PutOne(client gohbase.Client, tablename, key string, values map[string]map[string][]byte) error {
	if tablename == "" || key == "" || len(values) == 0 || client == nil {
		return errors.New("params error")
	}
	putRequest, err := hrpc.NewPutStr(context.Background(), tablename, key, values)
	if err != nil {
		return err
	}
	_, err = client.Put(putRequest)
	return err
}

// GetByKey 根据key查询
func GetByKey(client gohbase.Client, tablename, key string) (*hrpc.Result, error) {
	if tablename == "" || key == "" || client == nil {
		return nil, errors.New("params error")
	}
	getRequest, err := hrpc.NewGetStr(context.Background(), tablename, key)
	if err != nil {
		return nil, err
	}
	return client.Get(getRequest)
}

// DeleteByKey 根据key删除
func DeleteByKey(client gohbase.Client, tablename, key string) error {
	if tablename == "" || key == "" || client == nil {
		return errors.New("params error")
	}
	delRequest, err := hrpc.NewDelStr(context.Background(), tablename, key, nil)
	if err != nil {
		return err
	}
	_, err = client.Delete(delRequest)
	return err
}

// DeleteByKeyAndValues  根据key,列簇,列删除
// To delete entire row, values should be nil.
//
// To delete specific families, qualifiers map should be nil:
//
//	 map[string]map[string][]byte{
//			"cf1": nil,
//			"cf2": nil,
//	 }
//
// To delete specific qualifiers:
//
//	 map[string]map[string][]byte{
//	     "cf": map[string][]byte{
//				"q1": nil,
//				"q2": nil,
//			},
//	 }
//
func DeleteByKeyAndValues(client gohbase.Client, tablename, key string, values map[string]map[string][]byte) error {
	if tablename == "" || key == "" || client == nil {
		return errors.New("params error")
	}
	delRequest, err := hrpc.NewDelStr(context.Background(), tablename, key, values)
	if err != nil {
		return err
	}
	_, err = client.Delete(delRequest)
	return err
}
