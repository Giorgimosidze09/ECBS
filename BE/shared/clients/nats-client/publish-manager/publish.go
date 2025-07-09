package publish_manager

import (
	"errors"
	"sync"
	"time"

	"shared/common/utils"
	"shared/nats_client"

	"github.com/nats-io/nats.go"
)

type PublishManager struct {
	Name   string
	Client *nats.Conn
}

type PublishManagerConfig struct {
	Name    string
	NatsURL string
}

type PublishResponse struct {
	Data         []byte `json:"data"`
	ErrorMessage string `json:"error"`
	Error        error
}

var (
	sharedClient *nats.Conn
	once         sync.Once
)

func (m PublishManager) Request(subject string, data []byte) PublishResponse {
	msg, err := m.Client.Request(m.Name+"."+subject, data, 5*time.Second)
	if err != nil {
		return PublishResponse{
			Data:  nil,
			Error: err,
		}
	}

	res, err := utils.Decode[PublishResponse](msg.Data)
	if err != nil {
		return PublishResponse{
			Data:  nil,
			Error: err,
		}
	}

	if res.ErrorMessage != "" {
		return PublishResponse{
			Data:  nil,
			Error: errors.New(res.ErrorMessage),
		}
	}

	return res
}

func InitPublishManager(config PublishManagerConfig) *PublishManager {
	once.Do(func() {
		sharedClient = nats_client.Connect(config.NatsURL)
	})

	return &PublishManager{
		Name:   config.Name,
		Client: sharedClient,
	}
}
