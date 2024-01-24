// Code generated by goctl. DO NOT EDIT.
package types

type RegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisteraResp struct {
	Ok bool   `json:"ok"`
	Id string `json:"id"`
}

type InfoReq struct {
}

type InfoResp struct {
	Id       string  `json:"id"`
	Username string  `json:"username"`
	Balance  float64 `json:"balance"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResp struct {
	Id         string `json:"id"`
	Token      string `json:"token"`
	ExpireTime int64  `json:"expireTime"`
}
