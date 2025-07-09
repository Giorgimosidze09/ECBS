package subscribe_manager

import (
	"sync"

	"shared/common/utils"
	"shared/nats_client"
	publish_manager "shared/nats_client/publish-manager"

	"github.com/nats-io/nats.go"
)

type SubscribeManager struct {
	Name   string
	Client *nats.Conn
}

var (
	Subscriber SubscribeManager
	once       sync.Once
)

type SubscribeManagerConfig struct {
	Name    string
	NatsURL string
}

type Handler func(data []byte) ([]byte, error)

func InitManager(config SubscribeManagerConfig) *SubscribeManager {
	once.Do(func() {
		Subscriber = SubscribeManager{
			Name:   config.Name,
			Client: nats_client.Connect(config.NatsURL),
		}
	})

	return &Subscriber
}

func (sm *SubscribeManager) RegisterSyncListener(subject string, handler Handler) {
	var wrapper nats.MsgHandler = func(msg *nats.Msg) {
		response := publish_manager.PublishResponse{}

		data, err := handler(msg.Data)
		if err != nil {
			response.ErrorMessage = err.Error()
		} else {
			response.Data = data
		}

		reply, err := utils.Encode(response)
		if err != nil {
			msg.Respond([]byte{})
			return
		}

		msg.Respond(reply)
	}

	sm.Client.Subscribe(sm.Name+"."+subject, wrapper)
}
