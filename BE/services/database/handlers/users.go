package handlers

import (
	service "database/services"
	"shared/common/dto"
	"shared/common/utils"
	subscribe_manager "shared/nats_client/subscribe-manager"
)

var CreateUser subscribe_manager.Handler = func(data []byte) ([]byte, error) {
	input, err := utils.Decode[dto.CreateUsersInput](data)
	if err != nil {
		return nil, err
	}
	created, err := service.CreateUser(input)
	if err != nil {
		return nil, err
	}
	return utils.Encode(created)
}

var UsersList subscribe_manager.Handler = func(data []byte) ([]byte, error) {
	input, err := utils.Decode[dto.UsersListInput](data)
	if err != nil {
		return nil, err
	}
	totalBalance, err := service.UsersList(input)
	if err != nil {
		return nil, err
	}
	return utils.Encode(totalBalance)
}

var UpdateUser subscribe_manager.Handler = func(data []byte) ([]byte, error) {
	input, err := utils.Decode[dto.UserOutput](data)
	if err != nil {
		return nil, err
	}
	err = service.UpdateUser(input)
	if err != nil {
		return nil, err
	}
	return utils.Encode(input)
}

var SoftDeleteUser subscribe_manager.Handler = func(data []byte) ([]byte, error) {
	id, err := utils.Decode[int](data)
	if err != nil {
		return nil, err
	}
	err = service.SoftDeleteUser(int32(id))
	if err != nil {
		return nil, err
	}
	return nil, nil
}

var GetUserByID subscribe_manager.Handler = func(data []byte) ([]byte, error) {
	id, err := utils.Decode[int](data)
	if err != nil {
		return nil, err
	}
	user, err := service.GetUserByID(int32(id))
	if err != nil {
		return nil, err
	}
	return utils.Encode(user)
}

var SumBalance subscribe_manager.Handler = func(data []byte) ([]byte, error) {
	input, err := utils.Decode[dto.CustomerSumBalanceRequest](data)
	if err != nil {
		return nil, err
	}
	created, err := service.SumBalance(input)
	if err != nil {
		return nil, err
	}
	return utils.Encode(created)
}

var CountUsers subscribe_manager.Handler = func(data []byte) ([]byte, error) {
	counted, err := service.CountUsers()
	if err != nil {
		return nil, err
	}
	return utils.Encode(counted)
}
