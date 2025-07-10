package core

import (
	"shared/common/dto"

	repository_charges "database/repository/charges"
)

func ConvertChargesListInput(input dto.UsersListInput) repository_charges.ChargesListParams {
	return repository_charges.ChargesListParams{
		Limit:  int32(input.Limit),
		Offset: int32(input.Offset),
	}
}
func ConvertCharges(input repository_charges.ChargesListRow) *dto.Charges {
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
