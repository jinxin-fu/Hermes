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
	client, initErr = api.NewClient(api.Config{
		Address: "http://192.168.1.51:31445",
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

	v1api1 := v1.NewAPI(client)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
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
	wg.Wait()
	return qResps

}

func GetQuery(api v1.API, ctx context.Context, resp types.HermesResp, limiter chan bool, responseCh chan types.QueryResp, wg *sync.WaitGroup) {
	defer wg.Done()

	result, _, err := api.Query(ctx, resp.AggerateRules, time.Now())
	if err != nil {
		fmt.Printf("Error querying Prometheus: %v\n", err)
		os.Exit(1)
	}
	//obj := &unstructured.Unstructured{}
	//str := result.String()
	println(len(result.(model.Vector)))
	v := result.(model.Vector)[0]
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
	fmt.Printf("Result: \n%v\n", result.String())
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
	}
	<-limiter
}
