/**
 * Created with IntelliJ goland.
 * @Auther: jinxin
 * @Date: 2022/05/10/10:55
 * @Description:
 */
package realtimemprocess

import (
	"github.com/zeromicro/go-zero/core/stringx"
	"go.uber.org/zap"
	"time"
)

func init() {
	go Srs2RunningMap()
}

func equalSubscriber(src, dst *SubscribeRule) bool {
	if src.Callback != dst.Callback || len(src.Metrics) != len(dst.Metrics) {
		return false
	}
	for _, v := range src.Metrics {
		if !stringx.Contains(dst.Metrics, v) {
			return false
		}
	}
	return true
}

func Srs2RunningMap() {
	ticker := time.NewTicker(time.Second)
	for {
		<-ticker.C
		for k, v := range GlobalSrs {
			if ds, ok := RunningDsMap[k]; !ok {
				RunningDsMap[k] = newDistributor(&v)
				go RunningDsMap[k].Distribute()
			} else {
				if !equalSubscriber(ds.SubscribeRule, &v) {
					ds.DestroyDistributor()
					RunningDsMap[k] = newDistributor(&v)
					go RunningDsMap[k].Distribute()
				}
			}
		}
		for k, v := range RunningDsMap {
			if _, ok := GlobalSrs[k]; !ok {
				v.DestroyDistributor()
				Rlogger.Info("close distributor", zap.String("index", k))
				delete(RunningDsMap, k)
			}
		}
	}
}
