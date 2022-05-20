/**
 * Created with IntelliJ goland.
 * @Auther: jinxin
 * @Date: 2022/04/16/10:27
 * @Description:
 */
package parser

import (
	"fmt"
	"github.com/Hermes/api/inter/types"
	"github.com/prometheus/alertmanager/template"
)

const (
	ALERTNAME       = "alertname"
	AGGERATRRULES   = "aggerateRules"
	RECEIVERADDRESS = "receiverAddress"
	RETURNVALUEFLAG = "returnValueFlag"
)

func AlertInfoParser(data template.Data, alertNumber int) (types.AlertsFromAlertmanage, error) {
	var alertsInfo types.AlertsFromAlertmanage
	count := 0
	if alertNumber == 0 || alertNumber != len(data.Alerts) {
		return types.AlertsFromAlertmanage{}, fmt.Errorf("Alert number error.\n")
	}

	alertsInfo.Receiver = data.Receiver
	alertsInfo.Status = data.Status
	for _, v := range data.Alerts {
		if _, ok := v.Annotations[AGGERATRRULES]; !ok {
			continue
		}
		if _, ok := v.Labels[ALERTNAME]; !ok {
			continue
		}
		alertsInfo.Alerts = append(alertsInfo.Alerts, types.HermesReq{
			AlertName:       v.Labels[ALERTNAME],
			AggerateRules:   v.Annotations[AGGERATRRULES],
			ReceiverAddress: v.Annotations[RECEIVERADDRESS],
			ReturnValueFlag: v.Annotations[RETURNVALUEFLAG],
		})
		count++
	}
	alertsInfo.MacthedAlerts = count

	return alertsInfo, nil
}
