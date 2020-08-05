package main

import "context"

// Sub is an interface used for asynchronous messaging.
type Sub interface {
	RegisterSubscriber(topic string, h SubscriberHandler, configs ...SubscribeFunc) error
	RunSubscriber()
}

// SubscriberHandler represents handler function for incoming message
type SubscriberHandler func(context.Context, SubscribePayload) error

// SubscribeFunc :nodoc:
type SubscribeFunc func(*SubscribeConfig)

// SubscribePayload payload json data from event message
type SubscribePayload interface{}
