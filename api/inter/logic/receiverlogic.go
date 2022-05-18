package logic

import (
	"Hermes/api/inter/svc"
	"Hermes/api/inter/types"
	realtimemprocess "Hermes/pkg/realtimeprocess"
	"context"
	"github.com/prometheus/prometheus/prompb"
	"github.com/zeromicro/go-zero/core/logx"
)

type ReceiverLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

const METRICNAMELABELKEY = "__name__"

func NewReceiverLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReceiverLogic {
	return &ReceiverLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReceiverLogic) Receiver(req prompb.WriteRequest) (resp *types.ReveicerResp, err error) {
	for _, ts := range req.Timeseries {
		metricName := ""
		for _, v := range ts.Labels {
			if v.Name == METRICNAMELABELKEY {
				metricName = v.Value
				break
			}
			l.Logger.Error("Get metric name label fail")
			return
		}
		ds, is := realtimemprocess.FindDistributeTarget(metricName)
		if is {
			(*ds).ReceiverChan <- ts
		} else {
			l.Logger.Error("")
		}
	}

	return
}
