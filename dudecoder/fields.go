/**
 * @Author: chenwei
 * @Description:
 * @Date: 2022/11/10 11:21
 */

package dudecoder

import (
	"fmt"
	"github.com/goccy/go-json"
	"io/ioutil"
	"log"
	"strings"
)

type FieldListRes struct {
	Code  int         `json:"code"`
	Data  []FieldItem `json:"data"`
	Count int         `json:"count"`
}
type FieldItem struct {
	Id              string      `json:"id"`
	BoardId         int         `json:"board_id"`
	BoardName       string      `json:"board_name"`
	MessageTypeId   string      `json:"message_type_id"`
	MessageTypeName string      `json:"message_type_name"`
	ConfuseName     string      `json:"confuse_name"`
	SrcName         string      `json:"src_name"`
	StandardName    string      `json:"standard_name"`
	CloudName       string      `json:"cloud_name"`
	BigdataName     string      `json:"bigdata_name"`
	WebName         string      `json:"web_name"`
	Type            string      `json:"type"`
	Demonstration   string      `json:"demonstration"`
	SdkStartVersion string      `json:"sdk_start_version"`
	SdkEndVersion   string      `json:"sdk_end_version"`
	ApiStartVersion string      `json:"api_start_version"`
	ApiEndVersion   string      `json:"api_end_version"`
	MessageEncrypt  string      `json:"message_encrypt"`
	JsonLevel       string      `json:"json_level"`
	Sense           string      `json:"sense"`
	Description     string      `json:"description"`
	Deleted         int         `json:"deleted"`
	IsBigdataNeeded int         `json:"is_bigdata_needed"`
	Reason          interface{} `json:"reason"`
	Time            interface{} `json:"time"`
	UserId          interface{} `json:"user_id"`
	CreateTime      string      `json:"create_time"`
	UpdateTime      string      `json:"update_time"`
}

// GetFieldList : 获取字段列表
func GetFieldList(filePath string, biz int) {
	bs, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}
	parseFields(bs, biz)
}

// GetFieldListFromApi : 从接口获取最新的字段列表，仅用于字段文件更新程序，其他程序直接使用GetFieldList方法即可
func GetFieldListFromApi(addr string, channelTp, messageTypeId int) ([]byte, error) {
	page := 1
	var itemList []FieldItem
	for {
		bs, err := HttpPost(addr, nil, []byte(fmt.Sprintf(`{"page": %d,"limit": 99999,"channel_tp": %d,"message_type_id": %d}`, page, channelTp, messageTypeId)))
		page++
		if err != nil {
			return nil, err
		}
		r := FieldListRes{}
		err = json.Unmarshal(bs, &r)
		if err != nil {
			return nil, err
		}
		itemList = append(itemList, r.Data...)
		if len(itemList) >= r.Count {
			break
		}
	}
	allRes := FieldListRes{
		Code:  0,
		Data:  itemList,
		Count: len(itemList),
	}
	return json.Marshal(allRes)
}
func parseFields(bs []byte, biz int) error {
	r := FieldListRes{}
	err := json.Unmarshal(bs, &r)
	if err != nil {
		return err
	}
	nameMap := make(map[string]string)
	for _, e := range r.Data {
		if len(e.ConfuseName) < 2 {
			continue
		}
		nameMap[e.JsonLevel+"/"+e.SrcName] = e.ConfuseName
	}
	for _, e := range r.Data {
		if len(e.ConfuseName) < 2 {
			continue
		}
		var paths []string
		if len(e.JsonLevel) != 0 {
			levels := strings.Split(e.JsonLevel, "/")
			for i, p := range levels {
				var cname string
				if i == 0 {
					cname = nameMap["/"+p]
				} else {
					key := strings.Join(levels[:i], "/")
					cname = nameMap[key+"/"+p]
				}
				if cname == "" {
					if p == "0" {
						cname = "0"
					} else {
						log.Println("unknown path:", p)
						continue
					}
				}
				paths = append(paths, cname)
			}
		}
		
		paths = append(paths, e.ConfuseName)
		
		di := DictInfo{Field: e.SrcName}
		di.Plan, di.Anon = getPlan(e.MessageEncrypt)
		infoMap[biz][strings.Join(paths, "/")] = &di
	}
	return nil
}
func getPlan(s string) (int, bool) {
	if len(s) == 0 {
		return 0, false
	}
	if strings.HasPrefix(s, "方案一") {
		return 1, false
	}
	if strings.HasPrefix(s, "方案二") {
		return 2, false
	}
	if strings.HasPrefix(s, "方案三") {
		return 2, true
	}
	return 0, false
	
}
