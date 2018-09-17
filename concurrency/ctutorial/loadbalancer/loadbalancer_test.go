package loadbalancer

import (
	"encoding/base64"
	"fmt"
	"sync"
	"testing"
)

func TestLoadBalancer(t *testing.T) {
	worker := func(req Request) interface{} {
		//time.Sleep(200 * time.Millisecond)
		return base64.StdEncoding.EncodeToString([]byte(req.data.(string)))
	}
	b := NewBalancer(worker)
	//b.SetDebug(true)
	go b.Balance()

	var wg sync.WaitGroup
	for i := 0; i < 22; i++ {
		wg.Add(1)
		go func() {
			respChan := b.Add("1234567")
			resp := <-respChan
			if resp.(string) != "MTIzNDU2Nw==" {
				t.Errorf("Resp should be 'MTIzNDU2Nw==', instead of %s", resp.(string))
			}
			wg.Done()
		}()
		fmt.Println(b.NumOfWorkers())
	}
	wg.Wait()

}
