package duinflux

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func post(url string, token string, reqBody string) string {
	req, _ := http.NewRequest("POST", url, strings.NewReader(reqBody))
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	rsp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}

func HttpPost(uri string, body []byte) ([]byte, error) {
	escapeUrl := url.QueryEscape(string(body))
	req := &fasthttp.Request{}
	req.SetRequestURI(uri)
	req.SetBody([]byte(escapeUrl))

	req.Header.SetContentType("application/x-www-form-urlencoded")
	//req.Header.SetMethod("POST")
	req.Header.Add("Authorization", "Bearer eyJrIjoiM1RBYUtKc3cwd2hUUTQyanQ1RkdEckZXV1NTaHdhekIiLCJuIjoiY2xvdWQiLCJpZCI6MX0=\n")

	resp := &fasthttp.Response{}

	client := &fasthttp.Client{}
	if err := client.Do(req, resp); err != nil {
		fmt.Println("请求失败:", err.Error())
		return []byte{}, err
	}
	return resp.Body(), nil
}
