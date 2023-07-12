package server

import "git.du.com/cloud/du_component/durpcx"

var rpcServer *durpcx.RpcServer

func Init(cfg *durpcx.Config) {
	if cfg == nil {
		return
	}
	var err error
	rpcServer, err = durpcx.NewServer(*cfg, new(ExampleHandler))
	if err != nil {
		panic(err)
	}
}
func Close() {
	if rpcServer != nil {
		rpcServer.Stop()
	}
}
