package core

import (
	"context"
	database "database/db"
	repository_user "database/repository/users"
	"fmt"
	"log"
	"shared/common/dto"
)

func AddBalanceToCard(input dto.PayboxTopupRequest) error {
	ctx := context.Background()

	tx, err := database.DB.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %v", err)
	}
	defer tx.Rollback(ctx)

	q := repository_user.New(tx)

	pramas := AddBalanceToCardParams(input)
	total, err := q.AddBalanceToCard(context.Background(), pramas)
	log.Printf("balance list: %v", total)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}
