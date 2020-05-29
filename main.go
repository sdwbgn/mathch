package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
)

var ETag string

func main() {
	ETag = strconv.Itoa(int(Init()))
	test()
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	mux := http.NewServeMux()
	mux.HandleFunc("/get", NewRiddle)
	mux.HandleFunc("/check", CheckSolution)
	mux.HandleFunc("/", Gzip(LoadHTML))
	server := &http.Server{
		Addr:    ":801",
		Handler: mux,
	}
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		println("Starting server...")
		err := server.ListenAndServe()
		if err != nil {
			fmt.Printf("%v\n", err)
		}
	}()
	go func() {
		defer wg.Done()
		_ = <-signals
		println("Stopping server...")
		err := server.Shutdown(context.Background())
		if err != nil {
			fmt.Printf("%v\n", err)
		}
	}()
	wg.Wait()
}
