package dulogv2

import (
	"fmt"
	"strings"
)

type Config struct {

	//应用程序名称 例如 dna 名称中不能包含 '-'  必须 (远程配置/本地配置)
	Appid string `json:"appid" yaml:"appid"`

	//模块名称 例如 run,biz,raw,err  名称中不能包含 '-'  为空默认存储在default模块中 非必须 (远程配置/本地配置)
	Module string `json:"module" yaml:"module"`

	//rsyslog 服务地址  ip:port   必须  （远程配置）
	Addrs []string `json:"addrs" yaml:"addrs"`

	//开启健康检查 会有性能损耗  默认关闭 开启用true 非必须 （远程配置）
	IsCheck bool `json:"is_check" yaml:"is_check"`

	//开启es日志  默认关闭 开启用true 非必须 （远程配置）
	IsEs bool `json:"is_es" yaml:"is_es"`

	//以上是远程配置 如果只需要远程日志 到此为止
	//以下是本地配置 如果需要本地日志 需要配置以下参数

	//开启本地日志  默认值:false , 开启用true  开启后远程日志不生效 非必须  (本地配置)
	IsLoc bool `json:"is_loc" yaml:"is_loc"`

	//本地日志文件存储路径 默认值: ./logs , 如果存本机需配置 非必须   (本地配置)
	LocDir string `json:"loc_dir" yaml:"loc_dir"`

	//本地日志文件最大大小  默认是 10240 (10G),  非必须  (本地配置)
	LocMaxSize int `json:"loc_max_size" yaml:"max_size"`

	//本地日志回滚策略 默认: 0 0 * * * * (每小时回滚一次) , 非必须  (本地配置)
	LocRotate string `json:"loc_rotate" yaml:"loc_rotate"`

	//本地日志文件最大保存时间  默认值: 1 (1day) 单位day,  非必须  (本地配置)
	LocMaxAge int `json:"loc_max_age" yaml:"loc_max_age"`

	//本地日志文件是否压缩  默认是不压缩  非必须  (本地配置)
	LocCompress bool `json:"loc_compress" yaml:"loc_compress"`
}

func (c *Config) valid() {
	if c.Appid == "" {
		panic("dulog appid is empty")
	}
	if c.Module == "" {
		fmt.Println("dulog module is empty,default fill value 'default'")
		c.Module = defaultModule
	}
	if strings.Contains(c.Appid, logSplitChar) || strings.Contains(c.Module, logSplitChar) {
		fmt.Println("dulog appid or module contains '-'")
		panic("dulog appid or module contains '-';app:" + c.Appid + ",module:" + c.Module)
	}
	if c.IsLoc {
		if c.LocDir == "" {
			c.LocDir = locDefaultDir
		}
		if c.LocMaxSize == 0 {
			c.LocMaxSize = locDefaultMaxSize
		}
		if c.LocRotate == "" {
			c.LocRotate = locDefaultRotate
		}
		if c.LocMaxAge == 0 {
			c.LocMaxAge = locDefaultMaxAge
		}
	} else {
		if c.Addrs == nil {
			panic("dulog addrs is empty")
		}
		if !c.checkAddrsValid() {
			panic("dulog addrs is dup")
		}
	}
}

func (c *Config) checkAddrsValid() bool {
	m := make(map[string]struct{})
	for _, addr := range c.Addrs {
		m[addr] = struct{}{}
	}
	if len(m) != len(c.Addrs) {
		return false
	}
	return true
}
