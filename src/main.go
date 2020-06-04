package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
)

var ETag string

func main() {
	var port string
	if len(os.Args[1:]) == 0 {
		port = os.Getenv("PORT")
		if port == "" {
			port = "801"
		}
	} else if len(os.Args[1:]) != 1 {
		log.Panic("Invalid Arguments")
	} else {
		_, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Panic("Port is not numeric value")
		}
		port = os.Args[1]
	}
	ETag = strconv.Itoa(int(Init()))
	//test()
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	mux := http.NewServeMux()
	mux.HandleFunc("/get", Gzip(NewRiddle))
	mux.HandleFunc("/check", Gzip(CheckSolution))
	mux.HandleFunc("/", Gzip(LoadHTML))
	server := &http.Server{
		Addr:    ":" + port,
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
