package logic

import (
	"Hermes/rpc/transform/transformer"
	"context"

	"Hermes/api/internal/svc"
	"Hermes/api/internal/types"

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
	resp, err := l.svcCtx.Transformer.Expand(l.ctx, &transformer.ExpandReq{
		Hermesen: req.Hermesen,
	})
	if err != nil {
		return types.ExpandResp{}, err
	}
	return types.ExpandResp{
		Url: resp.Url,
	}, nil
}
