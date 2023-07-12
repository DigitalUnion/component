/**
 * @Author: chenwei
 * @Description:
 * @Date: 2022/11/10 11:17
 */

package dudecoder

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

func HttpPost(url string, header map[string]string, body []byte) ([]byte, error) {
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
	return resp.Body(), nil
}
