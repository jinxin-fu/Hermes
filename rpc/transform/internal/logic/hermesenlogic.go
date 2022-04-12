package logic

import (
	"context"

	"Hermes/rpc/transform/internal/svc"
	"Hermes/rpc/transform/pb/transform"

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
	// todo: add your logic here and delete this line

	return &transform.HermesResp{}, nil
}
