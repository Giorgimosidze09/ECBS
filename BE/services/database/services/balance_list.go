package core

import (
	"context"
	database "database/db"
	repository_user "database/repository/users"
	"fmt"
	"log"
	"shared/common/dto"
)

func BalanceList(input dto.UsersListInput) ([]*dto.BalanceOutput, error) {
	ctx := context.Background()

	tx, err := database.DB.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %v", err)
	}
	defer tx.Rollback(ctx)

	q := repository_user.New(tx)

	pramas := ConvertBalanceListInput(input)
	total, err := q.BalaneList(context.Background(), pramas)
	log.Printf("balance list: %v", total)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	var result []*dto.BalanceOutput
	for _, row := range total {
		result = append(result, ConvertBalanceListOutput(row))
	}
	return result, nil
}
