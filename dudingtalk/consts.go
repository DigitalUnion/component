package dudingtalk

const dingTalkURL = "https://oapi.dingtalk.com/robot/send?"
const dtmdFormat = "[%s](dtmd://dingtalkclient/sendMessage?content=%s)"
const formatSpliter = "$$"

const (
	H1     MarkType = "h1"
	H2     MarkType = "h2"
	H3     MarkType = "h3"
	H4     MarkType = "h4"
	H5     MarkType = "h5"
	H6     MarkType = "h6"
	RED    MarkType = "red"
	BLUE   MarkType = "blue"
	GREEN  MarkType = "green"
	GOLD   MarkType = "gold"
	YELLOW MarkType = "yellow"
	N      MarkType = ""
)

var hMap = map[MarkType]string{
	H1:    "# %s",
	H2:    "## %s",
	H3:    "### %s",
	H4:    "#### %s",
	H5:    "##### %s",
	H6:    "###### %s",
	RED:   "<font color=#ff0000>%s</font>",
	BLUE:  "<font color=#1E90FF>%s</font>",
	GREEN: "<font color=#00CD66>%s</font>",
	GOLD:  "<font color=#FFD700>%s</font>",
}

const (
	// markdown  markdown 消息格式
	markdownCommonFormat = `{"msgtype": "markdown","markdown": {"title":".","text": "%s"}}`
	// text  text 消息格式
	textCommonFormat = `{"msgtype": "text","text": {"content": "%s."}}`
)

const (
	imageServer = "http://172.17.129.178:8451/create_image"
	imageFormat = "![](%s)"
)
