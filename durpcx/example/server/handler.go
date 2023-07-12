package server

import (
	"context"
	"errors"
	"git.du.com/cloud/du_component/durpcx/example/common"
)

type ExampleHandler int

func (p ExampleHandler) Hello(_ context.Context, req *common.ExampleReq, res *common.ExampleRes) error {
	if len(req.Data) == 0 {
		return errors.New("row data not exists")
	}
	res.Data = []byte("Hello " + string(req.Data))
	return nil
}
