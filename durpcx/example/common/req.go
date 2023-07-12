package common

type ExampleReq struct {
	Ip   string
	Data []byte
}

func (p ExampleReq) GetClientIp() string {
	return p.Ip
}
