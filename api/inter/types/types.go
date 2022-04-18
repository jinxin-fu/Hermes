package types

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
