package {{.PkgName}}

import (
	"google.golang.org/grpc/status"
    "net/http"
    "ticket-server/common/httpc"

	"github.com/zeromicro/go-zero/rest/httpx"
	{{.ImportPackages}}
)

func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}
		if err := httpx.Parse(r, &req); err != nil {
			httpc.Response(r.Context(), w, errs.ParamParseFailed.Code(), errs.ParamParseFailed.Error(), nil)
            return
		}

		{{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
		if err != nil {
			if s, ok := status.FromError(err); ok {
                httpc.Response(r.Context(), w, s.Code(), s.Message(), nil)
            } else {
                httpc.Response(r.Context(), w, 1000, err.Error(), nil)
            }
		} else {
			{{if .HasResp}}httpc.Success(r.Context(), w, resp){{else}}httpc.Success(r.Context(), w, nil){{end}}
		}
	}
}
