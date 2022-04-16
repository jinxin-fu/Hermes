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

func (l *HermesLogic) Hermes(req types.AlertsFromAlertmanage) (types.HermesResp, error) {
	// add manually

	resp, err := l.svcCtx.Transformer.Hermesen(l.ctx, &transformer.HermesenReq{
		AlertName:       req.Alerts[1].AlertName,
		ReceiverAddress: req.Alerts[1].ReceiverAddress,
		ReturnValueFlag: req.Alerts[1].ReturnValueFlag,
		AggeratuRule:    req.Alerts[1].AggerateRules,
	})
	if err != nil {
		return types.HermesResp{}, err
	}

	return types.HermesResp{
		AlertName:       resp.AlertName,
		AggerateRules:   resp.AggeratuRule,
		ReceiverAddress: resp.ReceiverAddress,
		ReturnValueFlag: resp.ReturnValueFlag,
	}, nil

}
