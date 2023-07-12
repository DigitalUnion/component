# gindb

## 特性
- 提供单字段和数组的存储，以及倒排索引的实现
- 所有字段可选是否创建索引
- 索引可实现一对多映射
- 可直接新增字段，无需额外操作
- 新增字段可自动创建索引，无需额外操作

## 例子
### 写入（如主键已存在则更新）
```go

func TestNewHbaseApiPut(t *testing.T) {
    cfg := duhbase.HbaseConfig{
        Address:  address,
        User:     "root",
        Password: "root",
    }
    api := NewHbaseApi(cfg, dataTableName, indexTableName)
    data := Data{
        Id: "k6",
        IndexColumns: map[string][]string{ // 需要建索引的字段
            "macs": []string{"a6"},
        },
        NormalColumns: map[string][]string{ // 无需建索引的字段
            "model": []string{"HUAWEI7"},
        },
    }
    api.Put(&data,nil)
}
```

### 查询
```go

func TestNewHbaseApiGet(t *testing.T) {
    cfg := duhbase.HbaseConfig{
        Address:  address,
        User:     "root",
        Password: "root",
    }
    api := NewHbaseApi(cfg, dataTableName, indexTableName)
    res, _ := api.GetData("k6") // 按主键查询
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
    
    // 多条件查询,命中一个条件即结束；不限返回个数
    res = api.MatchAny(-1, Condition{"macs", "aa:bb:cc:d21"}, Condition{"macs", "aa:bb:cc:d21"}, Condition{"imei", "878877676271"})
    log.Printf("Total:%+v\n", res.Total)
    for _, e := range res.Datas {
        log.Printf("%+v\n", e)
    }
}
```

### 删除
```go

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
```