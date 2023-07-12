/**
 * @Author: chenwei
 * @Description:
 * @Date: 2022/11/04 16:43
 */

package dugindb

import (
	"context"
	"errors"
	"git.du.com/cloud/du_component/duhbase/gen-go/hbase"
)

type Condition struct {
	K string
	V string
}

type MatchRes struct {
	ConditionIndex int     // 命中条件的索引值
	Total          int     // 索引中的value对应的主键数量
	Datas          []*Data // 主表中读出的数据详情
	Error          error
}

// Match : 单条件查询
// field:字段名，value:字段值
func (p *HbaseApi) Match(limit int, field, value string) *MatchRes {
	return p.MatchAny(limit, Condition{field, value})
}

// MatchAny : 多条件查询，命中其中一个条件即可
// conditions : 条件
func (p *HbaseApi) MatchAny(limit int, conditions ...Condition) *MatchRes {
	conditionLen := len(conditions)
	if conditionLen == 0 {
		return &MatchRes{-1, 0, nil, errors.New("conditions required")}
	}
	
	c, _ := p.hbasePool.Get()
	client := c.(*hbase.THBaseServiceClient)
	defer p.hbasePool.Put(c)
	for index, condition := range conditions {
		indexRes, _ := p.getIndex(client, buildIndexKey(String2Bytes(condition.K), String2Bytes(condition.V)))
		if indexRes != nil && len(indexRes.Row) != 0 {
			var matchVals [][]byte
			for i := range indexRes.ColumnValues {
				matchVals = append(matchVals, indexRes.ColumnValues[i].Qualifier)
			}
			
			var results []*Data
			var matchCount = 0
			for _, e := range matchVals {
				res, _ := client.Get(context.Background(), p.dataTableName, &hbase.TGet{Row: HashKey(e)})
				if res != nil && len(res.Row) != 0 {
					data := rowToData(res)
					if data != nil {
						resVals := data.IndexColumns[condition.K]
						if SliceContains(resVals, condition.V) {
							results = append(results, data)
							matchCount++
							if limit > 0 && matchCount >= limit {
								break
							}
						} else {
							p.delIndex(client, buildIndexKey(String2Bytes(condition.K), String2Bytes(condition.V)), e)
						}
					}
				}
			}
			if len(results) != 0 {
				return &MatchRes{index, len(matchVals), results, nil}
			}
		}
	}
	return &MatchRes{-1, 0, nil, nil}
}

// MatchAnyCustomize : 多条件查询，命中其中一个条件即可 (定制)
// conditions : 条件
func (p *HbaseApi) MatchAnyCustomize(conditions ...Condition) []MatchRes {
	var ret []MatchRes
	conditionLen := len(conditions)
	if conditionLen == 0 {
		return ret
	}
	
	c, _ := p.hbasePool.Get()
	client := c.(*hbase.THBaseServiceClient)
	defer p.hbasePool.Put(c)
	for index, condition := range conditions {
		indexRes, _ := p.getIndex(client, buildIndexKey(String2Bytes(condition.K), String2Bytes(condition.V)))
		if indexRes != nil && len(indexRes.Row) != 0 {
			var matchVals [][]byte
			for i := range indexRes.ColumnValues {
				matchVals = append(matchVals, indexRes.ColumnValues[i].Qualifier)
			}
			
			var results []*Data
			for _, e := range matchVals {
				res, _ := client.Get(context.Background(), p.dataTableName, &hbase.TGet{Row: HashKey(e)})
				if res != nil && len(res.Row) != 0 {
					results = append(results, rowToData(res))
				}
			}
			retLen := len(results)
			if retLen != 0 {
				ret = append(ret, MatchRes{index, len(matchVals), results, nil})
				if retLen == 1 {
					return ret
				}
			}
		}
	}
	return ret
}

// MatchAll : 多条件查询，命中所有条件
// conditions : 条件
func (p *HbaseApi) MatchAll(limit int, conditions ...Condition) *MatchRes {
	conditionLen := len(conditions)
	if conditionLen == 0 {
		return &MatchRes{-1, 0, nil, errors.New("conditions required")}
	}
	
	c, _ := p.hbasePool.Get()
	client := c.(*hbase.THBaseServiceClient)
	defer p.hbasePool.Put(c)
	var results []*Data
	var matchCount = 0
	cacheMap := make(map[string]struct{})
	for _, condition := range conditions {
		indexRes, _ := p.getIndex(client, buildIndexKey(String2Bytes(condition.K), String2Bytes(condition.V)))
		if indexRes != nil && len(indexRes.Row) != 0 {
			var matchVals [][]byte
			for i := range indexRes.ColumnValues {
				matchVals = append(matchVals, indexRes.ColumnValues[i].Qualifier)
			}
			for _, e := range matchVals {
				res, _ := client.Get(context.Background(), p.dataTableName, &hbase.TGet{Row: HashKey(e)})
				if res != nil && len(res.Row) != 0 {
					data := rowToData(res)
					if _, ok := cacheMap[data.Id]; ok {
						break
					}
					var isNotMatch bool
					for _, cond := range conditions {
						vals, ok := data.IndexColumns[cond.K]
						if ok {
							if len(vals) == 0 {
								isNotMatch = true
								break
							} else {
								if vals[0] != cond.V {
									isNotMatch = true
									break
								}
							}
						} else {
							isNotMatch = true
							break
						}
					}
					if isNotMatch {
						continue
					}
					results = append(results, data)
					cacheMap[data.Id] = struct{}{}
					matchCount++
					if limit > 0 && matchCount >= limit {
						break
					}
				}
			}
			if len(results) == 0 {
				return &MatchRes{-1, 0, nil, nil}
				//return &MatchRes{index, len(matchVals), results, nil}
			}
		} else {
			return &MatchRes{-1, 0, nil, nil}
		}
	}
	if len(results) == 0 {
		return &MatchRes{-1, 0, nil, nil}
	}
	return &MatchRes{-1, matchCount, results, nil}
}
