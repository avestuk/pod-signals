package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	stopOnTerm := os.Getenv("STOP_ON_TERM")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM)

	done := make(chan bool)
	timer := time.NewTicker(2 * time.Second)

	go func() {
		for {
			select {
			case sig := <-sigs:
				fmt.Printf("got signal: %s\n", sig)

				if stopOnTerm != "true" {
					fmt.Println("nbd gonna carry on")
					continue
				}
				done <- true
			case <-timer.C:
				fmt.Println("doing stuff")
			}
		}
	}()

	<-done
	fmt.Println("aite I'm done")
}
