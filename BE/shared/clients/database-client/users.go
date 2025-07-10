package database_client

import (
	"shared/common/dto"
	"shared/common/utils"
)

func (c DatabaseClient) CreateUser(input dto.CreateUsersInput) (*dto.UserOutput, error) {
	data, err := utils.Encode(input)
	if err != nil {
		return nil, err
	}

	response := c.Publisher.Request("create", data)
	if response.Error != nil {
		return nil, response.Error
	}

	return utils.Decode[*dto.UserOutput](response.Data)
}

func (c DatabaseClient) GetUsersList(input dto.UsersListInput) ([]*dto.UsersListOutput, error) {
	data, err := utils.Encode(input)
	if err != nil {
		return nil, err
	}

	response := c.Publisher.Request("users_list", data)
	if response.Error != nil {
		return nil, response.Error
	}

	return utils.Decode[[]*dto.UsersListOutput](response.Data)
}

func (c DatabaseClient) UpdateUser(input dto.UserOutput) (*dto.UserOutput, error) {
	data, err := utils.Encode(input)
	if err != nil {
		return nil, err
	}
	response := c.Publisher.Request("update_user", data)
	if response.Error != nil {
		return nil, response.Error
	}
	return utils.Decode[*dto.UserOutput](response.Data)
}

func (c DatabaseClient) SoftDeleteUser(userID int) error {
	data, err := utils.Encode(userID)
	if err != nil {
		return err
	}
	response := c.Publisher.Request("delete_user", data)
	return response.Error
}

func (c DatabaseClient) GetUserByID(userID int) (*dto.UserOutput, error) {
	data, err := utils.Encode(userID)
	if err != nil {
		return nil, err
	}
	response := c.Publisher.Request("get_user_by_id", data)
	if response.Error != nil {
		return nil, response.Error
	}
	return utils.Decode[*dto.UserOutput](response.Data)
}

func (c DatabaseClient) SumBalance(input dto.CustomerSumBalanceRequest) (*dto.CustomerSumBalanceResponse, error) {
	data, err := utils.Encode(input)
	if err != nil {
		return nil, err
	}

	response := c.Publisher.Request("sum_balance", data)
	if response.Error != nil {
		return nil, response.Error
	}

	return utils.Decode[*dto.CustomerSumBalanceResponse](response.Data)
}
