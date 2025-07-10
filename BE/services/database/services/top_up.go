package core

import (
	"context"
	database "database/db"
	repository_balances "database/repository/balances"
	"fmt"
	"log"
	"shared/common/dto"
)

func TopUpBalance(input dto.TopUpInput) (*dto.BalanceOutput, error) {
	ctx := context.Background()

	tx, err := database.DB.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %v", err)
	}
	defer tx.Rollback(ctx)

	q := repository_balances.New(tx)

	dbParams := CreateBalanceParams(input)

	user, err := q.TopUpBalance(context.Background(), dbParams)

	log.Printf("Created user: %v", user)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return ConvertBalanceOutput(user), nil
}
