package logic

import (
	"context"

	"github.com/Hermes/api/inter/svc"
	"github.com/Hermes/api/inter/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExpandLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExpandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExpandLogic {
	return &ExpandLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExpandLogic) Expand(req types.ExpandReq) (types.ExpandResp, error) {
	// add manually
	//resp, err := l.svcCtx.Transformer.Expand(l.ctx, &transformer.ExpandReq{
	//	AlertName: req.AlertName,
	//})
	//if err != nil {
	//	return types.ExpandResp{}, err
	//}
	//return types.ExpandResp{
	//	AlertName:       resp.AlertName,
	//	ReturnValueFlag: resp.ReturnValueFlag,
	//	ReceiverAddress: resp.ReceiverAddress,
	//	AggerateRules:   resp.AggeratuRule,
	//}, nil
	return types.ExpandResp{}, nil
}
