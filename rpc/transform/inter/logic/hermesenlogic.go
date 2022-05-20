package logic

import (
	"context"
	"fmt"
	"github.com/Hermes/rpc/transform/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/Hermes/rpc/transform/inter/svc"
	"github.com/Hermes/rpc/transform/transform"

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

	_, err := l.svcCtx.Model.FindOne(l.ctx, in.AlertName)
	if err == nil {
		fmt.Printf("object exist , update data.\n")
		err := l.svcCtx.Model.Update(l.ctx, &model.Hermesd{
			Aletname:        in.AlertName,
			Aggeraterules:   in.AggeratuRule,
			Returnvalueflag: in.ReturnValueFlag,
			Receiveraddress: in.ReceiverAddress,
		})
		if err != nil {
			return &transform.HermesResp{}, fmt.Errorf("Update data faild error  %s", err.Error())
		}
		return &transform.HermesResp{
			AggeratuRule:    in.AggeratuRule,
			ReturnValueFlag: in.ReturnValueFlag,
			AlertName:       in.AlertName,
			ReceiverAddress: in.ReceiverAddress,
		}, nil

	} else if err == sqlx.ErrNotFound {
		_, err = l.svcCtx.Model.Insert(l.ctx, &model.Hermesd{
			Aletname:        in.AlertName,
			Receiveraddress: in.ReceiverAddress,
			Aggeraterules:   in.AggeratuRule,
			Returnvalueflag: in.ReturnValueFlag,
		})
		if err != nil {
			return nil, err
		}
	} else {
		return &transform.HermesResp{}, nil
	}
	return &transform.HermesResp{}, nil
}
