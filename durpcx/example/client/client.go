package client

import (
	"git.du.com/cloud/du_component/durpcx"
	"git.du.com/cloud/du_component/durpcx/example/common"
)

var rpcClient *durpcx.RpcClient

func Init(cfg *durpcx.Config) {
	rpcClient = durpcx.NewClient(*cfg)
}

func Hello(data []byte) *common.ExampleRes {
	req := common.ExampleReq{
		Ip:   "localhost",
		Data: data,
	}

	res := common.ExampleRes{}
	// 这里的 req 和 res 须实现common.go 中 Req 和 Res 的接口方法
	rpcClient.Call("Hello", &req, &res)
	return &res
}
