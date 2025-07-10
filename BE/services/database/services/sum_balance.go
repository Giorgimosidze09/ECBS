package core

import (
	"context"
	database "database/db"
	repository_user "database/repository/users"
	"fmt"
	"log"
	"shared/common/dto"

	"github.com/jackc/pgx/v5/pgtype"
)

func SumBalance(input dto.CustomerSumBalanceRequest) (*dto.CustomerSumBalanceResponse, error) {
	ctx := context.Background()

	tx, err := database.DB.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %v", err)
	}
	defer tx.Rollback(ctx)

	q := repository_user.New(tx)

	total, err := q.GetSumBalanceByDeviceID(context.Background(), input.DeviceID)
	log.Printf("balance list: %v", total)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	numeric, ok := total.(pgtype.Numeric)
	if !ok {
		return nil, fmt.Errorf("unexpected type for total: %T", total)
	}

	var totalBalance float64
	if numeric.Valid {
		f, err := numeric.Float64Value()
		if err != nil {
			return nil, fmt.Errorf("failed to convert pgtype.Numeric to float64: %v", err)
		}
		totalBalance = f.Float64
	} else {
		totalBalance = 0
	}

	response := &dto.CustomerSumBalanceResponse{
		TotalBalance: totalBalance,
	}
	return response, nil
}
