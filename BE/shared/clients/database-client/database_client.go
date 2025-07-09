package database_client

import (
	"log"
	js "shared/nats_client/jetstream-manager"
	pm "shared/nats_client/publish-manager"
	"sync"

	"github.com/nats-io/nats.go"
)

type DatabaseClient struct {
	Publisher *pm.PublishManager
	JetStream nats.JetStreamContext
}

var (
	Client DatabaseClient
	once   sync.Once
)

func SetupClient(natsURL string) *DatabaseClient {
	once.Do(func() {
		Client = DatabaseClient{
			Publisher: pm.InitPublishManager(pm.PublishManagerConfig{
				Name:    "service.database",
				NatsURL: natsURL,
			}),
			JetStream: func() nats.JetStreamContext {
				jsCtx, err := js.GetJetStreamContext(natsURL)
				if err != nil {
					log.Fatalf("Error getting JetStream context: %v", err)
				}
				return jsCtx
			}(),
		}
	})

	return &Client
}
