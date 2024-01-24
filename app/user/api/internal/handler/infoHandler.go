package handler

import (
	"google.golang.org/grpc/status"
	"net/http"
	"ticket-server/common/errs"
	"ticket-server/common/httpc"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ticket-server/app/user/api/internal/logic"
	"ticket-server/app/user/api/internal/svc"
	"ticket-server/app/user/api/internal/types"
)

func InfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.InfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpc.Response(r.Context(), w, errs.ParamParseFailed.Code(), errs.ParamParseFailed.Error(), nil)
			return
		}

		l := logic.NewInfoLogic(r.Context(), svcCtx)
		resp, err := l.Info(&req)
		if err != nil {
			if s, ok := status.FromError(err); ok {
				httpc.Response(r.Context(), w, s.Code(), s.Message(), nil)
			} else {
				httpc.Response(r.Context(), w, 1000, err.Error(), nil)
			}
		} else {
			httpc.Success(r.Context(), w, resp)
		}
	}
}
