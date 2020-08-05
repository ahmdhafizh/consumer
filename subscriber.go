package main

import (
	"context"
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// Subscriber :nodoc:
type Subscriber struct {
	kConsumer *kafka.Consumer
	topic     []string
	handler   map[string]SubscriberHandler
	config    *SubscribeConfig
}

// Newsubscriber :nodoc:
func Newsubscriber(config *SubscribeConfig) Sub {
	return &Subscriber{
		topic:   make([]string, 0),
		handler: make(map[string]SubscriberHandler),
		config:  config,
	}
}

// RegisterSubscriber :nodoc:
func (s *Subscriber) RegisterSubscriber(topic string, h SubscriberHandler, configs ...SubscribeConfigFunc) error {
	config := SubscribeConfig{}
	for _, c := range configs {
		c(&config)
	}

	kConfigMap := &kafka.ConfigMap{
		"bootstrap.servers": config.Address,
		"group.id":          config.GroupName,
		"auto.offset.reset": config.AutoOffsetReset,
	}

	if config.SASL.Enable {
		kConfigMap.SetKey("sasl.mechanisms", config.SASL.Mechanism)
		kConfigMap.SetKey("security.protocol", config.SASL.SecurityProtocol)
		kConfigMap.SetKey("sasl.username", config.SASL.User)
		kConfigMap.SetKey("sasl.password", config.SASL.Password)
	}

	kConsumer, err := kafka.NewConsumer(kConfigMap)
	if err != nil {
		return err
	}

	s.kConsumer = kConsumer
	s.topic = append(s.topic, topic)
	s.handler[topic] = h

	return nil
}

// RunSubscriber :nodoc:
func (s *Subscriber) RunSubscriber() {

	fmt.Println("Start consumer")
	s.kConsumer.SubscribeTopics(s.topic, nil)

consume:
	for {
		event := s.kConsumer.Poll(100)
		if event == nil {
			continue
		}

		switch evt := event.(type) {
		case *kafka.Message:
			s.handler[*evt.TopicPartition.Topic](context.Background(), evt.Value)
		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v: %v\n", evt.Code(), evt)
			if evt.Code() == kafka.ErrAllBrokersDown {
				break consume
			}
		}
	}

	fmt.Println("Closing consumer")
	s.kConsumer.Close()
}
