/**
 * Created with IntelliJ goland.
 * @Auther: jinxin
 * @Date: 2023/05/23/9:29
 * @Description:
 */
package handler

import (
	"Hermes/api/inter/logic"
	"Hermes/api/inter/svc"
	"Hermes/pkg/parser"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"net/http/httputil"
)

func QueryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req1, _ := httputil.DumpRequest(r, true)
		fmt.Println(string(req1))

		req, err := parser.QueryInfoParser(r)
		if err != nil {
			fmt.Printf("Parse Query Info erro : %s\n", err.Error())
			httpx.Error(w, err)
			return
		}
		l := logic.NewQueryLogic(r.Context(), svcCtx)
		resp, err := l.Query(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
