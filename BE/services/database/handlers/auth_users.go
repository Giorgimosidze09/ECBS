package handlers

import (
	"context"
	database "database/db"
	repository_auth_users "database/repository/auth_users"
	service "database/services"
	"shared/common/dto"
	"shared/common/utils"
	subscribe_manager "shared/nats_client/subscribe-manager"
)

var RegisterAuthUser subscribe_manager.Handler = func(data []byte) ([]byte, error) {
	input, err := utils.Decode[dto.RegisterRequest](data)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	q := repository_auth_users.New(database.DB)
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
	q := repository_auth_users.New(database.DB)
	user, err := service.GetAuthUserByUsername(ctx, q, input.Username)
	if err != nil {
		return nil, err
	}
	return utils.Encode(user)
}
