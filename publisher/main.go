package main

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"publisher/config"
	"strconv"
	"strings"
)

const STREAM = "JETSTREAM_EXAMPLE"
const SUBJECT = "JETSTREAM_SUBJECT.{:trace_id}"

func main() {
	js := config.ConnectJs()

	createStreamIfNotExists(js, STREAM, SUBJECT)

	for i := 0; i < 10; i++ {
		message := make(map[string]string)
		message["msg"] = "hello " + strconv.Itoa(i)

		messageAsByte, _ := json.Marshal(message)
		subject := strings.Replace(SUBJECT, "{:trace_id}", "randomTraceId"+strconv.Itoa(i), 1)

		_, err := js.Publish(subject, messageAsByte)
		if err != nil {
			panic(err)
		}
	}

}

func createStreamIfNotExists(js nats.JetStreamContext, streamName, subject string) {
	stream, err := js.StreamInfo(streamName)
	subject = strings.Replace(subject, "{:trace_id}", "*", 1)
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
