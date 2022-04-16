package logic

import (
	"Hermes/rpc/transform/model"
	"context"

	"Hermes/rpc/transform/inter/svc"
	"Hermes/rpc/transform/transform"

	"github.com/zeromicro/go-zero/core/logx"
)

type HermesenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHermesenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HermesenLogic {
	return &HermesenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HermesenLogic) Hermesen(in *transform.HermesenReq) (*transform.HermesResp, error) {
	_, err := l.svcCtx.Model.Insert(l.ctx, &model.Hermesd{
		Aletname:        in.AlertName,
		Receiveraddress: in.ReceiverAddress,
		Aggeraterules:   in.AggeratuRule,
		Returnvalueflag: in.ReturnValueFlag,
	})
	if err != nil {
		return nil, err
	}

	return &transform.HermesResp{
		AggeratuRule:    in.AggeratuRule,
		ReturnValueFlag: in.ReturnValueFlag,
		AlertName:       in.AlertName,
		ReceiverAddress: in.ReceiverAddress,
	}, nil
}
