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
	"Hermes/pkg/querier"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/prometheus/common/model"
	"github.com/zeromicro/go-zero/core/logx"
)

const (
	QUERY      = "query"
	QUERYRANGE = "query_range"
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
	fmt.Printf("%+v\n", req)
	var val model.Vector
	var err error
	switch req.MethodType {
	case QUERYRANGE:
		val, err = querier.QueryRange(req.QuerySql, req.StartTime, req.EndTime, req.Step)
	case QUERY:
		val, err = querier.Query(req.QuerySql, req.Time)
	default:
		l.Logger.Error("Query method error.")
		return types.QueryPromResp{}, errors.Errorf("Query method error. type: %s", req.MethodType)
	}
	if err != nil {
		return types.QueryPromResp{}, err
	}
	fmt.Printf("%+v\n", val)
	return types.QueryPromResp{}, nil
}
