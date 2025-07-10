package database_client

import (
	"shared/common/dto"
	"shared/common/utils"
)

func (c DatabaseClient) CountUsers() (*dto.CountOutput, error) {
	response := c.Publisher.Request("count_users", nil)
	if response.Error != nil {
		return nil, response.Error
	}
	return utils.Decode[*dto.CountOutput](response.Data)
}

func (c DatabaseClient) CountCards() (*dto.CountOutput, error) {
	response := c.Publisher.Request("count_cards", nil)
	if response.Error != nil {
		return nil, response.Error
	}
	return utils.Decode[*dto.CountOutput](response.Data)
}

func (c DatabaseClient) TotalBalance() (*dto.TotalBalanceOutput, error) {
	response := c.Publisher.Request("total_balance", nil)
	if response.Error != nil {
		return nil, response.Error
	}
	return utils.Decode[*dto.TotalBalanceOutput](response.Data)
}

func (c DatabaseClient) GetCharges(input dto.UsersListInput) ([]*dto.Charges, error) {
	data, err := utils.Encode(input)
	if err != nil {
		return nil, err
	}

	response := c.Publisher.Request("charges_list", data)
	if response.Error != nil {
		return nil, response.Error
	}

	return utils.Decode[[]*dto.Charges](response.Data)
}
