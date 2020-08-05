package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	subConfig := NewSubscribeConfig()
	sub := Newsubscriber(subConfig)

	sub.RegisterSubscriber("example_topic", ExampleHandler)

	go sub.RunSubscriber()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-done
	log.Print("Server Stopped")
}

// ExampleHandler :nodoc:
func ExampleHandler(ctx context.Context, payload SubscribePayload) error {
	fmt.Println("enter example handler")
	fmt.Println(payload)

	return nil
}
