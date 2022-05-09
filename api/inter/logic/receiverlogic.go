package logic

import (
	"context"

	"Hermes/api/inter/svc"
	"Hermes/api/inter/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReceiverLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReceiverLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReceiverLogic {
	return &ReceiverLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReceiverLogic) Receiver(req *types.ReceiverReq) (resp *types.ReveicerResp, err error) {
	// todo: add your logic here and delete this line

	return
}
