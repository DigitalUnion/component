package common

type ExampleRes struct {
	Code  int
	Error error
	Data  []byte
}

func (p *ExampleRes) GetResCode() int {
	return p.Code
}
func (p *ExampleRes) SetError(err error) {
	p.Code = 400
	p.Error = err
}
