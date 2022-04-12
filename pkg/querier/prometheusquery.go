/**
 * Created with IntelliJ goland.
 * @Auther: jinxin
 * @Date: 2022/04/12/18:56
 * @Description:
 */
package querier

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	"os"
	"time"
)

//var PromRule promonitv1.PrometheusRule
//var PromRuleSpec promonitv1.PrometheusRuleSpec

func main() {
	client, err := api.NewClient(api.Config{
		Address: "http://192.168.1.51:32699",
	})
	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
		os.Exit(1)
	}

	v1api1 := v1.NewAPI(client)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//result, _, err := v1api1.Query(ctx, "container_cpu_load_average_10s{container=\"activator\"}", time.Now())
	result, _, err := v1api1.Query(ctx, "rate(prometheus_tsdb_head_samples_appended_total[5m])", time.Now())
	if err != nil {
		fmt.Printf("Error querying Prometheus: %v\n", err)
		os.Exit(1)
	}
	//obj := &unstructured.Unstructured{}
	//str := result.String()
	println(len(result.(model.Vector)))
	for _, v := range result.(model.Vector) {
		fmt.Printf("value: %v\n", v.Value)
		for k, i := range v.Metric {
			println("k: ", k)
			println("v: ", i)

		}
	}
	//for k,v := range result.(model.Vector)
	//str := result.String()

	fmt.Printf("Result: \n%v\n", result.String())
}
