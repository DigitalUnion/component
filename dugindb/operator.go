/**
 * @Author: chenwei
 * @Description:
 * @Date: 2022/11/04 16:56
 */

package dugindb

import (
	"bytes"
	"context"
	"git.du.com/cloud/du_component/duhbase/gen-go/hbase"
	"math/big"
)

// GetData: 根据主键从主表中获取数据详情
func (p *HbaseApi) GetData(id string) (*Data, error) {
	client, _ := p.hbasePool.Get()
	defer p.hbasePool.Put(client)
	res, err := client.(*hbase.THBaseServiceClient).Get(context.Background(), p.dataTableName, &hbase.TGet{Row: HashKey(String2Bytes(id))})
	if err != nil {
		return nil, err
	}
	return rowToData(res), nil
}

// Put: 插入数据（如数据已存在则更新）
// oldData不必填
func (p *HbaseApi) Put(newData *Data, oldData *Data) error {
	c, _ := p.hbasePool.Get()
	client := c.(*hbase.THBaseServiceClient)
	defer p.hbasePool.Put(c)

	// check old newData
	if oldData == nil {
		oldData, _ = p.GetData(newData.Id)
	}

	// put newData
	tput := hbase.TPut{Row: HashKey(String2Bytes(newData.Id))}
	for k, v := range newData.IndexColumns {
		//if len(v) == 0 {
		//	continue
		//}
		if len(v) > 100 {
			continue
		}
		cv := hbase.TColumnValue{Family: []byte{'f'}, Qualifier: stringToIndexQualifier(k), Value: bytes.Join(Strings2Bytes(v), splitFlag)}
		tput.ColumnValues = append(tput.ColumnValues, &cv)

		var l, _, r []string
		if oldData != nil {
			l, _, r = GetSliceLMR(oldData.IndexColumns[k], newData.IndexColumns[k])
			if len(l) == 0 && len(r) == 0 {
				continue
			}
		} else {
			r = v
		}
		if len(l) != 0 {
			for _, e := range l {
				indexKey := buildIndexKey(String2Bytes(k), String2Bytes(e))
				p.delIndex(client, indexKey, String2Bytes(newData.Id))
			}
		}
		if len(r) != 0 {
			for _, e := range r {
				if len(k) == 0 || len(e) == 0 {
					continue
				}
				indexKey := buildIndexKey(String2Bytes(k), String2Bytes(e))
				p.putIndex(client, indexKey, String2Bytes(newData.Id))
			}
		}
	}
	for k, v := range newData.NormalColumns {
		//if len(v) == 0 {
		//	continue
		//}
		//var l, _, r []string
		//if oldData != nil {
		//	l, _, r = GetSliceLMR(oldData.NormalColumns[k], newData.NormalColumns[k])
		//	if len(l) == 0 && len(r) == 0 {
		//		continue
		//	}
		//}
		cv := hbase.TColumnValue{Family: []byte{'f'}, Qualifier: String2Bytes(k), Value: bytes.Join(Strings2Bytes(v), splitFlag)}
		tput.ColumnValues = append(tput.ColumnValues, &cv)
	}

	err := client.Put(context.Background(), p.dataTableName, &tput)
	if err != nil {
		return err
	}
	return nil
}

// Del: 删除数据
func (p *HbaseApi) Del(id string) error {
	c, _ := p.hbasePool.Get()
	client := c.(*hbase.THBaseServiceClient)
	defer p.hbasePool.Put(c)

	data, _ := p.GetData(id)
	if data == nil {
		return nil
	}

	// del index
	for k, v := range data.IndexColumns {
		for _, e := range v {
			if len(k) == 0 || len(e) == 0 {
				continue
			}
			indexKey := buildIndexKey(String2Bytes(k), String2Bytes(e))
			p.delIndex(client, indexKey, String2Bytes(data.Id))
		}
	}
	// del data
	return p.delData(client, String2Bytes(id))
}

func (p *HbaseApi) getIndex(client *hbase.THBaseServiceClient, k []byte) (*hbase.TResult_, error) {
	return client.Get(context.Background(), p.indexTableName, &hbase.TGet{Row: HashKey(k)})
}
func (p *HbaseApi) delIndex(client *hbase.THBaseServiceClient, k, v []byte) error {
	column := hbase.TColumn{Family: []byte{'f'}, Qualifier: v}
	columns := make([]*hbase.TColumn, 1)
	columns[0] = &column
	toDel := hbase.TDelete{
		Row:        HashKey(k),
		Columns:    columns,
		DeleteType: hbase.TDeleteType_DELETE_COLUMNS,
	}
	return client.DeleteSingle(context.Background(), p.indexTableName, &toDel)
}
func (p *HbaseApi) delData(client *hbase.THBaseServiceClient, k []byte) error {
	column := hbase.TColumn{Family: []byte{'f'}}
	columns := make([]*hbase.TColumn, 1)
	columns[0] = &column
	toDel := hbase.TDelete{
		Row:        HashKey(k),
		Columns:    columns,
		DeleteType: hbase.TDeleteType_DELETE_FAMILY,
	}
	return client.DeleteSingle(context.Background(), p.dataTableName, &toDel)
}

func (p *HbaseApi) putIndex(client *hbase.THBaseServiceClient, k, v []byte) error {
	if len(k) == 0 || len(v) == 0 {
		return nil
	}
	columnValue := hbase.TColumnValue{Family: []byte{'f'}, Qualifier: v, Value: []byte{1}}
	columnValues := make([]*hbase.TColumnValue, 1)
	columnValues[0] = &columnValue
	tput := &hbase.TPut{Row: HashKey(k), ColumnValues: columnValues}
	return client.Put(context.Background(), p.indexTableName, tput)
}
func buildIndexKey(field, value []byte) []byte {
	return bytes.Join([][]byte{value, field}, []byte{'_'})
}
func HashKey(k []byte) []byte {
	var sum uint64
	for _, e := range k {
		sum += uint64(e)
	}
	sum = sum << 2
	if sum > 238328 {
		sum %= 238328
	}
	i := Int10To62(uint64(sum))
	bf := bytes.Buffer{}
	bf.WriteString(i)
	bf.WriteByte('_')
	bf.Write(k)
	return bf.Bytes()
}
func Int10To62(s uint64) string {
	n := new(big.Int)
	t2 := n.SetUint64(s)
	return t2.Text(62)
}
