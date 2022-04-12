package logic

import (
	"context"

	"Hermes/rpc/transform/inter/svc"
	"Hermes/rpc/transform/pb/transform"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShortenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewShortenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShortenLogic {
	return &ShortenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ShortenLogic) Shorten(in *transform.HermesenReq) (*transform.HermesResp, error) {
	// todo: add your logic here and delete this line

	return &transform.HermesResp{}, nil
}
