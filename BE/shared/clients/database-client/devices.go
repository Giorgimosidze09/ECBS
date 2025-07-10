package database_client

import (
	"shared/common/dto"
	"shared/common/utils"
)

func (c DatabaseClient) CreateDevices(input dto.DevicesInput) (*dto.DevicesOutput, error) {
	data, err := utils.Encode(input)
	if err != nil {
		return nil, err
	}

	response := c.Publisher.Request("devices", data)
	if response.Error != nil {
		return nil, response.Error
	}

	return utils.Decode[*dto.DevicesOutput](response.Data)
}

func (c DatabaseClient) DevicesList(input dto.UsersListInput) ([]*dto.DevicesOutput, error) {
	data, err := utils.Encode(input)
	if err != nil {
		return nil, err
	}

	response := c.Publisher.Request("devices_list", data)
	if response.Error != nil {
		return nil, response.Error
	}

	return utils.Decode[[]*dto.DevicesOutput](response.Data)
}

func (c DatabaseClient) UpdateDevice(input dto.DeviceOutput) (*dto.DeviceOutput, error) {
	data, err := utils.Encode(input)
	if err != nil {
		return nil, err
	}
	response := c.Publisher.Request("update_device", data)
	if response.Error != nil {
		return nil, response.Error
	}
	return utils.Decode[*dto.DeviceOutput](response.Data)
}

func (c DatabaseClient) SoftDeleteDevice(deviceID int) error {
	data, err := utils.Encode(deviceID)
	if err != nil {
		return err
	}
	response := c.Publisher.Request("delete_device", data)
	return response.Error
}

func (c DatabaseClient) GetDeviceByID(deviceID int) (*dto.DeviceOutput, error) {
	data, err := utils.Encode(deviceID)
	if err != nil {
		return nil, err
	}
	response := c.Publisher.Request("get_device_by_id", data)
	if response.Error != nil {
		return nil, response.Error
	}
	return utils.Decode[*dto.DeviceOutput](response.Data)
}
