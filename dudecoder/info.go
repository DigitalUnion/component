package dudecoder

type DictInfo struct {
	Field string
	Path  string
	Plan  int
	Anon  bool
}

var infoMap = make(map[int]map[string]*DictInfo)
