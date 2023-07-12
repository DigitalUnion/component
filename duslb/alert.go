package duslb

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
	"log"
	"time"
)

const (
	NoticeUrl string = "https://oapi.dingtalk.com/robot/send?access_token=43ea364b107271283d3e328d215194d7dc07d5ff5c7d8d9c8d1f0ce9881c5379"
)

func httpPost(url string, reqBody []byte, sign string) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req) // 用完需要释放资源
	
	req.Header.SetContentType("application/json; charset=utf-8")
	req.Header.SetMethod("POST")
	if len(sign) > 0 {
		req.Header.Add("X-Cy-Sign", sign)
	}
	
	req.SetRequestURI(url)
	
	req.SetBody(reqBody)
	
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp) // 用完需要释放资源
	
	if err := fasthttp.Do(req, resp); err != nil {
		return []byte{}, err
	}
	return resp.Body(), nil
}

func Notice(action, lb, instanceId string, start time.Time, alertErr error) {
	errMsg := "nil"
	if alertErr != nil {
		errMsg = alertErr.Error()
	}
	now := time.Now()
	body := make(map[string]interface{})
	body["msgtype"] = "markdown"
	body["markdown"] = map[string]string{"title": "负载均衡调整", "text": fmt.Sprintf(
		"### %s \n - lb id: %s \n- ecs: %s \n - error: %s \n - 请求开始时间: %s \n  - 请求结束时间: %s \n ",
		action, lb, instanceId, errMsg, start.Format("01-02 15:04:05"), now.Format("01-02 15:04:05"))}
	req, _ := jsoniter.Marshal(body)
	url := NoticeUrl
	_, err := httpPost(url, req, "")
	if err != nil {
		log.Println("ding send msg err: ", err.Error())
	}
}
