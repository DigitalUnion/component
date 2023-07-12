package durpcx

import (
	"github.com/gin-gonic/gin"
	"strings"
)

type HttpReq struct {
	Ip     string
	Uri    string
	Query  map[string][]string
	Header map[string][]string
	Body   []byte      // http body,默认没有赋值，可根据需要自行赋值
	Data   interface{} // 自定义参数，可根据需要自行赋值
}
type HttpRes struct {
	Code  int
	Error error
	Data  []byte
}

// GinToRpcReq : 将gin请求转换为 HttpReq
func GinToRpcReq(c *gin.Context) *HttpReq {
	r := HttpReq{
		Ip:    c.ClientIP(),
		Query: c.Request.URL.Query(),
		Uri:   c.Request.RequestURI,
	}
	if len(c.Request.Header) != 0 {
		r.Header = make(map[string][]string)
		for k, v := range c.Request.Header {
			key := strings.ToLower(k)
			r.Header[key] = v
		}
	}
	return &r
}
func (p HttpReq) GetClientIp() string {
	return p.Ip
}
func (p *HttpReq) GetHeader(k string) string {
	k = strings.ToLower(k)
	val := p.Header[k]
	if len(val) != 0 {
		return val[0]
	}
	return ""
}
func (p *HttpReq) GetQuery(k string) string {
	val := p.Query[k]
	if len(val) != 0 {
		return val[0]
	}
	return ""
}

func (p *HttpRes) GetResCode() int {
	return p.Code
}
func (p *HttpRes) SetError(err error) {
	p.Code = 400
	p.Error = err
}
func (p *HttpRes) SetData(data []byte) {
	p.Code = 200
	p.Data = data
}
