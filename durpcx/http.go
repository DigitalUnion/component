package durpcx

import "github.com/valyala/fasthttp"

func Http(reqMethod, url string, reqBody []byte) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req) // 用完需要释放资源
	req.Header.SetContentType("application/json; charset=utf-8")
	req.Header.SetMethod(reqMethod)
	req.SetRequestURI(url)
	if len(reqBody) != 0 {
		req.SetBody(reqBody)
	}

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp) // 用完需要释放资源

	if err := fasthttp.Do(req, resp); err != nil {
		return []byte{}, err
	}
	return resp.Body(), nil

}
