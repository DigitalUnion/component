package dudownload

type Config struct {
	Files map[string]FileInfo
}
type FileInfo struct {
	Url    string
	Target string
}
