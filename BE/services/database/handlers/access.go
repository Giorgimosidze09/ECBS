package handlers

import (
	service "database/services"
	"shared/common/dto"
	"shared/common/utils"
	subscribe_manager "shared/nats_client/subscribe-manager"
)

var SyncAccessLogs subscribe_manager.Handler = func(data []byte) ([]byte, error) {
	input, err := utils.Decode[dto.SyncAccessLogInput](data)
	if err != nil {
		return nil, err
	}
	err = service.SyncAccessLogs(input)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
