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
