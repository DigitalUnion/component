/**
 * @Author: chenwei
 * @Description:
 * @Date: 2022/11/01 14:36
 */

package dugindb

import (
	"git.du.com/cloud/du_component/duhbase"
	"github.com/silenceper/pool"
)

type HbaseApi struct {
	hbasePool      pool.Pool
	dataTableName  []byte
	indexTableName []byte
}

// NewHbaseApi : 创建一个基于 hbase 的 gindb 可操作对象
func NewHbaseApi(cfg duhbase.HbaseConfig, dataTableName, indexTableName string) *HbaseApi {
	a := HbaseApi{
		hbasePool:      duhbase.NewHbasePool(&cfg),
		dataTableName:  []byte(dataTableName),
		indexTableName: []byte(indexTableName),
	}
	return &a
}
