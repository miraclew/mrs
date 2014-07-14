package api

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type RegisterResponse struct {
	Uid int64
}

type TokenResponse struct {
	Uid   int64
	Token string
}
