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
			if v.Name == "__name__" {
				metricName = v.Value
				break
			}
			return
		}
		ds, is := realtimemprocess.FindDistributeTarget(metricName)
		if is {
			(*ds).ReceiverChan <- ts
		}
	}

	return
}
