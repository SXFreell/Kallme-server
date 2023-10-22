package model

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRsp struct {
	RequstId string `json:"request_id"`
	Code     int    `json:"code"`
	Message  string `json:"message"`
}

type LinkWSReq struct {
	Token string `json:"token"`
}

type SendMessageReq struct {
	Token string      `json:"token"`
	Msg   interface{} `json:"msg"`
}
