package handler

import (
	"Hermes/api/inter/logic"
	"Hermes/api/inter/svc"
	"Hermes/pkg/parser"
	"encoding/json"
	"fmt"
	alertdata "github.com/prometheus/alertmanager/template"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func HermesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Printf("**************\n")

		data := alertdata.Data{}

		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			fmt.Printf("%s\n", err.Error())
		}
		alertNumber := len(data.Alerts)

		alertInfo, err := parser.AlertInfoParser(data, alertNumber)
		if err != nil {
			fmt.Printf("alert parse fail.\n")
			httpx.Error(w, err)
		}

		//res, _ := json.Marshal(data)
		//
		//var out bytes.Buffer
		//err := json.Indent(&out, res, "", "\t")
		//
		//out.WriteTo(os.Stdout)
		//println()

		fmt.Printf("**************\n")

		//var req []types.HermesReq

		//if err := httpx.Parse(r, &req); err != nil {
		//	httpx.Error(w, err)
		//	return
		//}

		l := logic.NewHermesLogic(r.Context(), svcCtx)
		resp, err := l.Hermes(alertInfo)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
