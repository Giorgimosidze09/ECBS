package database_client

import (
	"shared/common/dto"
	"shared/common/utils"
)

func (c DatabaseClient) CreateCard(input dto.AssignCardInput) (*dto.CardOutput, error) {
	data, err := utils.Encode(input)
	if err != nil {
		return nil, err
	}

	response := c.Publisher.Request("create_card", data)
	if response.Error != nil {
		return nil, response.Error
	}

	return utils.Decode[*dto.CardOutput](response.Data)
}

func (c DatabaseClient) GetCardsList(input dto.UsersListInput) ([]*dto.CardOutput, error) {
	data, err := utils.Encode(input)
	if err != nil {
		return nil, err
	}

	response := c.Publisher.Request("cards_list", data)
	if response.Error != nil {
		return nil, response.Error
	}

	return utils.Decode[[]*dto.CardOutput](response.Data)
}

func (c DatabaseClient) UpdateCard(input dto.CardOutput) (*dto.CardOutput, error) {
	data, err := utils.Encode(input)
	if err != nil {
		return nil, err
	}
	response := c.Publisher.Request("update_card", data)
	if response.Error != nil {
		return nil, response.Error
	}
	return utils.Decode[*dto.CardOutput](response.Data)
}

func (c DatabaseClient) SoftDeleteCard(cardID int) error {
	data, err := utils.Encode(cardID)
	if err != nil {
		return err
	}
	response := c.Publisher.Request("delete_card", data)
	return response.Error
}

func (c DatabaseClient) GetCardByID(cardID int) (*dto.CardOutput, error) {
	data, err := utils.Encode(cardID)
	if err != nil {
		return nil, err
	}
	response := c.Publisher.Request("get_card_by_id", data)
	if response.Error != nil {
		return nil, response.Error
	}
	return utils.Decode[*dto.CardOutput](response.Data)
}

func (c DatabaseClient) TopUpBalance(input dto.TopUpInput) (*dto.BalanceOutput, error) {
	data, err := utils.Encode(input)
	if err != nil {
		return nil, err
	}

	response := c.Publisher.Request("top_up", data)
	if response.Error != nil {
		return nil, response.Error
	}

	return utils.Decode[*dto.BalanceOutput](response.Data)
}

func (c DatabaseClient) AddCardActivation(input dto.CardActivation) (*dto.CardActivation, error) {
	data, err := utils.Encode(input)
	if err != nil {
		return nil, err
	}
	response := c.Publisher.Request("add_card_activation", data)
	if response.Error != nil {
		return nil, response.Error
	}
	return utils.Decode[*dto.CardActivation](response.Data)
}

func (c DatabaseClient) ValidateCard(input dto.ValidateCardInput) (*dto.ValidateCardOutput, error) {
	data, err := utils.Encode(input)
	if err != nil {
		return nil, err
	}

	response := c.Publisher.Request("validate", data)
	if response.Error != nil {
		return nil, response.Error
	}

	return utils.Decode[*dto.ValidateCardOutput](response.Data)
}

func (c DatabaseClient) BalanceList(input dto.UsersListInput) ([]*dto.BalanceOutput, error) {
	data, err := utils.Encode(input)
	if err != nil {
		return nil, err
	}

	response := c.Publisher.Request("balande_list", data)
	if response.Error != nil {
		return nil, response.Error
	}

	return utils.Decode[[]*dto.BalanceOutput](response.Data)
}

func (c DatabaseClient) AddBalanceToCard(input dto.PayboxTopupRequest) error {
	data, err := utils.Encode(input)
	if err != nil {
		return err
	}

	response := c.Publisher.Request("add_balance_to_card", data)
	if response.Error != nil {
		return response.Error
	}

	return nil
}

func (c DatabaseClient) ChangeRideCost(input dto.RideCostInput) (*bool, error) {
	data, err := utils.Encode(input)
	if err != nil {
		return nil, err
	}

	response := c.Publisher.Request("ride_cost", data)
	if response.Error != nil {
		return nil, response.Error
	}

	return utils.Decode[*bool](response.Data)
}
