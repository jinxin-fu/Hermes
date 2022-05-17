/**
 * Created with IntelliJ goland.
 * @Auther: jinxin
 * @Date: 2022/05/16/18:17
 * @Description:
 */
package realtimemprocess

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stringx"
	"time"
)

type SubscribeRule struct {
	Metrics  []string
	Callback string
}

type Distributor struct {
	ReceiverChan       chan string
	SubscribeRule      *SubscribeRule
	LifeControllerChan chan int
}

func newDistributor(sr *SubscribeRule) *Distributor {
	return &Distributor{
		ReceiverChan:       make(chan string),
		SubscribeRule:      sr,
		LifeControllerChan: make(chan int),
	}
}

func (d *Distributor) Distribute() {
	fmt.Printf("ds start: %s, callback:%s\n", d.SubscribeRule.Metrics, d.SubscribeRule.Callback)
	for {
		select {
		case req := <-d.ReceiverChan:
			// do request to callback
			fmt.Printf("%s -- ds:%s --callback:%s\n", time.Now(), req, d.SubscribeRule.Callback)
			time.Sleep(10 * time.Second)
		case <-d.LifeControllerChan:
			close(d.ReceiverChan)
			close(d.LifeControllerChan)
			fmt.Printf("Distributer close, consumer address:%s", d.SubscribeRule.Callback)
			return
		default:
		}
	}
}

func (d *Distributor) IsDistributeTarget(m string) bool {
	return stringx.Contains(d.SubscribeRule.Metrics, m)
}

func (d *Distributor) DestroyDistributor() {
	d.LifeControllerChan <- 1

}
