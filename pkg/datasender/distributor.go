/**
 * Created with IntelliJ goland.
 * @Auther: jinxin
 * @Date: 2022/04/18/11:30
 * @Description:
 */
package datasender

import (
	"Hermes/api/inter/types"
	"context"
	"log"
	"net/http"
	"time"
)

type RequestError struct {
	Name string
	Err  error
}

func Distributor(qResps []types.QueryResp) error {

	return nil
}

func DoRequest(address string, errCh chan RequestError) {
	req, err := http.NewRequest(http.MethodPost, address, nil)
	if err != nil {
		log.Fatal(err)
	}
	var t int64 = 5
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(t)*time.Second)
	
}
