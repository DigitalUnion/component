package dudingtalk

import (
	"fmt"
	"net/http"
)

// SendMarkdownMessageForString 发送 markdown 消息
func SendMarkdownMessageForString(url, msg string) error {
	body := fmt.Sprintf(markdownCommonFormat, msg)
	resp, status := SendPost(url, []byte(body))
	if status != http.StatusOK {
		return fmt.Errorf("send markdown message failed, status: %d, resp: %s", status, resp)
	}
	return nil
}

// SendMarkdownMessageForDingMap 发送 markdown 消息
func SendMarkdownMessageForDingMap(url string, dm *Dingmap) error {
	if dm == nil {
		return fmt.Errorf("send markdown message failed, dm is nil")
	}
	textList := dm.Slice()
	text := ""
	for _, t := range textList {
		text = text + "  \n  " + t
	}
	return SendMarkdownMessageForString(url, text)
}

// SendTextMessage 发送 text 消息
func SendTextMessage(url, msg string) error {
	body := fmt.Sprintf(textCommonFormat, msg)
	resp, status := SendPost(url, []byte(body))
	if status != http.StatusOK {
		return fmt.Errorf("send text message failed, status: %d, resp: %s", status, resp)
	}
	return nil
}

// SendTableMeaasge 发送 table 消息
func SendTableMeaasge(url string, tm *TableMessage) error {
	if tm == nil {
		return fmt.Errorf("send table message failed, tm is nil")
	}
	return SendMarkdownMessageForString(url, tm.string())
}
