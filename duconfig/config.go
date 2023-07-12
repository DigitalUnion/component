package duconfig

import (
	"bytes"
	"encoding/json"
	"errors"
	jsonIterator "github.com/json-iterator/go"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

const (
	defaultBaseURL = "http://172.17.0.130:9999"
)

var (
	defaultConfig = &DuConfig{
		url:      defaultBaseURL,
		cacheDir: getCurrentPath() + string(os.PathSeparator) + "cache",
	}
)

type DuConfig struct {
	url      string // du_config url，默认为 defaultBaseURL
	cacheDir string // 持久化配置文件的目录，默认为 cache
	env      string // 环境变量
}

type option func(c *DuConfig)

func WithUrl(url string) option {
	return func(c *DuConfig) {
		c.url = url
	}
}

func WithCacheDir(dir string) option {
	return func(c *DuConfig) {
		c.cacheDir = dir
	}
}

func WithEnv(env string) option {
	return func(c *DuConfig) {
		c.env = env
	}
}

func NewDuConfigByOptions(opts ...option) *DuConfig {
	c := defaultConfig
	for _, opt := range opts {
		opt(c)
	}
	if c.env == "" {
		c.setEnv()
	}
	return c
}

func NewDefaultConfig() *DuConfig {
	defaultConfig.setEnv()
	return defaultConfig
}

func (c *DuConfig) setEnv() {
	var err error
	c.env, err = getEnv("DU_ENV")
	log.Println("Read group from env:", c.env)
	if err != nil {
		panic(err)
	}
	if c.env == "" {
		panic("DU_ENV NOT FOUND!")
	}
}

func (c *DuConfig) Migrate(serviceName, projectName string, srcStruct interface{}) error {
	postData := make(map[string]interface{})
	postData["service_name"] = serviceName
	postData["project_name"] = projectName
	postData["env"] = c.env
	postData["configs"] = GetFields(srcStruct)
	url := c.url + "/config/migrate?token=8iXs299mj0KjgkHXCnJvxTbP5kgBaEl5"
	_, err := c.postWithJsonData(http.MethodPost, url, postData)
	return err
}

func (c *DuConfig) GetYamlContent(serviceName string) (string, error) {
	postData := make(map[string]string)
	postData["service_name"] = serviceName
	postData["env"] = c.env
	url := c.url + "/config/yaml_code?token=8iXs299mj0KjgkHXCnJvxTbP5kgBaEl5"
	data, err := c.postWithFormData(http.MethodPost, url, &postData)
	if err != nil {
		log.Printf("get config from server error: %v", err)
		content, err := readConfigFromFile(serviceName, c.cacheDir)
		if err != nil {
			log.Printf("get config from file error: %v", err)
			return "", errors.New("read config from both server and cache fail")
		}
		return content, nil
	}
	writeConfigToFile(serviceName, c.cacheDir, data)
	return data, nil
}

func (c *DuConfig) postWithFormData(method, url string, postData *map[string]string) (string, error) {
	body := new(bytes.Buffer)
	w := multipart.NewWriter(body)
	for k, v := range *postData {
		w.WriteField(k, v)
	}
	w.Close()
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", w.FormDataContentType())
	client := &http.Client{Timeout: time.Second * 3}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", err
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	resp.Body.Close()
	ret := jsonIterator.Get(bodyBytes, "data").ToString()
	return ret, nil
}

func (c *DuConfig) postWithJsonData(method, url string, postData map[string]interface{}) (string, error) {
	body, _ := json.Marshal(postData)
	reader := bytes.NewReader(body)
	req, err := http.NewRequest(method, url, reader)
	if err != nil {
		return "", err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", err
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	resp.Body.Close()
	ret := jsonIterator.Get(bodyBytes, "data").ToString()
	return ret, nil
}
