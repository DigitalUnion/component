package durpcx

// Req : 基础接口
type Req interface {
	GetClientIp() string
}

// Res : 基础接口
type Res interface {
	SetError(err error)
}
