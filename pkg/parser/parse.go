/**
 * Created with IntelliJ goland.
 * @Auther: jinxin
 * @Date: 2022/04/16/10:27
 * @Description:
 */
package parser

import (
	"Hermes/api/inter/types"
	"fmt"
	"github.com/pkg/errors"
	"github.com/prometheus/alertmanager/template"
	"github.com/prometheus/common/model"
	"math"
	"net/http"
	"strconv"
	"time"
)

const (
	ALERTNAME       = "alertname"
	AGGERATRRULES   = "aggerateRules"
	RECEIVERADDRESS = "receiverAddress"
	RETURNVALUEFLAG = "returnValueFlag"

	QUERY      = "query"
	QUERYRANGE = "query_range"
)

var (
	minTime = time.Unix(math.MinInt64/1000+62135596801, 0).UTC()
	maxTime = time.Unix(math.MaxInt64/1000-62135596801, 999999999).UTC()

	minTimeFormatted = minTime.Format(time.RFC3339Nano)
	maxTimeFormatted = maxTime.Format(time.RFC3339Nano)
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

func QueryInfoParser(r *http.Request) (types.QueryReq, error) {
	var queryReq types.QueryReq
	ts, err := parseTimeParam(r, "time", time.Now())
	if err != nil {
		return types.QueryReq{}, err
	}

	methodQuery := r.FormValue("query")
	methodQueryRange := r.FormValue("query_range")
	if methodQueryRange == "" && methodQuery == "" {
		return types.QueryReq{}, errors.Errorf("query method emprty.")
	}
	if methodQuery != "" && methodQueryRange != "" {
		return types.QueryReq{}, errors.Errorf("query method reduplicative, only one method allowed in a single request.")
	}

	if methodQuery == "" {
		queryReq.QuerySql = methodQueryRange
		queryReq.MethodType = QUERYRANGE
		queryReq.StartTime, _ = parseTime(r.FormValue("start"))
		queryReq.EndTime, _ = parseTime(r.FormValue("end"))
		step := r.FormValue("step")
		if step == "" {
			queryReq.Step = time.Minute //temporary query step value
		} else {
			stepVal, err := parseDuration(r.FormValue("step"))
			if err != nil {
				fmt.Printf("Parse duration step error, use default value 1 minute")
				queryReq.Step = time.Minute
			} else {
				queryReq.Step = stepVal
			}

		}

	} else {
		queryReq.QuerySql = methodQuery
		queryReq.MethodType = QUERY
		queryReq.Time = ts
	}
	return queryReq, nil
}

func parseTimeParam(r *http.Request, paramName string, defaultValue time.Time) (time.Time, error) {
	val := r.FormValue(paramName)
	if val == "" {
		return defaultValue, nil
	}
	result, err := parseTime(val)
	if err != nil {
		return time.Time{}, errors.Wrapf(err, "Invalid time value for '%s'", paramName)
	}
	return result, nil
}

func parseTime(s string) (time.Time, error) {
	if t, err := strconv.ParseFloat(s, 64); err == nil {
		s, ns := math.Modf(t)
		ns = math.Round(ns*1000) / 1000
		return time.Unix(int64(s), int64(ns*float64(time.Second))).UTC(), nil
	}
	if t, err := time.Parse(time.RFC3339Nano, s); err == nil {
		return t, nil
	}

	switch s {
	case minTimeFormatted:
		return minTime, nil
	case maxTimeFormatted:
		return maxTime, nil
	}
	return time.Time{}, errors.Errorf("cannot parse %q to a valid timestamp", s)
}

func parseDuration(s string) (time.Duration, error) {
	if d, err := strconv.ParseFloat(s, 64); err == nil {
		ts := d * float64(time.Second)
		if ts > float64(math.MaxInt64) || ts < float64(math.MinInt64) {
			return 0, errors.Errorf("cannot parse %q to a valid duration. It overflows int64", s)
		}
		return time.Duration(ts), nil
	}
	if d, err := model.ParseDuration(s); err == nil {
		return time.Duration(d), nil
	}
	return 0, errors.Errorf("cannot parse %q to a valid duration", s)
}
