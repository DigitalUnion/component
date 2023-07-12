package dunsq

import (
	"errors"
	"github.com/valyala/fasthttp"
)

func Http(reqMethod, url string, reqBody []byte) ([]byte, error) {

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req) // 用完需要释放资源

	// 默认是application/x-www-form-urlencoded
	req.Header.SetContentType("application/json; charset=utf-8")
	if reqMethod == "GET" {
		req.Header.SetMethod("GET")
	}
	if reqMethod == "POST" {
		req.Header.SetMethod("POST")
	}
	//req.Header.Add("Du-From", "Forward")
	//req.Header.Add("IsSimulate", "1")

	req.SetRequestURI(url)

	req.SetBody(reqBody)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp) // 用完需要释放资源

	if err := fasthttp.Do(req, resp); err != nil {
		return []byte{}, errors.New("Http-error:" + err.Error())
	}

	return resp.Body(), nil

}
