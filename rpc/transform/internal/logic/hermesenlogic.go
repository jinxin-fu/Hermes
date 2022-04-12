package logic

import (
	"Hermes/rpc/transform/model"
	"context"
	"github.com/zeromicro/go-zero/core/hash"

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
	key := hash.Md5Hex([]byte(in.Url))[:6]
	_, err := l.svcCtx.Model.Insert(l.ctx, &model.Hermesd{
		Hermes: key,
		Url:    in.Url,
	})
	if err != nil {
		return nil, err
	}

	return &transform.HermesResp{Hermesen: key}, nil
}
