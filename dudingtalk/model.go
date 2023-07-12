package dudingtalk

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"net/http"
	"strings"
)

type MarkType string

func (mt MarkType) String() string {
	return hMap[mt]
}

// 有序map
type Dingmap struct {
	m map[string]MarkType
	l []string
}

func NewDingmap() *Dingmap {
	return &Dingmap{m: make(map[string]MarkType), l: make([]string, 0, 0)}
}

func (d *Dingmap) Set(val string, t MarkType) *Dingmap {
	d.l = append(d.l, val)
	d.m[val] = t
	return d
}

func (d *Dingmap) Remove(val string) {
	if _, ok := d.m[val]; ok {
		for i, v := range d.l {
			if v == val {
				d.l = append(d.l[:i], d.l[i+1:]...)
				break
			}
		}
		delete(d.m, val)
	}
}

func (d *Dingmap) Slice() []string {
	resList := make([]string, 0, len(d.l))
	for _, val := range d.l {
		content := d.formatVal(val, d.m[val])
		resList = append(resList, content)
	}
	return resList
}

func (d *Dingmap) formatVal(val string, t MarkType) (res string) {
	var ok bool
	if res, ok = hMap[t]; ok {
		//vl := strings.Split(val, formatSpliter)
		//if len(vl) == 3 {
		//	res = fmt.Sprintf(res, vl[1])
		//	res = vl[0] + res + vl[2]
		//} else {
		res = fmt.Sprintf(res, val)
		//}
	} else {
		res = val
	}
	if !strings.HasPrefix(res, "- ") && !strings.HasPrefix(res, "#") {
		res = "- " + res
	}
	return
}

type TableMessage struct {
	Title    string     `json:"title,omitempty"`
	Header   []string   `json:"header,omitempty"`
	Data     [][]string `json:"data,omitempty"`
	Span     []int      `json:"span,omitempty"`
	ImageUrl string     `json:"image_url,omitempty"`
}

func NewTableMessage() *TableMessage {
	return &TableMessage{}
}

// AddTitle 添加表格标题
func (tm *TableMessage) AddTitle(m MarkType, title string) {
	tm.Title = fmt.Sprintf(m.String(), title)
}

// AddHeader 添加表格头
func (tm *TableMessage) AddHeader(cols ...string) {
	tm.Header = append(tm.Header, cols...)
}

// AddRaw 添加表格数据
func (tm *TableMessage) AddRaw(cols ...string) {
	tm.Data = append(tm.Data, cols)
}

// AddSpan 添加字体样式
func (tm *TableMessage) AddSpan(cols ...int) {
	tm.Span = append(tm.Span, cols...)
}

// createImage 生成图片
func (tm *TableMessage) createImage() {
	body, _ := jsoniter.Marshal(tm)
	respData, status := SendPost(imageServer, body)
	if status != http.StatusOK {
		return
	}
	tm.ImageUrl = jsoniter.Get(respData, "data").ToString()
}

func (tm *TableMessage) string() string {
	tm.createImage()
	//拼接字符串
	return tm.Title + "  " + fmt.Sprintf(imageFormat, tm.ImageUrl)
}
