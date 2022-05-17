/**
 * Created with IntelliJ goland.
 * @Auther: jinxin
 * @Date: 2022/05/16/18:17
 * @Description:
 */
package realtimemprocess

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/prometheus/prometheus/prompb"
	"github.com/zeromicro/go-zero/core/stringx"
	"net/http"
	"time"
)

var GlobalSrs map[string]SubscribeRule
var RunningDsMap map[string]*Distributor

func init() {
	GlobalSrs = make(map[string]SubscribeRule, 100)
	RunningDsMap = make(map[string]*Distributor, 100)
}

func AddGlobalSrs(name, callback string, metrics []string) {
	GlobalSrs[name] = SubscribeRule{
		Metrics:  metrics,
		Callback: callback,
	}
}

func DeleteGlobalSrs(name string) {
	delete(GlobalSrs, name)
}

func IsInGlobalSrs(name string) bool {
	if _, ok := GlobalSrs[name]; ok {
		return true
	}
	return false
}

type SubscribeRule struct {
	Metrics  []string
	Callback string
}

type Distributor struct {
	ReceiverChan       chan prompb.TimeSeries
	SubscribeRule      *SubscribeRule
	LifeControllerChan chan int
}

func newDistributor(sr *SubscribeRule) *Distributor {
	return &Distributor{
		ReceiverChan:       make(chan prompb.TimeSeries),
		SubscribeRule:      sr,
		LifeControllerChan: make(chan int),
	}
}

func FindDistributeTarget(metricName string) (*Distributor, bool) {
	for _, v := range RunningDsMap {
		if stringx.Contains(v.SubscribeRule.Metrics, metricName) {
			return v, true
		}
	}
	return nil, false
}

func (d *Distributor) Distribute() {
	fmt.Printf("ds start: %s, callback:%s\n", d.SubscribeRule.Metrics, d.SubscribeRule.Callback)
	for {
		select {
		case req := <-d.ReceiverChan:
			go d.doDistribute(req)
			fmt.Printf("%s -- ds:%s --callback:%s\n", time.Now(), req, d.SubscribeRule.Callback)
		case <-d.LifeControllerChan:
			close(d.ReceiverChan)
			close(d.LifeControllerChan)
			fmt.Printf("Distributer close, consumer address:%s", d.SubscribeRule.Callback)
			return
		default:
		}
	}
}

func (d *Distributor) doDistribute(body prompb.TimeSeries) {

	sendBody, err := json.Marshal(body)
	if err != nil {
		return
	}
	reader := bytes.NewReader(sendBody)
	req, err := http.NewRequest(http.MethodPost, d.SubscribeRule.Callback, reader)
	//req, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:5001/realtimereceiver", reader) //for test
	if err != nil {
		fmt.Printf("New request error: %s\n", err.Error())
	}
	req.Header.Set("subscribeType", "SubsRealTime")
	//params := req.URL.Query()
	//params.Set("subscribeType", "SubsRealTime")

	cli := http.Client{Timeout: 10 * time.Second}
	_, err = cli.Do(req)
	if err != nil {
		fmt.Printf("Distribute to consumer err :%s", err.Error())
	}
	//if response.Body != nil {
	//	response.Body.Close()
	//}

}

func (d *Distributor) IsDistributeTarget(m string) bool {
	return stringx.Contains(d.SubscribeRule.Metrics, m)
}

func (d *Distributor) DestroyDistributor() {
	d.LifeControllerChan <- 1
}
