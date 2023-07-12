package duconfig

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

var (
	osType string
	path   string
	// ['Redis', 'Mysql', 'Hbase','OSS', 'Consul', 'Kafka', 'Mongo', 'TableStore', 'Influx','Rpcx','DuLog','Customize']
	nameMap = map[string]string{
		"duredis.RedisConfig":    "Redis",
		"dumysql.Config":         "Mysql",
		"duhbase.HbaseConfig":    "Hbase",
		"dukafka.ConsumerConfig": "Kafka",
		"dukafka.ProducerConfig": "Kafka",
		"dumongo.config":         "Mongo",
		"duinflux.InfluxConfig":  "Influx",
		"durpcx.Config":          "Rpcx",
		"custom_log.ConfigLog":   "DuLog",
		"OSS":                    "OSS",
		"Consul":                 "Consul",
		"TableStore":             "TableStore",
		"Customize":              "Customize",
	}
)

const WINDOWS = "windows"

func init() {
	osType = runtime.GOOS
	if os.IsPathSeparator('\\') { //前边的判断是否是系统的分隔符
		path = "\\"
	} else {
		path = "/"
	}
}

func getEnv(key string) (string, error) {
	if osType == WINDOWS {
		return getWindowsEnv(key)
	} else {
		return getLinuxEnv(key), nil
	}
}

type PowerShell struct {
	powerShell string
}

func New() *PowerShell {
	ps, _ := exec.LookPath("powershell.exe")
	return &PowerShell{
		powerShell: ps,
	}
}

func (p *PowerShell) Execute(args ...string) (stdOut string, stdErr string, err error) {
	args = append([]string{"-NoProfile", "-NonInteractive"}, args...)
	cmd := exec.Command(p.powerShell, args...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	stdOut, stdErr = stdout.String(), stderr.String()
	return
}
func getWindowsEnv(key string) (string, error) {
	posh := New()
	stdout, _, err := posh.Execute(fmt.Sprintf(`[environment]::GetEnvironmentvariable("%s", "Machine")`, key))
	return strings.TrimSpace(stdout), err
}

func getLinuxEnv(key string) string {
	return os.Getenv(key)
}

func getFileName(cacheKey string, cacheDir string) string {
	return cacheDir + string(os.PathSeparator) + cacheKey
}

func readConfigFromFile(cacheKey string, cacheDir string) (string, error) {
	fileName := getFileName(cacheKey, cacheDir)
	b, err := os.ReadFile(fileName)
	if err != nil {
		return "", errors.New(fmt.Sprintf("failed to read config cache file:%s,err:%+v ", fileName, err))
	}
	return string(b), nil
}

func mkdirIfNecessary(createDir string) (err error) {
	s := strings.Split(createDir, path)
	startIndex := 0
	dir := ""
	if s[0] == "" {
		startIndex = 1
	} else {
		dir, _ = os.Getwd() //当前的目录
	}
	for i := startIndex; i < len(s); i++ {
		var d string
		if osType == WINDOWS && filepath.IsAbs(createDir) {
			d = strings.Join(s[startIndex:i+1], path)
		} else {
			d = dir + path + strings.Join(s[startIndex:i+1], path)
		}
		if _, e := os.Stat(d); os.IsNotExist(e) {
			err = os.Mkdir(d, os.ModePerm) //在当前目录下生成md目录
			if err != nil {
				break
			}
		}
	}

	return err
}

func getCurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Println("can not get current path")
	}
	return dir
}

func writeConfigToFile(cacheKey string, cacheDir string, content string) {
	mkdirIfNecessary(cacheDir)
	fileName := getFileName(cacheKey, cacheDir)
	if len(content) == 0 {
		// delete config snapshot
		if err := os.Remove(fileName); err != nil {
			if err == os.ErrNotExist {
				return
			}
			log.Printf("failed to delete config file,err:%+v,cache:%s,value:%s", err, fileName, content)
		}
		return
	}
	err := os.WriteFile(fileName, []byte(content), 0666)
	if err != nil {
		log.Printf("failed to write config file,err:%+v,cache:%s,value:%s", err, fileName, content)
	}
}

type Fields struct {
	YamlTag string `json:"yaml_tag"`
	Desc    string `json:"desc"`
	Type    string `json:"type"`
}

func GetFields(obj interface{}) []Fields {
	// 取 Value
	var ret []Fields
	v := reflect.ValueOf(obj)
	reflectType := v.Type()
	// 解析字段
	for i := 0; i < v.NumField(); i++ {

		field := reflectType.Field(i)
		tag := field.Tag
		if field.Anonymous {
			subRet := GetFields(v.Field(i).Interface())
			ret = append(ret, subRet...)
			continue
		}
		yamlTag := tag.Get("yaml")
		desc := tag.Get("desc")
		if yamlTag == "" {
			yamlTag = strings.ToLower(field.Name)
		}
		ret = append(ret, Fields{YamlTag: yamlTag, Desc: desc, Type: getType(field.Type.String())})
	}

	return ret
}

func getType(name string) string {
	if v, ok := nameMap[strings.TrimLeft(name, "*")]; ok {
		return v
	}
	return "Customize"
}
