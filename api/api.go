package api

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type RegisterResponse struct {
	Uid   int64
	Token string
}

type TokenResponse struct {
	Uid   int64
	Token string
}

func response(code int, data interface{}) *Response {
	r := &Response{
		code,
		data,
	}

	return r
}
