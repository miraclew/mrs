package util

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)



func ApiResonpse(w http.ResponseWriter, code int, message string, data interface{}) {
	response, err := json.Marshal(struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{
		code,
		message,
		data,
	})

	if err != nil {
		response = []byte(fmt.Sprintf(`{"code":500,"message":"%s","data":null}`, err.Error()))
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Content-Length", strconv.Itoa(len(response)))
	w.WriteHeader(code)
	w.Write(response)
}
