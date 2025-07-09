package nats_client

import (
	"fmt"
	"sync"

	"github.com/nats-io/nats.go"
)

var (
	once       sync.Once
	connection *nats.Conn
)

func Connect(natsURL string) *nats.Conn {
	once.Do(func() {
		var err error
		connection, err = nats.Connect(natsURL)
		if err != nil {
			panic(fmt.Sprintf("Failed to connect to NATS server at %s: %v", natsURL, err))
		}

		// Check if JetStream is enabled
		_, err = connection.JetStream()
		if err != nil {
			panic(fmt.Sprintf("JetStream is not enabled on NATS server at %s. Please enable JetStream in your NATS server configuration. Error: %v", natsURL, err))
		}
	})

	return connection
}
