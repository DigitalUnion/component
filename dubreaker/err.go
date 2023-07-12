package dubreaker

import (
	"errors"
)

var (
	baseErr               = errors.New("sentinel")
	UnknownErr            = &SentinelErr{"Unknown", baseErr}
	FlowControlErr        = &SentinelErr{"FlowControl", baseErr}
	BlockTypeIsolationErr = &SentinelErr{"BlockTypeIsolation", baseErr}
	CircuitBreakingErr    = &SentinelErr{"CircuitBreaking", baseErr}
	SystemErr             = &SentinelErr{"System", baseErr}
	HotSpotParamFlowErr   = &SentinelErr{"HotSpotParamFlow", baseErr}
)

type SentinelErr struct {
	Type string
	Err  error
}

func (e *SentinelErr) Error() string { return e.Type + e.Err.Error() }

func (e *SentinelErr) Unwrap() error { return e.Err }

func getError(blockMsg string) error {
	switch blockMsg {
	case "Unknown":
		return UnknownErr
	case "FlowControl":
		return FlowControlErr
	case "BlockTypeIsolation":
		return BlockTypeIsolationErr
	case "CircuitBreaking":
		return CircuitBreakingErr
	case "System":
		return SystemErr
	case "HotSpotParamFlow":
		return HotSpotParamFlowErr
	default:
		return &SentinelErr{blockMsg, baseErr}
	}
}
