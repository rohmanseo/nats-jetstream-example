package config

import "github.com/nats-io/nats.go"

func ConnectJs() nats.JetStreamContext {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}
	js, err := nc.JetStream()
	if err != nil {
		panic(err)
	}
	return js
}
