package dudeploy

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type response struct {
	Content Content `json:"content"`
	Error   bool    `json:"error"`
}

type Content struct {
	Brokers string `json:"brokers"`
}

// HttpGet topicToken topic:token
func HttpGet(topicToken string) (string, error) {
	url := "http://tkafka-8W114n7y.kafka.sfcloud.local:1080/mom-mon/monitor/requestService.pub?cluster_name=isic_dfp_core_1qqpeivw&validate_type=1&topic_tokens=" + topicToken
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var res response
	err = json.Unmarshal(body, &res)
	if err != nil {
		return "", err
	}
	return res.Content.Brokers, nil
}
