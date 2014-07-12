package missle

import (
	"fmt"
)

const (
	ERR_INVALID_STATE    = 100
	ERR_INVALID_POSITION = 101
)

var ErrMessageMap = map[int]string{
	ERR_INVALID_STATE:    "invalid state: %s",
	ERR_INVALID_POSITION: "invalid position: %f,%f",
}

type MissleErr struct {
	Code    int
	Message string
}

func (e *MissleErr) Error() string {
	return fmt.Sprintf("%d %s", e.Code, e.Message)
}

func NewMissleErr(code int, a ...interface{}) *MissleErr {
	desc := fmt.Sprintf(ErrMessageMap[code], a...)
	return &MissleErr{
		Code:    code,
		Message: desc,
	}
}
