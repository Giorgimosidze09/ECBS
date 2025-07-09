package subscribe_manager

import (
	"log"

	jetstream_manager "shared/nats_client/jetstream-manager"

	"github.com/nats-io/nats.go"
)

type KVSubscribeConfig struct {
	Name    string
	NatsURL string
}

func SubscribeToKV(config KVSubscribeConfig, callback func([]byte)) {
	js, err := jetstream_manager.GetJetStreamContext(config.NatsURL)
	if err != nil {
		log.Fatalf("Failed to get JetStream context: %v", err)
	}

	kv, err := jetstream_manager.CreateKVBucket(js, config.Name, 50)
	if err != nil {
		log.Fatalf("Failed to ensure KV bucket %s exists: %v", config.Name, err)
	}

	watcher, err := kv.WatchAll(nats.UpdatesOnly())
	if err != nil {
		log.Fatalf("Failed to watch KV store %s: %v", config.Name, err)
	}
	defer watcher.Stop()

	log.Printf("Listening for KV updates on stream: %s", config.Name)

	for {
		update := <-watcher.Updates()
		if update == nil {
			continue
		}

		callback(update.Value())
	}
}
