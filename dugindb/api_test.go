/**
 * @Author: chenwei
 * @Description:
 * @Date: 2022/11/01 15:56
 */

package dugindb

import (
	"fmt"
	"git.du.com/cloud/du_component/duhbase"
	"log"
	"testing"
)

const (
	address        = "http://ld-2zehayx6ygihdl6qj-proxy-lindorm.lindorm.rds.aliyuncs.com:9190"
	dataTableName  = "ab_test:data_test"
	indexTableName = "ab_test:index_test"
)

/*
*
address: http://ld-2zed8a8y16bd4lx43-proxy-hbaseue.hbaseue.rds.aliyuncs.com:9190

	user: root
	pw: root
	namespace: mdna
*/
func TestNewHbaseApiPut(t *testing.T) {
	cfg := duhbase.HbaseConfig{
		Address:  address,
		User:     "root",
		Password: "root",
	}
	api := NewHbaseApi(cfg, dataTableName, indexTableName)
	data := Data{
		Id: "k7",
		IndexColumns: map[string][]string{
			"macs": []string{"a7", "a8"},
		},
		NormalColumns: map[string][]string{
			"model": []string{"HUAWEI8"},
		},
	}
	api.Put(&data, nil)
}
func TestNewHbaseApiDel(t *testing.T) {
	cfg := duhbase.HbaseConfig{
		Address:  address,
		User:     "root",
		Password: "root",
	}
	api := NewHbaseApi(cfg, dataTableName, indexTableName)
	res := api.Del("k6")
	log.Printf("%+v\n", res)
}
func TestNewHbaseApiGet(t *testing.T) {
	cfg := duhbase.HbaseConfig{
		Address:  address,
		User:     "root",
		Password: "root",
	}
	api := NewHbaseApi(cfg, dataTableName, indexTableName)
	res, _ := api.GetData("k6")
	log.Printf("%v\n", res)
}
func TestNewHbaseApiMatch(t *testing.T) {
	cfg := duhbase.HbaseConfig{
		Address:  address,
		User:     "root",
		Password: "root",
	}
	api := NewHbaseApi(cfg, dataTableName, indexTableName)
	// 单条件查询，如有多个did，只返回2个
	res := api.Match(2, "macs", "a7")
	log.Printf("%+v\n", *res)
	for _, e := range res.Datas {
		log.Printf("%+v\n", e)
	}
	
	// 多条件查询,b不限返回个数
	//res = api.MatchAny(-1, Condition{"macs", "aa:bb:cc:d21"}, Condition{"macs", "aa:bb:cc:d21"}, Condition{"imei", "878877676271"})
	//log.Printf("Total:%+v\n", res.Total)
	//for _, e := range res.Datas {
	//	log.Printf("%+v\n", e)
	//}
}

func Test_hashKey(t *testing.T) {
	for i := 0; i < 100; i++ {
		id := HashKey([]byte("D2HUfOb4Ce5AV8MjmydxXhsJYudg+o4y85MVhcDuZ7gRMXd2"))
		log.Println(string(id))
	}
}

func TestHbaseApi_MatchAll(t *testing.T) {
	cfg := duhbase.HbaseConfig{
		Address:  address,
		User:     "root",
		Password: "root",
	}
	api := NewHbaseApi(cfg, dataTableName, indexTableName)
	var conds []Condition
	conds = append(conds, Condition{
		K: "cid",
		V: "102",
	})
	conds = append(conds, Condition{
		K: "strategy",
		V: "Z1002",
	})
	
	matchRes := api.MatchAll(-1, conds...)
	fmt.Println(matchRes.Total)
}
