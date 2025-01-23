package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"pixare40/hello/handlers"
	"time"
)

func main() {
	l := log.New(os.Stdout, "hello: ", log.LstdFlags)
	hh := handlers.NewHello(l)
	goodbye := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", goodbye)
	sm.Handle("/products/*", handlers.NewProducts(l))

	s := &http.Server{
		Addr: ":8080",
		Handler: sm,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	log.Fatal(s.ListenAndServe())

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	signal := <-sigChan
	l.Println("Received terminate, graceful shutdown", signal)
	tc, cancelFunc := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancelFunc()
	s.Shutdown(tc)
}