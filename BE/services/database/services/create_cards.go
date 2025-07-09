package core

import (
	"context"
	database "database/db"
	repository_user "database/repository/users"
	"fmt"
	"shared/common/dto"

	"github.com/jackc/pgx/v5/pgtype"
)

func CreateCard(input dto.AssignCardInput) (*dto.CardOutput, error) {
	ctx := context.Background()

	tx, err := database.DB.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %v", err)
	}
	defer tx.Rollback(ctx)

	q := repository_user.New(tx)

	dbParams := CreateCardParams(input)

	card, err := q.CreateCard(ctx, dbParams)
	if err != nil {
		return nil, err
	}

	// If type is activation and dates are provided, insert into card_activations
	if input.Type == "activation" && input.ActivationStart != "" && input.ActivationEnd != "" {
		activationParams := repository_user.CreateCardActivationParams{
			CardID:          card.ID,
			ActivationStart: parseDate(input.ActivationStart),
			ActivationEnd:   parseDate(input.ActivationEnd),
		}
		_, err := q.CreateCardActivation(ctx, activationParams)
		if err != nil {
			return nil, fmt.Errorf("failed to create card activation: %v", err)
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return ConvertCardOutput(card), nil
}

// parseDate parses a YYYY-MM-DD string to pgtype.Date
func parseDate(dateStr string) pgtype.Date {
	var d pgtype.Date
	_ = d.Scan(dateStr)
	return d
}
