package handlers

import (
	service "database/services"
	"shared/common/dto"
	"shared/common/utils"
	subscribe_manager "shared/nats_client/subscribe-manager"
)

var TopUpBalance subscribe_manager.Handler = func(data []byte) ([]byte, error) {
	input, err := utils.Decode[dto.TopUpInput](data)
	if err != nil {
		return nil, err
	}
	created, err := service.TopUpBalance(input)
	if err != nil {
		return nil, err
	}
	return utils.Encode(created)
}

var BalanceList subscribe_manager.Handler = func(data []byte) ([]byte, error) {
	input, err := utils.Decode[dto.UsersListInput](data)
	if err != nil {
		return nil, err
	}
	totalBalance, err := service.BalanceList(input)
	if err != nil {
		return nil, err
	}
	return utils.Encode(totalBalance)
}

var RideCost subscribe_manager.Handler = func(data []byte) ([]byte, error) {
	input, err := utils.Decode[dto.RideCostInput](data)
	if err != nil {
		return nil, err
	}
	totalBalance, err := service.RideCost(input)
	if err != nil {
		return nil, err
	}
	return utils.Encode(totalBalance)
}

var AddBalanceToCard subscribe_manager.Handler = func(data []byte) ([]byte, error) {
	input, err := utils.Decode[dto.PayboxTopupRequest](data)
	if err != nil {
		return nil, err
	}
	err = service.AddBalanceToCard(input)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

var TotalBalance subscribe_manager.Handler = func(data []byte) ([]byte, error) {
	totalBalance, err := service.TotalBalance()
	if err != nil {
		return nil, err
	}
	return utils.Encode(totalBalance)
}
