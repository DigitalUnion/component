package dudingtalk

import (
	"fmt"
	"testing"
)

func TestSendTableMeaasge(t *testing.T) {
	url := "https://oapi.dingtalk.com/robot/send?access_token=d301739dcc90bec756ed1da3c759def56ccacc8e9b057d7f898101259d532277"
	tm := NewTableMessage()
	tm.AddTitle(H1, "哈哈哈")
	tm.AddHeader("姓名", "年龄", "性别")
	tm.AddRaw("张三", "18", "男")
	tm.AddRaw("李四", "19", "女")
	tm.AddSpan(5, 5, 5)
	err := SendTableMeaasge(url, tm)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func TestSendMarkdownMessageForDingMap(t *testing.T) {
	url := "https://oapi.dingtalk.com/robot/send?access_token=d301739dcc90bec756ed1da3c759def56ccacc8e9b057d7f898101259d532277"
	dm := NewDingmap()
	dm.Set("哈哈哈1", H1)
	dm.Set("哈哈哈2", H2)
	dm.Set("哈哈哈3", H3)
	dm.Set("哈哈哈4", H4)
	SendMarkdownMessageForDingMap(url, dm)
}
