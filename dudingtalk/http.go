package dudingtalk

import (
	"github.com/valyala/fasthttp"
	"net/http"
)

// SendPost POET 请求
func SendPost(url string, body []byte) ([]byte, int) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.Header.SetContentType("application/json;charset=utf-8")
	req.Header.SetMethod("POST")
	req.SetRequestURI(url)
	req.SetBody(body)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	err := fasthttp.Do(req, resp)
	if err != nil {
		return nil, http.StatusBadGateway
	}
	return resp.Body(), resp.StatusCode()
}
