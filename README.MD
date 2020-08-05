# Kafka Confluent Consumer

Create config:
```go
subConfig := NewSubscribeConfig()
```

Create new instance:
```go
sub := Newsubscriber(subConfig)
```

Create handler:
```go
func ExampleHandler(ctx context.Context, payload SubscribePayload) error {
	fmt.Println("enter example handler")
	fmt.Println(payload)

	return nil
}
```

Register Subscriber:
```go
sub.RegisterSubscriber("example_topic", ExampleHandler)
sub.RegisterSubscriber("example_topic_2", ExampleHandler2)
sub.RegisterSubscriber("example_topic_3", ExampleHandler3)
```

Listen:
```go
go sub.RunSubscriber()
```
