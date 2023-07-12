/**
 * @Author: chenwei
 * @Description:
 * @Date: 2022/11/01 14:46
 */

package dugindb

import (
	"bytes"
	"git.du.com/cloud/du_component/duhbase/gen-go/hbase"
	"strings"
)

const (
	indexSuffix = "&"
)

var splitFlag = []byte{1, ';'}

type Data struct {
	Id            string              // 主键
	IndexColumns  map[string][]string // 需要建索引的字段
	NormalColumns map[string][]string // 无需建索引的字段
}

func rowToData(r *hbase.TResult_) *Data {
	if r == nil || len(r.Row) == 0 {
		return nil
	}
	d := Data{Id: Bytes2String(r.Row)}
	index := strings.Index(d.Id, "_")
	if index != -1 {
		d.Id = d.Id[index+1:]
	}
	d.IndexColumns = make(map[string][]string)
	d.NormalColumns = make(map[string][]string)
	for _, e := range r.ColumnValues {
		k := Bytes2String(e.Qualifier)
		ps := bytes.Split(e.Value, splitFlag)
		ss := make([]string, len(ps))
		for i, v := range ps {
			ss[i] = Bytes2String(v)
		}
		if strings.HasSuffix(k, indexSuffix) {
			d.IndexColumns[k[:len(k)-1]] = ss
		} else {
			d.NormalColumns[k] = ss
		}
	}
	return &d
}
func stringToIndexQualifier(s string) []byte {
	bf := bytes.Buffer{}
	bf.WriteString(s)
	bf.WriteString(indexSuffix)
	return bf.Bytes()
}
