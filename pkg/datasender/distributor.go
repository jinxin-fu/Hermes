/**
 * Created with IntelliJ goland.
 * @Auther: jinxin
 * @Date: 2022/04/18/11:30
 * @Description:
 */
package datasender

import (
	"Hermes/api/inter/types"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type RequestError struct {
	Name string
	Err  error
}

func Distributor(qResps []types.QueryResp) ([]types.DistributeResult, error) {
	var dRes []types.DistributeResult
	wg := &sync.WaitGroup{}
	limiter := make(chan bool, 10)
	defer close(limiter)
	responseCh := make(chan types.DistributeResult)
	wgResponse := &sync.WaitGroup{}
	go func() {
		wgResponse.Add(1)
		for response := range responseCh {
			dRes = append(dRes, response)
		}
		wgResponse.Done()
	}()

	for _, v := range qResps {
		wg.Add(1)
		limiter <- true
		go DoRequest(v, responseCh, limiter, wg)
	}
	wg.Wait()
	fmt.Println("Distributor process finished.")
	close(responseCh)
	wgResponse.Wait()
	return dRes, nil
}

func DoRequest(qReq types.QueryResp, resCh chan types.DistributeResult, limiter chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	address := "http://127.0.0.1:5000/parsePrometheusAlert"
	//ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	req, err := http.NewRequest(http.MethodPost, address, nil)
	if err != nil {
		fmt.Printf("New request error: %s\n", err.Error())
	}
	//req.WithContext(ctx)
	//var t int64 = 5
	//ctx, _ := context.WithTimeout(context.Background(), time.Duration(t)*time.Second)
	params := req.URL.Query()
	params.Add("alertName", qReq.Name)
	params.Add("witchValue", strconv.FormatBool(qReq.Flag))
	if qReq.Flag == true {
		params.Add("value", strconv.FormatFloat(qReq.Value, 'E', -1, 64))
	}
	cli := http.Client{Timeout: 3 * time.Second}
	response, err := cli.Do(req)
	//response, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Distribute to consumer err :%s", err.Error())
	}
	defer response.Body.Close()
	resCh <- types.DistributeResult{
		Receiver:   qReq.Destination,
		Status:     response.Status,
		StatusCode: response.StatusCode,
		Err:        err,
	}

	<-limiter
}
