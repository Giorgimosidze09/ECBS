package jetstream_manager

import (
	"fmt"

	"shared/nats_client"

	"github.com/nats-io/nats.go"
)

const (
	KVBucketWithdrawals = "customer_withdrawals"
	KVBucketExpenses    = "customer_expenses"
	KVAllWebhookData    = "webhook_data"
)

func GetJetStreamContext(natsURL string) (nats.JetStreamContext, error) {
	nc := nats_client.Connect(natsURL)
	if nc == nil {
		return nil, fmt.Errorf("failed to connect to NATS")
	}
	js, err := nc.JetStream()
	if err != nil {
		return nil, fmt.Errorf("failed to get JetStream context: %w", err)
	}
	fmt.Println("Successfully connected to JetStream")
	return js, nil
}

func CreateKVBucket(js nats.JetStreamContext, bucketName string, history uint8) (nats.KeyValue, error) {
	kv, err := js.KeyValue(bucketName)
	if err != nil {
		kvConfig := &nats.KeyValueConfig{
			Bucket:   bucketName,
			Storage:  nats.FileStorage,
			Replicas: 1,
			MaxBytes: 10485760,
			History:  history,
		}
		kv, err = js.CreateKeyValue(kvConfig)
		if err != nil {
			return nil, fmt.Errorf("failed to create KV bucket: %w", err)
		}
	}
	return kv, nil
}
