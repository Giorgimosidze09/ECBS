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

func (c DatabaseClient) SyncAuthorizedAccess(input dto.AuthorizedInput) ([]dto.AuthorizedAccessSyncDTO, error) {
	data, err := utils.Encode(input)
	if err != nil {
		return nil, err
	}

	response := c.Publisher.Request("authorized", data)
	if response.Error != nil {
		return nil, response.Error
	}

	return utils.Decode[[]dto.AuthorizedAccessSyncDTO](response.Data)
}

func (c DatabaseClient) SyncAccessLogs(input dto.SyncAccessLogInput) error {
	data, err := utils.Encode(input)
	if err != nil {
		return err
	}

	response := c.Publisher.Request("access_logs", data) // more specific subject
	if response.Error != nil {
		return response.Error
	}

	return nil
}

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
