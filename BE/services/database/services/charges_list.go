package core

import (
	"context"
	database "database/db"
	repository_charges "database/repository/charges"
	"fmt"
	"log"
	"shared/common/dto"
)

func ChargesList(input dto.UsersListInput) ([]*dto.Charges, error) {
	ctx := context.Background()

	tx, err := database.DB.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %v", err)
	}
	defer tx.Rollback(ctx)

	q := repository_charges.New(tx)

	pramas := ConvertChargesListInput(input)
	total, err := q.ChargesList(context.Background(), pramas)
	log.Printf("Charges list: %v", total)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	var result []*dto.Charges
	for _, row := range total {
		result = append(result, ConvertCharges(row))
	}
	return result, nil
}
