package dulogv2

const (
	defaultModule = "default"
	logSplitChar  = "-"
	tcpType       = "tcp"
)

const (
	accessLogFormat = "[%s]\t%s\t%s\t%s\t%s\t%d\t%s\tcost:%s"
)

const (
	ALL = iota
	DEBUG
	INFO
	ERROR
	FATAL
	OFF
)

const (
	levleFormat     = "[%s|%s] %s"
	printTimeFormat = "2006-01-02 15:04:05.000"
)

// local
const (
	locDefaultDir     = "./logs"
	locDefaultRotate  = "0 0 * * * *" //每小时
	locDefaultMaxAge  = 1             //1天
	locDefaultMaxSize = 10240         //10G
)

// timeOut
const (
	maxSize = 1000
	timeOut = 100 //ms
)
