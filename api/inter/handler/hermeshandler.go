package handler

import (
	"encoding/json"
	"fmt"
	"github.com/jinxin-fu/hermes/api/inter/logic"
	"github.com/jinxin-fu/hermes/api/inter/svc"
	"github.com/jinxin-fu/hermes/pkg/parser"
	alertdata "github.com/prometheus/alertmanager/template"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func HermesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

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

		fmt.Printf("Trigger alerts number: %d\n", len(data.Alerts))
		fmt.Printf("Valid hyperos alert number: %d\n", len(alertInfo.Alerts))
		//res, _ := json.Marshal(alertInfo.Alerts)
		//
		//var out bytes.Buffer
		//if err := json.Indent(&out, res, "", "\t"); err != nil {
		//	panic(err)
		//}
		//
		//out.WriteTo(os.Stdout)
		//println()
		//var req []types.HermesReq

		//if err := httpx.Parse(r, &req); err != nil {
		//	httpx.Error(w, err)
		//	return
		//}
		if alertInfo.MacthedAlerts == 0 {
			httpx.Error(w, fmt.Errorf("no matched hyperos-rule trigger in hermes."))
		}

		l := logic.NewHermesLogic(r.Context(), svcCtx)
		resp, err := l.Hermes(alertInfo)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
