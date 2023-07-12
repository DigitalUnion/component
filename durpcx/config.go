package durpcx

type Config struct {
	RpcPort            int    // RPC服务端口，服务端配置，客户端无需配置
	ServicePath        string // 服务路径，一个服务使用一个，须保证唯一
	Metadata           string // meta信息，非必须
	SerializeType      byte   // 序列化方式
	TimeoutMillseconds int64  // 请求超时时间，单位为毫秒（仅客户端端有效）
	ReadTimeout        int64  // 服务端read超时时间，单位为毫秒（仅服务端有效）
	WriteTimeout       int64  // 服务端write超时时间，单位为毫秒（仅服务端有效）
	ConsulServer       string // consul 连接地址
	RedisServer        string // redis 连接地址
	K8sService         string // k8s service name
	RedisPassword      string // redis 密码
	RedisDb            string // redis db
	Group              string // 分组
	CheckStateUrl      string // 检查服务器是否监控的url
	OnUnregister       func()
}
