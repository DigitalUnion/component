package tools

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
	"strings"
)

func HttpGet(url string) ([]byte, error) {
	status, resp, err := fasthttp.Get(nil, url)
	if err != nil {
		// 请求发生错误
		return []byte{}, err
	}

	if status != fasthttp.StatusOK {
		return []byte{}, errors.New(fmt.Sprintf("请求失败，status: %d", status))
	}

	return resp, nil
}

func HttpPost(url string, header map[string]string, body []byte) ([]byte, error) {
	if !strings.Contains(url, "http://") {
		url = "http://" + url
	}
	req := &fasthttp.Request{}
	req.SetRequestURI(url)
	if len(body) != 0 {
		req.SetBody(body)
	}

	// 默认是application/x-www-form-urlencoded
	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")
	for k, v := range header {
		req.Header.Set(k, v)
	}

	resp := &fasthttp.Response{}

	client := &fasthttp.Client{}
	if err := client.Do(req, resp); err != nil {
		fmt.Println("请求失败:", err.Error())
		return []byte{}, err
	}
	var res Resp
	err := json.Unmarshal(resp.Body(), &res)
	if err != nil {
		return nil, err
	}
	return []byte(res.Data), nil
}

type Resp struct {
	Code    int    `json:"code"`
	Data    string `json:"data"`
	Message string `json:"message"`
}
