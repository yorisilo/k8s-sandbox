// 異なる port で 別の httpd(web server) を立ち上げる
package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	mux1, srv1 := newServer(":8081", 5)
	mux2, srv2 := newServer(":8082", 5)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		mux1.HandleFunc("/", helloHandler)
		srv1.ListenAndServe()
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		mux2.HandleFunc("/", goodbyeHandler)
		srv2.ListenAndServe()
	}()
	wg.Wait()
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func goodbyeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "goodbye\n")
}

func newServer(addr string, timeoutSec int) (*http.ServeMux, *http.Server) {
	mux := http.NewServeMux()
	return mux, &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: time.Duration(timeoutSec) * time.Second,
	}
}
