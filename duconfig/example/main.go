package main

//type name struct {
//	Name string        `yaml:"n_ame"`
//	Age  durpcx.Config `yaml:"a_ge"`
//}

//type Config struct {
//	A
//	RpcCfg         *durpcx.Config      `json:"rpc_cfg" yaml:"rpc_cfg" desc:"rpc_cfg info"`
//	BlackListRedis duredis.RedisConfig `json:"blacklist_redis" yaml:"blacklist_redis" desc:"blacklist_redis info"`
//	EndRedis       duredis.RedisConfig `json:"end_redis" yaml:"end_redis" desc:"end_redis info"`
//	GlobalRedis    duredis.RedisConfig `json:"global_redis" yaml:"global_redis" desc:"global_redis info"`
//}

func main() {
	// 使用默认配置实例，获取服务名字是 H5_UNLOAD 的配置 yaml 文件
	/*c := duconfig.NewDefaultConfig()
	content, err := c.GetYamlContent("H5_UNLOAD")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))*/

	// 指定 du_config url，并指定本地持久化配置目录，获取服务 id 是 7 的配置 yaml 文件
	//url := "http://127.0.0.1:8080"
	//c := duconfig.NewDuConfigByOptions(duconfig.WithUrl(url), duconfig.WithEnv("CAOCAO-PROD"))
	/*content, err := c.GetYamlContent("DNA_FRONT")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))*/
	//err := c.Migrate("migrate_test", "migrate", Config{})
	//if err != nil {
	//	panic(err)
	//}
	//fields := duconfig.GetFields(Config{
	//	A: A{},
	//})
	//fmt.Printf("%+v", fields)

}

type A struct {
	Name string `yaml:"n_ame"`
}
