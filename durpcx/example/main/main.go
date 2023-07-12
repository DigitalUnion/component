package main

import (
	"git.du.com/cloud/du_component/durpcx"
	"git.du.com/cloud/du_component/durpcx/example/client"
	"git.du.com/cloud/du_component/durpcx/example/server"
	"log"
)

func main() {
	cfg := durpcx.Config{
		RpcPort:            8888,
		ServicePath:        "hello",
		Metadata:           "",
		TimeoutMillseconds: 1000,
		RedisServer:        "r-2ze5zgn5ul01p8iqrz.redis.rds.aliyuncs.com:6379",
		RedisPassword:      "v2wyewfBJ0bD22GZ",
		RedisDb:            "15",
		//ConsulServer:       "172.17.129.178:8500",
	}
	// 以上配置信息应写在配置文件中，而不是硬编码

	// init server
	server.Init(&cfg)
	log.Println("Start server success")

	// init client
	client.Init(&cfg)
	res := client.Hello([]byte("World"))
	log.Println(string(res.Data))

	select {}
}
