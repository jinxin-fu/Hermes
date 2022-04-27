/**
 * Created with IntelliJ goland.
 * @Auther: jinxin
 * @Date: 2022/04/12/18:56
 * @Description:
 */
package querier

import (
	"Hermes/api/inter/types"
	"context"
	"fmt"
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	"log"
	"os"
	"sync"
	"time"
)

//var PromRule promonitv1.PrometheusRule
//var PromRuleSpec promonitv1.PrometheusRuleSpec

var client api.Client
var initErr error

func init() {
	add := os.Getenv("PROMETHEUSADDRESS")
	if add == "" {
		add = "http://192.168.1.51:31445"
	}
	client, initErr = api.NewClient(api.Config{
		Address: add,
	})
	if initErr != nil {
		fmt.Printf("Error creating client: %v\n", initErr)
		log.Fatal(initErr)
	}

}

func PrometheusQuery(resp []types.HermesResp) []types.QueryResp {
	var qResps []types.QueryResp
	wg := &sync.WaitGroup{}
	limiter := make(chan bool, 20)
	defer close(limiter)
	v1api1 := v1.NewAPI(client)
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()

	responseCh := make(chan types.QueryResp)
	wgResponse := &sync.WaitGroup{}
	go func() {
		wgResponse.Add(1)
		for reponse := range responseCh {
			qResps = append(qResps, reponse)
		}
		wgResponse.Done()
	}()

	for _, v := range resp {
		wg.Add(1)
		limiter <- true
		go GetQuery(v1api1, ctx, v, limiter, responseCh, wg)

	}

	wg.Wait()
	fmt.Println("Prometheus query process finished")
	close(responseCh)
	wgResponse.Wait()
	return qResps

}

func GetQuery(api v1.API, ctx context.Context, resp types.HermesResp, limiter chan bool, responseCh chan types.QueryResp, wg *sync.WaitGroup) {
	defer wg.Done()

	result, _, err := api.Query(ctx, resp.AggerateRules, time.Now())
	if err != nil {
		fmt.Printf("Error querying Prometheus, alertName:%s error:%v\n", resp.AlertName, err)

	}

	//obj := &unstructured.Unstructured{}
	//str := result.String()
	//println(len(result.(model.Vector)))
	if result.Type() != model.ValVector {
		responseCh <- types.QueryResp{
			Err: fmt.Errorf("query result is not type vector"),
		}
		return
	}
	v := result.(model.Vector)[0] // TODO默认去第一个metric值，具体逻辑等上真是环境上调试
	value := v.Value
	//for _, v := range result.(model.Vector) {
	//	fmt.Printf("value: %v\n", v.Value)
	//	for k, i := range v.Metric {
	//		println("k: ", k)
	//		println("v: ", i)
	//
	//	}
	//}
	//for k,v := range result.(model.Vector)
	//str := result.String()
	var flag bool
	//fmt.Printf("Result: \n%v\n", result.String())
	if resp.ReturnValueFlag == "false" {
		flag = false
	} else {
		flag = true
	}
	responseCh <- types.QueryResp{
		Name:        resp.AlertName,
		Destination: resp.ReceiverAddress,
		Expression:  resp.AggerateRules,
		Flag:        flag,
		Value:       float64(value),
		Err:         nil,
	}
	<-limiter
}
