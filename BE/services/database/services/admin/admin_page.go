package core

import (
	"context"
	database "database/db"
	repository_balances "database/repository/balances"
	repository_cards "database/repository/cards"
	repository_user "database/repository/users"
	"fmt"
	"log"
	"shared/common/dto"

	"github.com/jackc/pgx/v5/pgtype"
)

func CountCards() (*dto.CountOutput, error) {
	ctx := context.Background()

	tx, err := database.DB.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %v", err)
	}
	defer tx.Rollback(ctx)

	q := repository_cards.New(tx)

	countCards, err := q.CountCards(context.Background())

	log.Printf("Created cards: %v", countCards)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return &dto.CountOutput{Count: int(countCards)}, nil
}

func CountUsers() (*dto.CountOutput, error) {
	ctx := context.Background()

	tx, err := database.DB.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %v", err)
	}
	defer tx.Rollback(ctx)

	q := repository_user.New(tx)

	countCards, err := q.CountUsers(context.Background())

	log.Printf("Created user: %v", countCards)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return &dto.CountOutput{Count: int(countCards)}, nil
}

func TotalBalance() (*dto.TotalBalanceOutput, error) {
	ctx := context.Background()

	tx, err := database.DB.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %v", err)
	}
	defer tx.Rollback(ctx)

	balances := repository_balances.New(tx)

	total, err := balances.TotalBalance(context.Background())
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	numericTotal, ok := total.(pgtype.Numeric)
	if !ok {
		return nil, fmt.Errorf("unexpected type for total: %T", total)
	}

	f, err := numericTotal.Float64Value()
	if err != nil {
		return nil, fmt.Errorf("failed to convert total to float64: %v", err)
	}
	return &dto.TotalBalanceOutput{Total: f.Float64}, nil
}
