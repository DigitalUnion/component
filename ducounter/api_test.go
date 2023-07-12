package ducounter

import (
	"testing"
	"time"
)

func TestCounterAdd(t *testing.T) {
	dsn := "root:root@(39.100.46.101:3306)/test0928?charset=utf8&parseTime=True&loc=Local"
	var tmp = make(map[string]string)
	tmp["DNA"] = "tt_dna_stat"
	tmp["DDI"] = "tt_ddi_stat"
	Init(dsn, tmp)
	for i := 0; i < 36000; i++ {
		CounterAdd("DNA")
		if i%2 == 0 {
			CounterAdd("DDI")
		}
		time.Sleep(10 * time.Microsecond)
	}
}
