package database_client

import (
	"shared/common/dto"
	"shared/common/utils"
)

func (c DatabaseClient) RegisterAuthUser(input dto.RegisterRequest) (*dto.RegisterResponse, error) {
	data, err := utils.Encode(input)
	if err != nil {
		return nil, err
	}
	response := c.Publisher.Request("auth_users.create", data)
	if response.Error != nil {
		return nil, response.Error
	}
	return utils.Decode[*dto.RegisterResponse](response.Data)
}

func (c DatabaseClient) LoginAuthUserHandler(input dto.LoginRequest) (*struct {
	ID           int32  `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`
	Role         string `json:"role"`
	DeviceID     string `json:"device_id"`
}, error) {
	data, err := utils.Encode(input)
	if err != nil {
		return nil, err
	}
	response := c.Publisher.Request("auth_users.get_by_username", data)
	if response.Error != nil {
		return nil, response.Error
	}
	return utils.Decode[*struct {
		ID           int32  `json:"id"`
		Username     string `json:"username"`
		PasswordHash string `json:"password_hash"`
		Role         string `json:"role"`
		DeviceID     string `json:"device_id"`
	}](response.Data)
}
