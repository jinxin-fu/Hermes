package logic

import (
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

func (l *HermesLogic) Hermes(req types.HermesReq) (types.HermesResp, error) {
	// add manually

	//resp, err := l.svcCtx.Transformer.Hermesen(l.ctx, &transformer.HermesenReq{
	//	Url: req.Url,
	//})
	//if err != nil {
	//	return types.HermesResp{}, err
	//}
	//
	//return types.HermesResp{
	//	Hermesen: resp.Hermesen,
	//}, nil

	return types.HermesResp{}, nil
}
