package main

import (
	"fmt"
	"git.du.com/cloud/du_component/dudingtalk"
)

//发送钉钉消息 钉钉机器人需要添加关键字 .
func main() {
	var err error
	url := "https://oapi.dingtalk.com/robot/send?access_token=d301739dcc90bec756ed1da3c759def56ccacc8e9b057d7f898101259d532277"
	//发送表格消息
	tm := dudingtalk.NewTableMessage()
	tm.AddTitle(dudingtalk.H1, "哈哈哈")
	tm.AddHeader("姓名", "年龄", "性别")
	tm.AddRaw("张三", "18", "男")
	tm.AddRaw("李四", "19", "女")
	err = dudingtalk.SendTableMeaasge(url, tm)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//发送markdown消息 格式化
	dm := dudingtalk.NewDingmap()
	dm.Set("哈哈哈1", dudingtalk.H1)
	dm.Set("哈哈哈2", dudingtalk.H2)
	dm.Set("哈哈哈3", dudingtalk.RED)
	dm.Set("哈哈哈4", dudingtalk.GREEN)
	dm.Set("哈哈哈5", dudingtalk.N)
	err = dudingtalk.SendMarkdownMessageForDingMap(url, dm)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//发送markdown消息 自定义
	str := "### aaabbbb \n cdvdvdd"
	err = dudingtalk.SendMarkdownMessageForString(url, str)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//发送文本消息
	err = dudingtalk.SendTextMessage(url, "哈哈哈")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
