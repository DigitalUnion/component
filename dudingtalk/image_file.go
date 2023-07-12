package dudingtalk

import (
	"bytes"
	jsoniter "github.com/json-iterator/go"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

// UploadFileToImgServer 上传文件到图片服务器
func UploadFileToImgServer(filename string) (string, error) {
	req, err := newfileUploadRequest("http://172.17.129.178:8451/upload", nil, "file", filename)
	if err != nil {
		return "", err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//fmt.Println(err.Error())
		return "", err
	}
	return jsoniter.Get(body, "data", "url").ToString(), nil
}

func newfileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, path)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)
	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodPost, uri, body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	return request, err
}
