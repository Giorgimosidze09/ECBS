package handlers

import (
	"context"
	"shared/common/dto"
	"shared/common/utils"
	subscribe_manager "shared/nats_client/subscribe-manager"

	database "database/db"
	repository_users "database/repository/users"
	service "database/services"
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

var CreateDevices subscribe_manager.Handler = func(data []byte) ([]byte, error) {
	input, err := utils.Decode[dto.DevicesInput](data)
	if err != nil {
		return nil, err
	}

	created, err := service.CreateDevices(input)
	if err != nil {
		return nil, err
	}

	return utils.Encode(created)
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

var Authorization subscribe_manager.Handler = func(data []byte) ([]byte, error) {
	input, err := utils.Decode[dto.AuthorizedInput](data)
	if err != nil {
		return nil, err
	}

	created, err := service.Authorization(input)
	if err != nil {
		return nil, err
	}

	return utils.Encode(created)
}

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

var CountUsers subscribe_manager.Handler = func(data []byte) ([]byte, error) {
	counted, err := service.CountUsers()
	if err != nil {
		return nil, err
	}

	return utils.Encode(counted)
}

var CountCards subscribe_manager.Handler = func(data []byte) ([]byte, error) {

	counted, err := service.CountCards()
	if err != nil {
		return nil, err
	}

	return utils.Encode(counted)
}

var TotalBalance subscribe_manager.Handler = func(data []byte) ([]byte, error) {

	totalBalance, err := service.TotalBalance()
	if err != nil {
		return nil, err
	}

	return utils.Encode(totalBalance)
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

var DevicesList subscribe_manager.Handler = func(data []byte) ([]byte, error) {

	input, err := utils.Decode[dto.UsersListInput](data)
	if err != nil {
		return nil, err
	}
	totalBalance, err := service.DevicesList(input)
	if err != nil {
		return nil, err
	}

	return utils.Encode(totalBalance)
}

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

var UpdateDevice subscribe_manager.Handler = func(data []byte) ([]byte, error) {
	input, err := utils.Decode[dto.DeviceOutput](data)
	if err != nil {
		return nil, err
	}
	err = service.UpdateDevice(input)
	if err != nil {
		return nil, err
	}
	return utils.Encode(input)
}

var SoftDeleteDevice subscribe_manager.Handler = func(data []byte) ([]byte, error) {
	id, err := utils.Decode[int](data)
	if err != nil {
		return nil, err
	}
	err = service.SoftDeleteDevice(int32(id))
	if err != nil {
		return nil, err
	}
	return nil, nil
}

var GetDeviceByID subscribe_manager.Handler = func(data []byte) ([]byte, error) {
	id, err := utils.Decode[int](data)
	if err != nil {
		return nil, err
	}
	device, err := service.GetDeviceByID(int32(id))
	if err != nil {
		return nil, err
	}
	return utils.Encode(device)
}

var RegisterAuthUser subscribe_manager.Handler = func(data []byte) ([]byte, error) {
	input, err := utils.Decode[dto.RegisterRequest](data)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	q := repository_users.New(database.DB)
	created, err := service.CreateAuthUser(ctx, q, input.Username, input.Password, input.Role)
	if err != nil {
		return nil, err
	}
	return utils.Encode(created)
}

var LoginAuthUserHandler subscribe_manager.Handler = func(data []byte) ([]byte, error) {
	input, err := utils.Decode[dto.LoginRequest](data)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	q := repository_users.New(database.DB)
	user, err := service.GetAuthUserByUsername(ctx, q, input.Username)
	if err != nil {
		return nil, err
	}
	return utils.Encode(user)
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
