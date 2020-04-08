package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	w.Header().Set("Connection", "keep-alive")
	time.Sleep(300 * time.Microsecond)
	fmt.Fprintf(w, "mes: %s", time.Since(t))
}

func main() {
	s := http.Server{
		Addr:              ":8888",
		//Handler:           nil,
		//TLSConfig:         nil,
		//ReadTimeout:       0,
		//ReadHeaderTimeout: 0,
		//WriteTimeout:      0,
		//IdleTimeout:       0,
		//MaxHeaderBytes:    0,
		//TLSNextProto:      nil,
		//ConnState:         nil,
		//ErrorLog:          nil,
		//BaseContext:       nil,
		//ConnContext:       nil,
	}
	http.HandleFunc("/", handler)
	s.ListenAndServe()
}
