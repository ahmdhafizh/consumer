package main

// SubscribeConfig :nodoc:
type SubscribeConfig struct {
	GroupName string
	Address   string
	// SASL based authentication with broker. While there are multiple SASL authentication methods
	// the current implementation is limited to plaintext (SASL/PLAIN) authentication
	SASL struct {
		// Whether or not to use SASL authentication when connecting to the broker
		// (defaults to false).
		Enable bool
		// SASLMechanism is the name of the enabled SASL mechanism.
		// Possible values: OAUTHBEARER, PLAIN (defaults to PLAIN).
		Mechanism string
		//username and password for SASL/PLAIN  or SASL/SCRAM authentication
		User     string
		Password string
		// SecurityProtocol SASL_SSL
		SecurityProtocol string
	}
	AutoOffsetReset string
}

var defaultGroupName = "default_group_name"
var defaultAddress = "127.0.0.1:9092"
var defaultAutoOffsetReset = "earliest"

// NewSubscribeConfig :nodoc:
func NewSubscribeConfig() *SubscribeConfig {
	c := &SubscribeConfig{}
	c.Address = defaultAddress
	c.GroupName = defaultGroupName
	c.AutoOffsetReset = defaultAutoOffsetReset

	return c
}
