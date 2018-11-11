package loadbalancer

import (
	"encoding/base64"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestLoadBalancer(t *testing.T) {
	//defer leaktest.Check(t)()
	worker := func(req Request) interface{} {
		time.Sleep(20 * time.Millisecond)
		return base64.StdEncoding.EncodeToString([]byte(req.data.(string)))
	}
	b := NewBalancer(worker)
	//b.SetDebug(true)
	go b.Balance()

	var counter = 0
	var counterLock = &sync.RWMutex{}
	for i := 0; i < 2200; i++ {
		var wg = sync.WaitGroup{}
		for j := 0; j < 100; j++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				respChan := b.Add("1234567")
				select {
				case resp := <-respChan:
					if resp.(string) != "MTIzNDU2Nw==" {
						t.Errorf("Resp should be 'MTIzNDU2Nw==', instead of %s", resp.(string))
					}
				case <-time.After(100 * time.Millisecond):
					return
				}

				counterLock.Lock()
				counter++
				counterLock.Unlock()

			}()

		}
		fmt.Println(b.NumOfWorkers())
		wg.Wait()
	}
	//time.Sleep(100*time.Millisecond)
	fmt.Println(counter)
	//b.Interrupt()
}
