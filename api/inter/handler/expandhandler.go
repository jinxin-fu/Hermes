package handler

import (
	"net/http"

	"Hermes/api/inter/logic"
	"Hermes/api/inter/svc"
	"Hermes/api/inter/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ExpandHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ExpandReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewExpandLogic(r.Context(), svcCtx)
		resp, err := l.Expand(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
