package logic

import (
	"Hermes/rpc/transform/transformer"
	"context"

	"Hermes/api/inter/svc"
	"Hermes/api/inter/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HermesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHermesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HermesLogic {
	return &HermesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HermesLogic) Hermes(req types.AlertsFromAlertmanage) (types.AlertmanagerResp, error) {
	// add manually
	inProcessReuest := 0
	for i := 0; i < req.MacthedAlerts; i++ {
		_, err := SendToRpc(l, req.Alerts[i])
		if err != nil {
			return types.AlertmanagerResp{}, err
		}
		inProcessReuest++
	}

	return types.AlertmanagerResp{
		Receiver:        req.Receiver,
		MatchedAlerts:   req.MacthedAlerts,
		InProcessNumber: inProcessReuest,
	}, nil

}

func SendToRpc(l *HermesLogic, req types.HermesReq) (types.HermesResp, error) {
	resp, err := l.svcCtx.Transformer.Hermesen(l.ctx, &transformer.HermesenReq{
		AlertName:       req.AlertName,
		ReceiverAddress: req.ReceiverAddress,
		ReturnValueFlag: req.ReturnValueFlag,
		AggeratuRule:    req.AggerateRules,
	})
	if err != nil {
		return types.HermesResp{}, err
	}
	return types.HermesResp{
		AggerateRules:   resp.AggeratuRule,
		AlertName:       resp.AlertName,
		ReturnValueFlag: resp.ReturnValueFlag,
		ReceiverAddress: resp.ReceiverAddress,
	}, nil

}
