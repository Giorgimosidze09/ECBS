package core

import (
	"fmt"
	"shared/common/dto"

	repository_user "database/repository/users"

	"github.com/jackc/pgx/v5/pgtype"
)

func CreateUserParams(input dto.CreateUsersInput) repository_user.CreateUserParams {
	return repository_user.CreateUserParams{
		Name:  input.Name,
		Email: pgtype.Text{String: input.Email, Valid: true},
		Phone: pgtype.Text{String: input.Phone, Valid: true},
	}
}

func CreateDeviceParams(input dto.DevicesInput) repository_user.CreateDeviceParams {
	return repository_user.CreateDeviceParams{
		DeviceID: input.DeviceID,
		Location: pgtype.Text{String: input.Location, Valid: true},
	}
}

func ConvertOutput(input repository_user.User) *dto.UserOutput {
	return &dto.UserOutput{
		ID:    input.ID,
		Name:  input.Name,
		Email: input.Email.String,
		Phone: input.Phone.String,
	}
}

func ConvertDeviceOutput(input repository_user.Device) *dto.DevicesOutput {
	return &dto.DevicesOutput{
		ID:          int(input.ID),
		DeviceID:    input.DeviceID,
		Location:    input.Location.String,
		InstalledAt: input.InstalledAt.Time.Format("2006-01-02 15:04:05"),
		Active:      input.Active.Bool,
	}
}

func CreateCardParams(input dto.AssignCardInput) repository_user.CreateCardParams {
	return repository_user.CreateCardParams{
		CardID:   input.CardID,
		UserID:   int32(input.UserID),
		DeviceID: int32(input.DeviceID),
		Active:   pgtype.Bool{Bool: true, Valid: true},
		Type:     input.Type,
	}
}

func ConvertCardOutput(input repository_user.CreateCardRow) *dto.CardOutput {
	return &dto.CardOutput{
		ID:       int(input.ID),
		UserID:   int(input.UserID),
		CardID:   input.CardID,
		DeviceID: int(input.DeviceID),
		Active:   input.Active.Bool,
	}
}

func CreateBalanceParams(input dto.TopUpInput) repository_user.TopUpBalanceParams {
	numeric := pgtype.Numeric{}
	err := numeric.Scan(fmt.Sprintf("%.2f", input.Balance))
	if err != nil {
		panic(fmt.Sprintf("failed to scan balance: %v", err))
	}
	rideCostNumeric := pgtype.Numeric{}
	err = rideCostNumeric.Scan(fmt.Sprintf("%.2f", input.RideCost))
	if err != nil {
		panic(fmt.Sprintf("failed to scan ride cost: %v", err))
	}
	return repository_user.TopUpBalanceParams{
		UserID:   pgtype.Int4{Int32: int32(input.UserID), Valid: true},
		Balance:  numeric,
		CardID:   int32(input.CardID),
		RideCost: rideCostNumeric,
	}
}

func ConvertBalanceOutput(input repository_user.TopUpBalanceRow) *dto.BalanceOutput {
	var balanceFloat float64
	if input.Balance.Valid {
		balanceFloatStruct, _ := input.Balance.Float64Value()
		balanceFloat = balanceFloatStruct.Float64
	}
	return &dto.BalanceOutput{
		UserID:    int(input.UserID.Int32),
		Balance:   balanceFloat,
		UpdatedAt: input.UpdatedAt.Time.Format("2006-01-02 15:04:05"),
		CardID:    int(input.CardID),
	}
}

func ConvertUserListInput(input dto.UsersListInput) repository_user.UsersListParams {
	return repository_user.UsersListParams{
		Limit:  int32(input.Limit),
		Offset: int32(input.Offset),
	}
}

func ConvertCardsListInput(input dto.UsersListInput) repository_user.CardsListParams {
	return repository_user.CardsListParams{
		Limit:  int32(input.Limit),
		Offset: int32(input.Offset),
	}
}

func ConvertBalanceListInput(input dto.UsersListInput) repository_user.BalaneListParams {
	return repository_user.BalaneListParams{
		Limit:  int32(input.Limit),
		Offset: int32(input.Offset),
	}
}

func ConvertDevicesListInput(input dto.UsersListInput) repository_user.DeviceListParams {
	return repository_user.DeviceListParams{
		Limit:  int32(input.Limit),
		Offset: int32(input.Offset),
	}
}

func ConvertUsersListOutput(input repository_user.UsersListRow) *dto.UsersListOutput {

	return &dto.UsersListOutput{
		ID:           input.ID,
		Name:         input.Name,
		Email:        input.Email.String,
		Phone:        input.Phone.String,
		CardCount:    input.CardCount,
		TotalBalance: float64(input.TotalBalance),
		Total:        int(input.Total),
	}
}

func ConvertCardsListOutput(input repository_user.CardsListRow) *dto.CardOutput {
	return &dto.CardOutput{
		ID:         int(input.ID),
		UserID:     int(input.UserID),
		CardID:     input.CardID,
		Active:     input.Active.Bool,
		AssignedAt: input.AssignedAt.Time.Format("2006-01-02 15:04:05"),
		Total:      int(input.Total),
	}
}

func ConvertBalanceListOutput(input repository_user.BalaneListRow) *dto.BalanceOutput {
	var balance float64
	if input.Balance.Valid {
		if f, err := input.Balance.Float64Value(); err == nil {
			balance = f.Float64
		}
	}

	var rideCost float64
	if input.RideCost.Valid {
		if f, err := input.RideCost.Float64Value(); err == nil {
			rideCost = f.Float64
		}
	}

	return &dto.BalanceOutput{
		UserID:   int(input.UserID.Int32),
		CardID:   int(input.CardID),
		Balance:  balance,
		RideCost: rideCost,
	}
}

func ConvertDevicesListOutput(input repository_user.DeviceListRow) *dto.DevicesOutput {
	return &dto.DevicesOutput{
		ID:          int(input.ID),
		DeviceID:    input.DeviceID,
		Location:    input.Location.String,
		InstalledAt: input.InstalledAt.Time.Format("2006-01-02 15:04:05"),
		Active:      input.Active.Bool,
	}
}

func ConvertChargesListInput(input dto.UsersListInput) repository_user.ChargesListParams {
	return repository_user.ChargesListParams{
		Limit:  int32(input.Limit),
		Offset: int32(input.Offset),
	}
}
func ConvertCharges(input repository_user.ChargesListRow) *dto.Charges {
	var amount float64
	if input.Amount.Valid {
		f, _ := input.Amount.Float64Value()
		amount = f.Float64
	} else {
		amount = 0
	}

	return &dto.Charges{
		ID:          int(input.ID),
		UserID:      int(input.UserID.Int32),
		Amount:      amount,
		Type:        input.Type,
		Description: input.Description.String,
		CreatedAt:   input.CreatedAt.Time.Format("2006-01-02 15:04:05"),
		Total:       int(input.Total),
	}
}

func CreateCardActivationParams(input dto.CardActivation) repository_user.CreateCardActivationParams {
	return repository_user.CreateCardActivationParams{
		CardID:          int32(input.CardID),
		ActivationStart: pgtype.Date{Time: input.ActivationStart, Valid: true},
		ActivationEnd:   pgtype.Date{Time: input.ActivationEnd, Valid: true},
	}
}

func ConvertCardActivationOutput(input repository_user.CardActivation) *dto.CardActivation {
	return &dto.CardActivation{
		ID:              int(input.ID),
		CardID:          int(input.CardID),
		ActivationStart: input.ActivationStart.Time,
		ActivationEnd:   input.ActivationEnd.Time,
	}
}
