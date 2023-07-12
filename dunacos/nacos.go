package dunacos

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"log"
	"strconv"
	"strings"
)

const (
	default_addr = "172.17.0.122:8888"
)

// Connect : 使用默认地址连接 nacos 并返回client
func Connect(namespace string) config_client.IConfigClient {
	nacosAddr, err := GetEnv("NACOS_ADDR")
	if err != nil {
		log.Println(err.Error())
	}
	if nacosAddr == "" {
		nacosAddr = default_addr
		log.Println("Cannot get NACOS_ADDR from env,will use default_addr:", default_addr)
	} else {
		log.Println("Got NACOS_ADDR from env:", nacosAddr)
	}
	return ConnectWithAddr(nacosAddr, namespace)
}

// ConnectWithAddr : 使用自定义地址连接 nacos 并返回client
func ConnectWithAddr(addr, namespace string) config_client.IConfigClient {
	if addr == "" {
		panic("nacos addr is empty")
	}
	index := strings.Index(addr, ":")
	if index == -1 {
		panic("nacos addr format error:[ip:port]")
	}
	var err error
	ip := addr[:index]
	port, err := strconv.Atoi(addr[index+1:])
	if err != nil {
		panic("Parse nacos port error:" + err.Error())
	}
	clientConfig := constant.ClientConfig{
		NamespaceId:     namespace,
		UpdateThreadNum: 1,
		LogLevel:        "error",
	}
	
	// At least one ServerConfig
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      ip,
			ContextPath: "/nacos",
			Port:        uint64(port),
		},
	}
	// Create config client for dynamic configuration
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		panic("Get nacos config error:" + err.Error())
	}
	return configClient
}

// GetCommonConfig : 获取通用配置信息（生产环境和测试环境可共用，且不敏感的配置信息）
//
// configClient: nacos client
// dataId: 对应nacos中的dataId,通常为项目名称
// onChange: 配置信息变更回调函数
func GetCommonConfig(configClient config_client.IConfigClient, dataId string, onChange func(data string)) string {
	return getConfig(configClient, dataId, "COMMON", onChange)
}

// GetConfig : 获取配置信息
//
// configClient: nacos client
// dataId: 对应nacos中的dataId,通常为项目名称
// onChange: 配置信息变更回调函数
func GetConfig(configClient config_client.IConfigClient, dataId string, onChange func(data string)) string {
	group, _ := GetEnv("DU_ENV")
	group = strings.TrimSpace(group)
	log.Println("Read group from env:", group)
	if group == "" {
		group = "TEST"
	}
	return getConfig(configClient, dataId, group, onChange)
}
func getConfig(configClient config_client.IConfigClient, dataId string, group string, onChange func(data string)) string {
	content, err := configClient.GetConfig(vo.ConfigParam{DataId: dataId, Group: group})
	if err != nil {
		panic(err)
	}
	if onChange != nil {
		err := configClient.ListenConfig(vo.ConfigParam{
			DataId: dataId,
			Group:  group,
			OnChange: func(namespace, group, dataId, data string) {
				onChange(data)
			},
		})
		if err != nil {
			log.Println("listen nacos error:", err.Error())
		}
	}
	return content
}
