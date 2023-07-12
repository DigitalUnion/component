package dulogv2

import (
	"fmt"
	"git.du.com/cloud/du_component/dulogv2/du_syslog"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/silenceWe/loggo"
	"log"
	"path"
	"strings"
	"sync"
	"time"
)

type Logger struct {
	ws       []*slog //远程日志
	badWs    chan *slog
	facility du_syslog.Priority
	addrs    []string
	tag      string
	length   int
	total    int
	ck       bool
	l        sync.RWMutex
	isLoc    bool          //是否本地日志
	localW   *loggo.Logger //本地日志
	isEs     bool          //是否写入es
}

func NewLogger(cfg Config) *Logger {
	//配置校验
	cfg.valid()

	if cfg.IsLoc {
		//本地日志
		logger := &Logger{
			isLoc:  true,
			localW: &loggo.Logger{},
		}
		filename := path.Join(cfg.LocDir, cfg.Appid, cfg.Module, cfg.Appid+"_"+cfg.Module+".log")
		logger.localW.SetWriter(&loggo.FileWriter{
			FileName:   filename,
			MaxSize:    cfg.LocMaxSize,
			MaxAge:     cfg.LocMaxAge,
			Compress:   cfg.LocCompress,
			RotateCron: cfg.LocRotate,
			Prefix:     cfg.Appid + "_" + cfg.Module,
		})
		log.Printf("local_log init is success: appid[%s], module[%s], filepath[%s]\n", cfg.Appid, cfg.Module, filename)
		return logger
	} else {
		//远程日志
		logger := &Logger{
			facility: du_syslog.LOG_LOCAL0,
			tag:      cfg.Appid + logSplitChar + cfg.Module,
			addrs:    cfg.Addrs,
			badWs:    make(chan *slog, 100),
			isEs:     cfg.IsEs,
		}
		logger.connect()
		if cfg.IsCheck {
			logger.ck = true
			go logger.check()
		}
		if logger.length == 0 {
			log.Printf("remote_log init is error: appid[%s], module[%s]\n", cfg.Appid, cfg.Module)
		} else {
			log.Printf("remote_log init is success: appid[%s], module[%s]\n", cfg.Appid, cfg.Module)
		}
		return logger
	}
}

func GinDulogMiddleware(logger *Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		if logger != nil {
			logger.PlainTextln(fmt.Sprintf(accessLogFormat, time.Now().Format("2006-01-02 15:04:05"), c.ClientIP(), c.Request.Method, c.Request.URL.Path, c.Request.Proto, c.Writer.Status(), c.Request.UserAgent(), time.Since(start)))
		}
	}
}

func (l *Logger) connect() {
	for _, addr := range l.addrs {
		w, err := du_syslog.Dial(tcpType, addr, l.facility, l.tag)
		sl := &slog{
			w:    w,
			addr: addr,
		}
		if err != nil {
			log.Println("dulog err:", err.Error())
			l.badWs <- sl
		} else {
			l.ws = append(l.ws, sl)
			l.length = l.length + 1
		}
	}
}

func (l *Logger) removeBadw(index int, m string) {
	if index >= l.length || index < 0 {
		return
	}
	l.ws[index].m = m
	l.badWs <- l.ws[index]
	l.ws = append(l.ws[:index], l.ws[index+1:]...)
	l.length = l.length - 1
}

func (l *Logger) addw(w *slog) {
	l.l.Lock()
	defer l.l.Unlock()
	l.ws = append(l.ws, w)
	l.length = l.length + 1
}

func (l *Logger) write(m string) {
	if l.length == 0 {
		return
	}
	l.total++
	index := l.total % l.length
	if l.isEs {
		if validJson(m) {
			temp := msg2Json(m)
			if temp != "" {
				l.ws[index].w.Debug(temp)
			}

		}
	}
	l.ws[index].w.Info(m)
}

func (l *Logger) writeAndCheck(m string) {
	l.l.Lock()
	defer l.l.Unlock()
	if l.length == 0 {
		return
	}
	l.total++
	index := l.total % l.length
	if l.isEs {
		if validJson(m) {
			temp := msg2Json(m)
			if temp != "" {
				l.ws[index].w.Debug(temp)
			}
		}
	}
	err := l.ws[index].w.Info(m)
	if err != nil {
		l.removeBadw(index, m)
	}
}

func (l *Logger) check() {
	for {
		bw := <-l.badWs
		if bw.w == nil {
			var err error
			bw.w, err = du_syslog.Dial(tcpType, bw.addr, l.facility, l.tag)
			if err != nil {
				l.badWs <- bw
			} else {
				l.addw(bw)
			}
		} else {
			if l.isEs {
				if validJson(bw.m) {
					temp := msg2Json(bw.m)
					if temp != "" {
						bw.w.Debug(temp)
					}
				}
			}
			err := bw.w.Info(bw.m)
			if err != nil {
				l.badWs <- bw
			} else {
				l.addw(bw)
			}
		}
		time.Sleep(time.Second)
	}
}

// PlainTextln 文本日志
func (l *Logger) PlainTextln(s ...string) {
	m := strings.Join(s, ",")
	m = removeWrap(m)
	if l.isLoc {
		l.localW.PlainTextln(m)
		return
	}
	if l.ck {
		l.writeAndCheck(m)
	} else {
		l.write(m)
	}
}

// PlainTextfn 文本日志
func (l *Logger) PlainTextfn(format string, v ...interface{}) {
	m := fmt.Sprintf(format, v...)
	l.PlainTextln(m)
}

func (l *Logger) Infoln(s ...string) {
	m := strings.Join(s, ",")
	m = fmt.Sprintf(levleFormat, getTime(), getLevelStr(INFO), m)
	l.PlainTextln(m)
}

func (l *Logger) Infofn(format string, v ...interface{}) {
	m := fmt.Sprintf(levleFormat, getTime(), getLevelStr(INFO), fmt.Sprintf(format, v...))
	l.PlainTextln(m)
}

func (l *Logger) Debugln(s ...string) {
	m := strings.Join(s, ",")
	m = fmt.Sprintf(levleFormat, getTime(), getLevelStr(DEBUG), m)
	l.PlainTextln(m)
}

func (l *Logger) Debugfn(format string, v ...interface{}) {
	m := fmt.Sprintf(levleFormat, getTime(), getLevelStr(DEBUG), fmt.Sprintf(format, v...))
	l.PlainTextln(m)
}

func (l *Logger) Errorln(s ...string) {
	m := strings.Join(s, ",")
	m = fmt.Sprintf(levleFormat, getTime(), getLevelStr(ERROR), m)
	l.PlainTextln(m)
}

func (l *Logger) Errorfn(format string, v ...interface{}) {
	m := fmt.Sprintf(levleFormat, getTime(), getLevelStr(ERROR), fmt.Sprintf(format, v...))
	l.PlainTextln(m)
}

func (l *Logger) Fatalln(s ...string) {
	m := strings.Join(s, ",")
	m = fmt.Sprintf(levleFormat, getTime(), getLevelStr(FATAL), m)
	l.PlainTextln(m)
	panic(m)
}

func (l *Logger) Fatalfn(format string, v ...interface{}) {
	m := fmt.Sprintf(levleFormat, getTime(), getLevelStr(FATAL), fmt.Sprintf(format, v...))
	l.PlainTextln(m)
	panic(m)
}

func (l *Logger) Write(p []byte) (n int, err error) {
	l.PlainTextln(string(p))
	return len(p), nil
}

func removeWrap(m string) string {
	if strings.Contains(m, "\n") {
		m = strings.ReplaceAll(m, "\n", "")
	}
	if strings.Contains(m, "\r") {
		m = strings.ReplaceAll(m, "\r", "")
	}
	return m
}

type slog struct {
	w    *du_syslog.Writer
	addr string
	m    string
}

func getLevelStr(level int) string {
	switch level {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO "
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	}
	return "UNKNOWN"
}

func getTime() string {
	return time.Now().Format(printTimeFormat)
}

func validJson(m string) bool {
	return jsoniter.Valid([]byte(m))
}

func msg2Json(m string) string {
	now := time.Now().Format(time.RFC3339)
	res := make(map[string]interface{})
	err := jsoniter.Unmarshal([]byte(m), &res)
	if err != nil {
		return ""
	}
	replaceKey(res)
	m, _ = jsoniter.MarshalToString(res)
	return fmt.Sprintf(`{"@timestamp":"%s","msg":%s}`, now, m)
}

// 使用递归替换map[string]interface{}所有key中的'.'为'_'
func replaceKey(m map[string]interface{}) {
	for key, val := range m {
		if strings.Contains(key, ".") {
			newKey := strings.Replace(key, ".", "_", -1)
			m[newKey] = val
			delete(m, key)
		}
		if val, ok := val.(map[string]interface{}); ok {
			replaceKey(val)
		}
	}
}
