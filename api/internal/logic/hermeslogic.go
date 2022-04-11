package logic

import (
	"context"

	"Hermes/api/internal/svc"
	"Hermes/api/internal/types"

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

func (l *HermesLogic) Hermes(req *types.HermesReq) (resp *types.HermesResp, err error) {
	// todo: add your logic here and delete this line

	return
}
