package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

var c = &http.Client{
	Transport: &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: time.Minute,
		}).Dial,
		MaxIdleConnsPerHost: 1024,
		TLSHandshakeTimeout: 0 * time.Second,
	},
}

func main() {
	for i := 0; i < 10; i++ {
		req, err := http.NewRequest("GET", "http://localhost:8888", nil)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("Connection", "keep-alive")

		t := time.Now()
		resp, err := c.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(time.Since(t))
		r, err := ioutil.ReadAll(resp.Body)
		fmt.Println(string(r))
		fmt.Println(resp.Header)
		resp.Body.Close()
		time.Sleep(1*time.Second)
	}
}
