package logic

import (
	"Hermes/api/inter/svc"
	"Hermes/api/inter/types"
	"context"
	"fmt"
	"github.com/prometheus/common/model"
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
		m := make(model.Metric, len(ts.Labels))
		for _, l := range ts.Labels {
			m[model.LabelName(l.Name)] = model.LabelValue(l.Value)
		}
		fmt.Println(m)

		for _, s := range ts.Samples {
			s = s
			//fmt.Printf("  %f %d %s %s\n", s.Value, s.Timestamp, time.UnixMilli(s.Timestamp), time.Now())
		}
	}

	return
}
