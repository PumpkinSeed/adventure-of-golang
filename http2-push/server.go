package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/push", handlePush)
	http.HandleFunc("/nopush", handle)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handlePush(w http.ResponseWriter, r *http.Request) {
	// Unlock HTTP/2 server push
	p, ok := w.(http.Pusher)
	fmt.Println("Push: ", ok)
	if ok {
		fmt.Println("Push")
		p.Push("/static/style.css", nil)
	}
	http.ServeFile(w, r, "./static/index.html")
}

func handle(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}
