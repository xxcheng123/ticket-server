package httpc

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc/codes"
	"net/http"
)

/**
*	@Author: xxcheng
*	@Email developer@xxcheng.cn
*	@Date: 2024/1/24
 */

type Resp struct {
	Code codes.Code `json:"code"`
	Msg  string     `json:"msg"`
	Data any        `json:"data,omitempty"`
}

func response(ctx context.Context, w http.ResponseWriter, code codes.Code, msg string, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	data := &Resp{
		Code: code,
		Msg:  msg,
		Data: v,
	}
	s, _ := json.Marshal(data)
	_, _ = w.Write(s)
}

func Success(ctx context.Context, w http.ResponseWriter, v any) {
	response(ctx, w, http.StatusOK, "success", v)
}
func Response(ctx context.Context, w http.ResponseWriter, code codes.Code, msg string, v any) {
	response(ctx, w, code, msg, v)
}
