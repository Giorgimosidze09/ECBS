package core

import (
	"context"
	database "database/db"
	repository_cards "database/repository/cards"
	"fmt"
	"shared/common/dto"
)

func AddCardActivation(input dto.CardActivation) (*dto.CardActivation, error) {
	ctx := context.Background()
	tx, err := database.DB.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %v", err)
	}
	defer tx.Rollback(ctx)

	q := repository_cards.New(tx)
	params := CreateCardActivationParams(input)
	created, err := q.CreateCardActivation(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to insert card activation: %v", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ConvertCardActivationOutput(created), nil
}
