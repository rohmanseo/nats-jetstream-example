package main

import (
	"consumer/config"
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"sync"
)

const STREAM = "JETSTREAM_EXAMPLE"
const SUBJECT = "JETSTREAM_SUBJECT.*"

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	js := config.ConnectJs()

	createStreamIfNotExists(js, STREAM, SUBJECT)

	message := make(map[string]string)

	_, err := js.QueueSubscribe(SUBJECT, "1-for-1", func(msg *nats.Msg) {
		json.Unmarshal(msg.Data, &message)
		fmt.Println("MESSAGE RECEIVED---")
		fmt.Println(message)

	})
	if err != nil {
		panic(err)
	}
	wg.Wait()

}

func createStreamIfNotExists(js nats.JetStreamContext, streamName, subject string) {
	stream, err := js.StreamInfo(streamName)

	if stream == nil {
		_, err = js.AddStream(&nats.StreamConfig{
			Name:     streamName,
			Subjects: []string{subject},
		})

		if err != nil {
			panic(err)
		}
	}
}
