package handlers

import (
	service "database/services"
	"shared/common/dto"
	"shared/common/utils"
	subscribe_manager "shared/nats_client/subscribe-manager"
)

var CreateCard subscribe_manager.Handler = func(data []byte) ([]byte, error) {
	input, err := utils.Decode[dto.AssignCardInput](data)
	if err != nil {
		return nil, err
	}
	created, err := service.CreateCard(input)
	if err != nil {
		return nil, err
	}
	return utils.Encode(created)
}

var CardsList subscribe_manager.Handler = func(data []byte) ([]byte, error) {
	input, err := utils.Decode[dto.UsersListInput](data)
	if err != nil {
		return nil, err
	}
	totalBalance, err := service.CardsList(input)
	if err != nil {
		return nil, err
	}
	return utils.Encode(totalBalance)
}

var AddCardActivation subscribe_manager.Handler = func(data []byte) ([]byte, error) {
	input, err := utils.Decode[dto.CardActivation](data)
	if err != nil {
		return nil, err
	}
	created, err := service.AddCardActivation(input)
	if err != nil {
		return nil, err
	}
	return utils.Encode(created)
}

var ValidateCard subscribe_manager.Handler = func(data []byte) ([]byte, error) {
	input, err := utils.Decode[dto.ValidateCardInput](data)
	if err != nil {
		return nil, err
	}
	created, err := service.ValidateCard(input)
	if err != nil {
		return nil, err
	}
	return utils.Encode(created)
}

var UpdateCard subscribe_manager.Handler = func(data []byte) ([]byte, error) {
	input, err := utils.Decode[dto.CardOutput](data)
	if err != nil {
		return nil, err
	}
	err = service.UpdateCard(input)
	if err != nil {
		return nil, err
	}
	return utils.Encode(input)
}

var SoftDeleteCard subscribe_manager.Handler = func(data []byte) ([]byte, error) {
	id, err := utils.Decode[int](data)
	if err != nil {
		return nil, err
	}
	err = service.SoftDeleteCard(int32(id))
	if err != nil {
		return nil, err
	}
	return nil, nil
}

var GetCardByID subscribe_manager.Handler = func(data []byte) ([]byte, error) {
	id, err := utils.Decode[int](data)
	if err != nil {
		return nil, err
	}
	card, err := service.GetCardByID(int32(id))
	if err != nil {
		return nil, err
	}
	return utils.Encode(card)
}

var CountCards subscribe_manager.Handler = func(data []byte) ([]byte, error) {
	counted, err := service.CountCards()
	if err != nil {
		return nil, err
	}
	return utils.Encode(counted)
}
