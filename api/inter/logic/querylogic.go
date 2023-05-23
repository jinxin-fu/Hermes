/**
 * Created with IntelliJ goland.
 * @Auther: jinxin
 * @Date: 2023/05/23/14:14
 * @Description:
 */
package logic

import (
	"Hermes/api/inter/svc"
	"Hermes/api/inter/types"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type QueryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryLogic {
	return &QueryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryLogic) Query(req types.QueryReq) (types.QueryPromResp, error) {
	return types.QueryPromResp{}, nil
}
