package handler

import (
	"net/http"

	"Hermes/api/internal/logic"
	"Hermes/api/internal/svc"
	"Hermes/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func HermesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HermesReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewHermesLogic(r.Context(), svcCtx)
		resp, err := l.Hermes(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
