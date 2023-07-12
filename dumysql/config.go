package dumysql

type Config struct {
	Address      string
	Username     string
	Password     string
	Database     string
	Charset      string
	MaxOpenConns int
	MaxIdleConns int
}
