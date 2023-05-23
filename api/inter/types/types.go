package types

import "github.com/prometheus/common/model"

type ExpandReq struct {
	AlertName string `form:"alertName"`
}

type ExpandResp struct {
	AlertName       string `json:"alertName"`
	AggerateRules   string `json:"aggerateRules"`
	ReceiverAddress string `json:"receiveraddress"`
	ReturnValueFlag string `json:"returnValueFlag"`
}

type AlertsFromAlertmanage struct {
	Alerts        HermesReqs
	Receiver      string
	Status        string
	MacthedAlerts int
}

type HermesReqs []HermesReq

type HermesReq struct {
	AlertName       string `from:"alertName"`
	AggerateRules   string `from:"aggerateRules"`
	ReceiverAddress string `from:"receiveraddress"`
	ReturnValueFlag string `from:"returnValueFlag"`
}

type HermesResp struct {
	AlertName       string `json:"alertName"`
	AggerateRules   string `json:"aggerateRules"`
	ReceiverAddress string `json:"receiveraddress"`
	ReturnValueFlag string `json:"returnValueFlag"`
}

type AlertmanagerResp struct {
	InProcessNumber int    `json:"inProcessNumber"`
	Receiver        string `json:"receiver"`
	MatchedAlerts   int    `json:"matchedAlerts"`
}

type QueryResp struct {
	Name        string
	Expression  string
	Flag        bool
	Destination string
	QValue      model.Vector
	Err         error
}

type DistributeResult struct {
	Receiver   string
	Status     string
	StatusCode int
	Err        error
}

type ReceiverReq struct {
}

type ReveicerResp struct {
}

type QueryReq struct {
	MethodType string `from:"methodType"`
	QuerySql   string `from:"querySlq"`
	StartTime  string `from:"startTime"`
	EndTime    string `from:"endTime"`
}

type QueryPromResp struct {
}
