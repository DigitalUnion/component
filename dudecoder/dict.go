package dudecoder

import (
	"errors"
	"fmt"
	"github.com/goccy/go-json"
	"log"
	"reflect"
	"strconv"
	"strings"
)

func parseInterface(verCode interface{}) string {
	switch verCode.(type) {
	case string:
		return verCode.(string)
	case int:
		return strconv.Itoa(verCode.(int))
	case float64:
		return strconv.FormatFloat(verCode.(float64), 'f', -1, 64)
	default:
		return ""
	}
}

func (p Imap) GetStringByProperty(f []string) string {
	if len(f) == 1 {
		if v, ok := p[f[0]]; ok {
			return parseInterface(v)
		}
		return ""
	}
	m := p
	for i, k := range f {
		if i != len(f)-1 {
			if v, ok := m[k]; ok {
				if tmpV, ok := v.(map[string]interface{}); ok {
					m = tmpV
					continue
				}
			}
			return ""
		}
		if v, ok := m[k]; ok {
			return parseInterface(v)
		}
	}
	return ""
}

var showUnknownFields = false

type Imap map[string]interface{}

func (p Imap) GetAsMap(k string) Imap {
	if p[k] == nil {
		return nil
	}
	switch p[k].(type) {
	case map[string]interface{}:
		return p[k].(map[string]interface{})
	default:
		return nil
	}
}

func (p Imap) Get(ks ...string) interface{} {
	if len(ks) == 0 {
		return nil
	}
	if len(ks) == 1 {
		k := ks[0]
		if p[k] == nil {
			return nil
		} else {
			return p[k]
		}
	} else {
		t := p
		for i := 0; i < len(ks); i++ {
			if i != len(ks)-1 {
				t = t.GetAsMap(ks[i])
				if t == nil {
					return nil
				}
				continue
			}
			if t[ks[i]] == nil {
				return nil
			}
			return t[ks[i]]
		}
	}
	return nil
}

func (p Imap) GetString(ks ...string) string {
	if len(ks) == 0 {
		return ""
	}
	if len(ks) == 1 {
		k := ks[0]
		if p[k] == nil {
			return ""
		} else {
			return p[k].(string)
		}
	} else {
		t := p
		for i := 0; i < len(ks); i++ {
			if i != len(ks)-1 {
				t = t.GetAsMap(ks[i])
				if t == nil {
					return ""
				}
				continue
			}
			if t[ks[i]] == nil {
				return ""
			}
			return t[ks[i]].(string)
		}
	}
	return ""
}
func (p Imap) GetInt64(ks ...string) int64 {
	if len(ks) == 0 {
		return 0
	}
	if len(ks) == 1 {
		k := ks[0]
		if p[k] == nil {
			return 0
		} else {
			return p[k].(int64)
		}
	} else {
		t := p
		for i := 0; i < len(ks); i++ {
			if i != len(ks)-1 {
				t = t.GetAsMap(ks[i])
				if t == nil {
					return 0
				}
				continue
			}
			if t[ks[i]] == nil {
				return 0
			}
			return t[ks[i]].(int64)
		}
	}
	return 0
}
func (p Imap) GetFloat64(ks ...string) float64 {
	if len(ks) == 0 {
		return 0
	}
	if len(ks) == 1 {
		k := ks[0]
		if p[k] == nil {
			return 0
		} else {
			return p[k].(float64)
		}
	} else {
		t := p
		for i := 0; i < len(ks); i++ {
			if i != len(ks)-1 {
				t = t.GetAsMap(ks[i])
				if t == nil {
					return 0
				}
				continue
			}
			if t[ks[i]] == nil {
				return 0
			}
			return t[ks[i]].(float64)
		}
	}
	return 0
}
func (p Imap) Set(k string, v interface{}) {
	p[k] = v
}
func (p Imap) Del(k string) {
	delete(p, k)
}

func DecodeDNA(bs []byte) (Imap, error) {
	m := make(map[string]interface{}, 100)
	err := json.UnmarshalWithOption(bs, &m, json.DecodeFieldPriorityFirstWin())
	if err != nil {
		return nil, err
	}
	_, isKeyEncode := m["AAA"]
	_, isValEncode := m["BBB"]
	if isKeyEncode {
		var tk int
		if isValEncode {
			timeVal := m["2cO"]
			switch timeVal.(type) {
			case string:
				tk = timeKey(timeVal.(string))
			case float64:
				curtime, ok := timeVal.(float64)
				if !ok {
					return m, errors.New("curtime not found")
				}
				curtimeStr := strconv.FormatFloat(curtime, 'f', 0, 64)
				tk = timeKey(curtimeStr)
			}
		}
		var paths []string
		DecryptMap(paths, m, tk, isValEncode, infoMap[DNA], platformAndroid)
	}
	return m, nil
}
func DecodeDAI(bs []byte) (Imap, error) {
	m := make(map[string]interface{}, 100)
	err := json.UnmarshalWithOption(bs, &m, json.DecodeFieldPriorityFirstWin())
	if err != nil {
		return nil, err
	}
	_, isKeyEncode := m["AAA"]
	_, isValEncode := m["BBB"]
	if isKeyEncode {
		var tk int
		if isValEncode {
			timeVal := m["2cO"]
			switch timeVal.(type) {
			case string:
				tk = timeKey(timeVal.(string))
			case float64:
				curtime, ok := timeVal.(float64)
				if !ok {
					return m, errors.New("curtime not found")
				}
				curtimeStr := strconv.FormatFloat(curtime, 'f', 0, 64)
				tk = timeKey(curtimeStr)
			}
		}
		var paths []string
		DecryptMap(paths, m, tk, isValEncode, infoMap[DAI], platformAndroid)
	}
	return m, nil
}
func DecodeIDAA(bs []byte) (Imap, error) {
	m := make(map[string]interface{}, 100)
	err := json.UnmarshalWithOption(bs, &m, json.DecodeFieldPriorityFirstWin())
	if err != nil {
		return nil, err
	}
	_, isKeyEncode := m["AAA"]
	_, isValEncode := m["BBB"]
	if isKeyEncode {
		var tk int
		if isValEncode {
			timeVal := m["th1"]
			switch timeVal.(type) {
			case string:
				tk = timeKey(timeVal.(string))
			case float64:
				curtime, ok := timeVal.(float64)
				if !ok {
					return m, errors.New("curtime not found")
				}
				curtimeStr := strconv.FormatFloat(curtime, 'f', 0, 64)
				tk = timeKey(curtimeStr)
			}
		}
		var paths []string
		DecryptMap(paths, m, tk, isValEncode, infoMap[IDAA], platformIos)
	}
	return m, nil
}
func DecodeIDNA(bs []byte) (Imap, error) {
	m := make(map[string]interface{}, 100)
	err := json.UnmarshalWithOption(bs, &m, json.DecodeFieldPriorityFirstWin())
	if err != nil {
		return nil, err
	}
	_, isKeyEncode := m["AAA"]
	_, isValEncode := m["BBB"]
	if isKeyEncode {
		var tk int
		if isValEncode {
			timeVal := m["th1"]
			switch timeVal.(type) {
			case string:
				tk = timeKey(timeVal.(string))
			case float64:
				curtime, ok := timeVal.(float64)
				if !ok {
					return m, errors.New("curtime not found")
				}
				curtimeStr := strconv.FormatFloat(curtime, 'f', 0, 64)
				tk = timeKey(curtimeStr)
			}
		}
		var paths []string
		DecryptMap(paths, m, tk, isValEncode, infoMap[IDNA], platformIos)
	}
	return m, nil
}

func DecodeApplet(bs []byte) (Imap, error) {
	m := make(map[string]interface{}, 100)
	err := json.UnmarshalWithOption(bs, &m, json.DecodeFieldPriorityFirstWin())
	if err != nil {
		return nil, err
	}
	//_, isKeyEncode := m["AAA"]
	_, isValEncode := m["BBB"]
	if true {
		//if isKeyEncode {
		var tk int
		if isValEncode {
			timeVal := m["2cO"]
			switch timeVal.(type) {
			case string:
				tk = timeKey(timeVal.(string))
			case float64:
				curtime, ok := timeVal.(float64)
				if !ok {
					return m, errors.New("curtime not found")
				}
				curtimeStr := strconv.FormatFloat(curtime, 'f', 0, 64)
				tk = timeKey(curtimeStr)
			}
		}
		var paths []string
		DecryptMap(paths, m, tk, isValEncode, infoMap[Applet], platformApplet)
	}
	return m, nil
}
func DecodeDAA(bs []byte) (Imap, error) {
	m := make(map[string]interface{}, 100)
	err := json.UnmarshalWithOption(bs, &m, json.DecodeFieldPriorityFirstWin())
	if err != nil {
		return nil, err
	}
	_, isKeyEncode := m["AAA"]
	_, isValEncode := m["BBB"]
	if isKeyEncode {
		var tk int
		if isValEncode {
			timeVal := m["2cO"]
			switch timeVal.(type) {
			case string:
				tk = timeKey(timeVal.(string))
			case float64:
				curtime, ok := timeVal.(float64)
				if !ok {
					return m, errors.New("curtime not found")
				}
				curtimeStr := strconv.FormatFloat(curtime, 'f', 0, 64)
				tk = timeKey(curtimeStr)
			}
		}
		var paths []string
		DecryptMap(paths, m, tk, isValEncode, infoMap[DAA], platformAndroid)
	}
	return m, nil
}
func DecryptMap(paths []string, m map[string]interface{}, tk int, isValEncode bool, infoMap map[string]*DictInfo, platform int) map[string]interface{} {
	tm := make(map[string]interface{})
	var unknownFields map[string]bool
	if showUnknownFields {
		unknownFields = make(map[string]bool)
	}
	for k, v := range m {
		if len(k) < 5 {
			if IsNumber(k) {
				k = "0"
			}
		}
		var info *DictInfo
		if len(paths) == 0 {
			info = infoMap[k]
			if showUnknownFields {
				if info == nil {
					unknownFields[k] = true
				}
			}
		} else {
			ps := append(paths, k)
			path := strings.Join(ps, "/")
			info = infoMap[path]
			if showUnknownFields {
				if info == nil {
					if !strings.HasSuffix(path, "/0") {
						unknownFields[k] = true
					}
				}
			}
		}
		switch v.(type) {
		case string:
			if info != nil {
				if isValEncode {
					// 字段值解析
					switch info.Plan {
					case 1:
						newBytes, _ := decryptPlan1(interfaceToBytes(v))
						v = Bytes2String(newBytes)
					case 2:
						if info.Anon {
							md5Key := getMd5Key(k)
							if m[md5Key] == nil {
								continue
							}
							if showUnknownFields {
								delete(unknownFields, md5Key)
							}
							md5Val := m[md5Key].(string)
							newBytes, _ := decryptPlan2Anon(interfaceToBytes(v), md5Val, tk, platform)
							v = Bytes2String(newBytes)
							delete(m, md5Key)
						} else {
							if len(v.(string)) <= 1024 {
								newBytes, _ := decryptPlan2(interfaceToBytes(v), tk, platform)
								v = Bytes2String(newBytes)
							}
						}
					}
				}
			}
		case float64:
		case bool:
		case []interface{}:
			if info != nil {
				if info.Plan != 0 {
					vals := v.([]interface{})
					newVals := make([]string, len(vals))
					for i, e := range vals {
						switch e.(type) {
						case []byte:
							newVals[i] = DecryptString(e.([]byte), tk, info.Plan, platform)
						case string:
							newVals[i] = DecryptString(String2Bytes(e.(string)), tk, info.Plan, platform)
						default:
							log.Println("unhandled []interface type:", reflect.TypeOf(v))
						}
					}
					v = newVals
				}
			}
		case map[string]interface{}:
			v = DecryptMap(append(paths, k), v.(map[string]interface{}), tk, isValEncode, infoMap, platform)
		default:
			log.Println("unhandled type:", reflect.TypeOf(v))
		}
		// 字段名转换
		if info != nil {
			tm[info.Field] = v
			delete(m, k)
		}
	}
	for k, v := range tm {
		m[k] = v
	}
	if showUnknownFields {
		for k := range unknownFields {
			fmt.Println("unknown field:", k)
		}
	}
	return m
}
func DecryptString(bs []byte, tk int, plan int, platform int) string {
	if len(bs) == 0 {
		return ""
	}
	switch plan {
	case 1:
		v, _ := decryptPlan1(bs)
		return Bytes2String(v)
	case 2:
		newBytes, _ := decryptPlan2(bs, tk, platform)
		return Bytes2String(newBytes)
	}
	return Bytes2String(bs)
}

func IsNumber(s string) bool {
	for _, e := range s {
		if e < 48 || e > 57 {
			return false
		}
	}
	return true
}
