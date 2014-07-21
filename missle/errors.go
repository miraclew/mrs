package missle

import (
	"fmt"
)

const (
	ERR_INVALID_ARGS     = 100
	ERR_INVALID_STATE    = 101
	ERR_INVALID_POSITION = 102
	ERR_DATA_NOT_FOUND   = 103

	ERR_INVALID_CREDENTIAL = 201
	ERR_INVALID_TOKEN      = 202
)

var ErrMessageMap = map[int]string{
	ERR_INVALID_ARGS:     "invalid args: %s",
	ERR_INVALID_STATE:    "invalid state: %s",
	ERR_INVALID_POSITION: "invalid position: %f,%f",

	ERR_DATA_NOT_FOUND: "data not found: %s",
	ERR_INVALID_TOKEN:  "invalid token: %s",
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
