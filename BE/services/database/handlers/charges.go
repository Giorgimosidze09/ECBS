package handlers

import (
	service "database/services"
	"shared/common/dto"
	"shared/common/utils"
	subscribe_manager "shared/nats_client/subscribe-manager"
)

var ChargesList subscribe_manager.Handler = func(data []byte) ([]byte, error) {
	input, err := utils.Decode[dto.UsersListInput](data)
	if err != nil {
		return nil, err
	}
	totalBalance, err := service.ChargesList(input)
	if err != nil {
		return nil, err
	}
	return utils.Encode(totalBalance)
}
