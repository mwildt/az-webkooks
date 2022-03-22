package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	serveHttp()
}

func serveHttp() {

	httpHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[http] %s %s\n", r.Method, r.URL.Path)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Hallo schöne welt"))
	})

	fmt.Printf("Listening for incomming http-requests on port 80\n")
	if err := http.ListenAndServe(":80", httpHandler); err != nil {
		log.Fatal(err)
	}
}
