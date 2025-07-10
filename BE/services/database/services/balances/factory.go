package core

import (
	"fmt"
	"shared/common/dto"

	repository_balances "database/repository/balances"

	"github.com/jackc/pgx/v5/pgtype"
)

func CreateBalanceParams(input dto.TopUpInput) repository_balances.TopUpBalanceParams {
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
	return repository_balances.TopUpBalanceParams{
		UserID:   pgtype.Int4{Int32: int32(input.UserID), Valid: true},
		Balance:  numeric,
		CardID:   int32(input.CardID),
		RideCost: rideCostNumeric,
	}
}

func ConvertBalanceOutput(input repository_balances.TopUpBalanceRow) *dto.BalanceOutput {
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

func ConvertBalanceListInput(input dto.UsersListInput) repository_balances.BalaneListParams {
	return repository_balances.BalaneListParams{
		Limit:  int32(input.Limit),
		Offset: int32(input.Offset),
	}
}

func ConvertBalanceListOutput(input repository_balances.BalaneListRow) *dto.BalanceOutput {
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

func AddBalanceToCardParams(input dto.PayboxTopupRequest) repository_balances.AddBalanceToCardParams {
	numericAmount := pgtype.Numeric{}
	_ = numericAmount.Scan(input.Amount)
	return repository_balances.AddBalanceToCardParams{
		CardID:  int32(input.CardID),
		Balance: numericAmount,
	}
}
