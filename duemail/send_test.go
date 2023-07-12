package duemail

import (
	"fmt"
	"testing"
)

func TestEmail_UploadsFiles(t *testing.T) {
	a := Email{}

	err := a.UploadsFiles([]string{
		"/Users/likai/Desktop/test.csv",
		//"/Users/likai/Desktop/test1.csv",
	})

	fmt.Println(err)
}

func TestEmail_SendEmail(t *testing.T) {
	a := Email{Result: func(res string, err error) {
		fmt.Printf("send email err:%s  res:%s\n", err, res)
	}}

	a.SendEmail(EmailNormalParams{
		// 多个以逗号分割
		To:      "likai@shuzilm.cn",
		Subject: "test subject",
		Content: "test content",
		// 多个以逗号分割
		//Cc:      "likai@shuzilm.cn",
		FilesPath: []string{
			//"/Users/likai/Desktop/test.csv",
			//"/Users/likai/Desktop/test1.csv",
		},
	})
}

func TestMakeTableHtml(t *testing.T) {
	data := MakeTableHtml(Table{
		Header: []string{"客户", "请求量"},
		Content: [][]string{
			[]string{"test1", "0"},
			[]string{"test2", "1"},
			[]string{"test3", "2"},
		},
	})

	a := Email{Result: func(res string, err error) {
		fmt.Printf("send email err:%s  res:%s\n", err, res)
	}}

	a.SendEmail(EmailNormalParams{
		// 多个以逗号分割
		To:      "likai@shuzilm.cn",
		Subject: "test subject",
		Content: data,
		// 多个以逗号分割
		//Cc:      "xunjx@shuzilm.cn",
		FilesPath: []string{
			//"/Users/likai/Desktop/test.csv",
			"/data/202204.csv",
		},
	})
}
