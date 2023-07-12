package dunsq

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetNsqdNodesInfo(t *testing.T) {
	getNsqdNodesInfo([]string{"192.168.200.224:4161"})
}

func TestNsq(t *testing.T) {
	Connect(NsqConfig{
		Addrs: "192.168.200.224:4161,192.168.200.224:4261",
		Topics: []string{
			"t1",
		},
	})
	for i := 0; i < 100; i++ {
		PushString("t1", "qqqqqqqpppppppp11111")
	}
}

func TestCreateTopic(t *testing.T) {
	nsqdAddr := "http://192.168.200.224:4151/topic/create?topic=t2"
	body, err := Http(http.MethodPost, nsqdAddr, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(body))
}
