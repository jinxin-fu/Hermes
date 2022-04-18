package logic

import (
	"Hermes/rpc/transform/transformer"
	"context"
	"fmt"
	"sync"

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

func processAlerts(l *HermesLogic, req types.AlertsFromAlertmanage) []types.HermesResp {
	var result []types.HermesResp
	wg := &sync.WaitGroup{}
	limiter := make(chan bool, 20) //限制最大并发
	defer close(limiter)

	responseCh := make(chan types.HermesResp, req.MacthedAlerts)
	wgResponse := &sync.WaitGroup{}
	go func() {
		wgResponse.Add(1)
		for response := range responseCh {
			result = append(result, response)
		}
		wgResponse.Done()
	}()

	for _, v := range req.Alerts {
		wg.Add(1)
		limiter <- true
		go sendToRpc(l, v, limiter, responseCh, wg)
	}

	wg.Wait()
	fmt.Println("Alerts process finished.")
	close(responseCh)

	wgResponse.Wait()
	return result

}

func sendToRpc(l *HermesLogic, req types.HermesReq, limiter chan bool, responseCh chan types.HermesResp, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := l.svcCtx.Transformer.Hermesen(l.ctx, &transformer.HermesenReq{
		AlertName:       req.AlertName,
		ReceiverAddress: req.ReceiverAddress,
		ReturnValueFlag: req.ReturnValueFlag,
		AggeratuRule:    req.AggerateRules,
	})
	if err != nil {
		fmt.Printf("process fail %s", err.Error())
		responseCh <- types.HermesResp{}
	}
	responseCh <- types.HermesResp{
		AlertName:       resp.AlertName,
		AggerateRules:   resp.AggeratuRule,
		ReceiverAddress: resp.ReceiverAddress,
		ReturnValueFlag: resp.ReturnValueFlag,
	}
	<-limiter

}
func (l *HermesLogic) Hermes(req types.AlertsFromAlertmanage) (types.AlertmanagerResp, error) {
	// add manually

	res := processAlerts(l, req)
	return types.AlertmanagerResp{
		Receiver:        req.Receiver,
		MatchedAlerts:   req.MacthedAlerts,
		InProcessNumber: len(res),
	}, nil

}
