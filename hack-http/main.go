package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	var wg = new(sync.WaitGroup)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			t := time.Now()
			req, err := http.NewRequest("GET", "http://192.168.5.27:3306", nil)
			if err != nil {
				log.Fatal(err)
			}
			client := http.DefaultClient
			for j := 0; j < 100; j++ {
				resp, err := client.Do(req)
				if err != nil {
					log.Fatal(err)
				}
				resp.Body.Close()
			}
			fmt.Printf("100 done %v\n", time.Since(t))
			wg.Done()
		}()
	}

	wg.Wait()

}
