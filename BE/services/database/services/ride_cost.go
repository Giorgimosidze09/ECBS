package core

import (
	"context"
	database "database/db"
	repository_user "database/repository/users"
	"fmt"
	"log"
	"shared/common/dto"
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
)

func RideCost(input dto.RideCostInput) (*bool, error) {
	ctx := context.Background()

	tx, err := database.DB.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %v", err)
	}
	defer tx.Rollback(ctx)

	q := repository_user.New(tx)

	log.Printf("Received ride cost input: %v", input.RideCost)

	numeric := pgtype.Numeric{}
	stringValue := strconv.FormatFloat(input.RideCost, 'f', -1, 64)
	log.Printf("Converted to string: %v", stringValue)

	if err := numeric.Scan(stringValue); err != nil {
		return nil, fmt.Errorf("failed to convert ride cost to numeric: %w", err)
	}

	err = q.CostOfRide(ctx, numeric)

	log.Printf("RIDE COST CHANGED")
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	result := true
	return &result, nil
}
