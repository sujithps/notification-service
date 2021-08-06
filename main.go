package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	http.HandleFunc("/", requestHandler)
	go func() {
		http.ListenAndServe(":8888", nil)
	}()

	defer waitTillInterrupt()
}

func waitTillInterrupt() {
	onApplicationExit := make(chan os.Signal, 1)
	signal.Notify(onApplicationExit, os.Interrupt, os.Kill, syscall.SIGTERM)

	<- onApplicationExit
	fmt.Println("Shutting down the system")
}

func requestHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Received a request")

}
