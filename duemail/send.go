package duemail

import (
	"bytes"
	"errors"
	jsoniter "github.com/json-iterator/go"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var Json = jsoniter.ConfigCompatibleWithStandardLibrary

const (
	EMAIL_DOMAIN         = "https://mailer.shuzilm.cn"
	EMAIL_SEND           = EMAIL_DOMAIN + "/api/mailer/normal"
	EMAIL_SEND_WITH_FILE = EMAIL_DOMAIN + "/api/mailer/attachment"
	EMAIL_UPLOAD_FILE    = EMAIL_DOMAIN + "/api/uploads"
	EMAIL_TOKEN          = "17a03a2fbbd249fbf2126abed4c10367"

	COMMON_SQPARATOR = ","
)

type Email struct {
	uploadFiles []uploadFilesRet
	Result      func(res string, err error)
}

type EmailNormalParams struct {
	Content   string   `json:"content"`    // 邮件正文
	To        string   `json:"to"`         // 收件人列表 多个邮件地址用 , 分割
	Cc        string   `json:"cc"`         // 抄送人列表 多个邮件地址用 , 分割
	Subject   string   `json:"subject"`    // 邮件主题
	FilesPath []string `json:"files_path"` // 需要携带的附件
}

type EmailPostJsonBody struct {
	Token       string        `json:"token"`
	Subject     string        `json:"subject"`
	Content     string        `json:"content"`
	Cc          []string      `json:"cc"`
	To          []string      `json:"to"`
	Attachments []Attachments `json:"attachments"`
}

type Attachments struct {
	Filename string `json:"filename"`
	Path     string `json:"path"`
}

type uploadFilesRet struct {
	Filename string `json:"filename"`
	Path     string `json:"path"`
}

/**
* 频率限制：
* 同一 token、同一 ip、同一 subject 在 1 分内请求上限制 10 次，超出将锁定 10 分钟
* 在锁定期间发生请求将会重置锁定时间
 */
func (e *Email) SendEmail(params EmailNormalParams) {
	var (
		err error
		res string
	)

	if len(params.FilesPath) == 0 {
		err, res = e.sendEmailNotContainFile(params)
	} else {
		err, res = e.sendEmailContainFile(params)
	}
	e.Result(res, err)
}

func (e *Email) sendEmailNotContainFile(params EmailNormalParams) (err error, res string) {
	fields := make(map[string]string)
	fields["to"] = params.To
	fields["subject"] = params.Subject
	fields["content"] = params.Content
	if params.Cc != "" {
		fields["cc"] = params.Cc
	}
	fields["token"] = EMAIL_TOKEN
	return postForm(EMAIL_SEND_WITH_FILE, fields, params.FilesPath)
}

func (e *Email) sendEmailContainFile(params EmailNormalParams) (err error, res string) {
	err = e.UploadsFiles(params.FilesPath)
	if err != nil {
		return
	}

	if len(e.uploadFiles) == 0 {
		return errors.New("no file uploaded"), res
	}

	tempAttachments := []Attachments{}
	for i := 0; i < len(e.uploadFiles); i++ {
		temp := Attachments{
			Filename: e.uploadFiles[i].Filename,
			Path:     e.uploadFiles[i].Path,
		}
		tempAttachments = append(tempAttachments, temp)
	}

	tempBody := EmailPostJsonBody{
		Token:       EMAIL_TOKEN,
		To:          strings.Split(params.To, COMMON_SQPARATOR),
		Subject:     params.Subject,
		Content:     params.Content,
		Cc:          []string{},
		Attachments: tempAttachments,
	}

	if params.Cc != "" {
		tempBody.Cc = strings.Split(params.Cc, COMMON_SQPARATOR)
	}

	tempByte, err := Json.Marshal(tempBody)
	if err != nil {
		return
	}

	return postJson(EMAIL_SEND, tempByte)
}

/**
upload mail Files
*/
func (e *Email) UploadsFiles(fileNmaes []string) (err error) {
	fields := make(map[string]string)
	fields["token"] = EMAIL_TOKEN
	err, res := postForm(EMAIL_UPLOAD_FILE, fields, fileNmaes)
	if err != nil {
		return
	}
	err = Json.UnmarshalFromString(res, &e.uploadFiles)
	return
}

/**
表单请求
*/
func postForm(url string, fields map[string]string, filename []string) (err error, res string) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	if len(filename) != 0 {
		for i := range filename {
			fileWriter, err1 := bodyWriter.CreateFormFile("uploadfile"+strconv.Itoa(i), filename[i])
			if err1 != nil {
				return err1, res
			}
			fh, err1 := os.Open(filename[i])
			if err1 != nil {
				return err1, res
			}
			defer fh.Close()
			_, err1 = io.Copy(fileWriter, fh)
			if err1 != nil {
				return err1, res
			}
		}
	}

	if len(fields) != 0 {
		for i := range fields {
			err = bodyWriter.WriteField(i, fields[i])
			if err != nil {
				return
			}
		}
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(url, contentType, bodyBuf)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if !strings.HasPrefix(resp.Status, "200") {
		return errors.New("resp status error:" + resp.Status), string(resp_body)
	}
	return nil, string(resp_body)
}

func postJson(url string, post []byte) (err error, res string) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(post))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return err, string(body)
}
