package duinflux

type InfluxConfig struct {
	Token        string
	Bucket       string
	Url          string
	Org          string
	GrafanaUrl   string
	PathList     []string // options 黑白名单
	PathListType int      // options 名单类型：-1:黑名单,1:白名单
	BatchSize    int      // options 累计内部累计数量,default:50000
}
