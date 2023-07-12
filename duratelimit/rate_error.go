package duratelimit

import "errors"

var (
	nameEmpty   = errors.New("name is empty")
	rateEmpty   = errors.New("rate is empty")
	allowEmpty  = errors.New("allow is empty")
	allowExceed = errors.New("allow exceed the limit")
)
