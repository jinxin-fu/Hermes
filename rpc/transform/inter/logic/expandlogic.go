package logic

import (
	"context"

	"github.com/Hermes/rpc/transform/inter/svc"
	"github.com/Hermes/rpc/transform/transform"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExpandLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewExpandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExpandLogic {
	return &ExpandLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ExpandLogic) Expand(in *transform.ExpandReq) (*transform.ExpandResp, error) {

	res, err := l.svcCtx.Model.FindOne(l.ctx, in.AlertName)
	if err != nil {
		return nil, err
	}

	return &transform.ExpandResp{
		AlertName:       res.Aletname,
		AggeratuRule:    res.Aggeraterules,
		ReturnValueFlag: res.Returnvalueflag,
		ReceiverAddress: res.Receiveraddress,
	}, nil
}
