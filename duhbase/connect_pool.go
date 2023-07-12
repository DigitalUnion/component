package duhbase

import (
	"fmt"
	"git.du.com/cloud/du_component/duhbase/gen-go/hbase"
	"github.com/apache/thrift/lib/go/thrift"
	"net/http"
	"time"
)

var defaultHttpClient *http.Client
var cfg *HbaseConfig

func init() {
	defaultHttpClient = http.DefaultClient
	defaultHttpClient.Timeout = time.Second * 20

	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = 2000
	http.DefaultTransport.(*http.Transport).MaxConnsPerHost = 2000
	http.DefaultTransport.(*http.Transport).IdleConnTimeout = 30 * time.Second
}
func newClient() (interface{}, error) {
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transport, err := NewHttpClient(cfg.Address)
	if err != nil {
		fmt.Printf("error resolving address:%s\n", err)
		return nil, err
	}
	httClient := transport.(*thrift.THttpClient)
	httClient.SetHeader("ACCESSKEYID", cfg.User)
	httClient.SetHeader("ACCESSSIGNATURE", cfg.Password)
	hbaseClient := hbase.NewTHBaseServiceClientFactory(transport, protocolFactory)
	if err := transport.Open(); err != nil {
		fmt.Printf("Error opening %s "+cfg.Address+"\n", err)
		return nil, err
	}
	return hbaseClient, nil
}

func NewHttpClient(urlstr string) (thrift.TTransport, error) {
	return thrift.NewTHttpClientWithOptions(urlstr,
		thrift.THttpClientOptions{
			Client: defaultHttpClient,
		})
}

func closeClient(c interface{}) error {
	return nil
}
