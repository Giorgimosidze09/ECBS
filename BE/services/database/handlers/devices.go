package handlers

import (
	service "database/services/devices"
	"shared/common/dto"
	"shared/common/utils"
	subscribe_manager "shared/nats_client/subscribe-manager"
)

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
