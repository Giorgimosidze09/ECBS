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

func CardsList(input dto.UsersListInput) ([]*dto.CardOutput, error) {
	ctx := context.Background()

	tx, err := database.DB.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %v", err)
	}
	defer tx.Rollback(ctx)

	q := repository_user.New(tx)

	pramas := ConvertCardsListInput(input)
	total, err := q.CardsList(context.Background(), pramas)
	log.Printf("Cards list: %v", total)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	var result []*dto.CardOutput
	for _, row := range total {
		result = append(result, &dto.CardOutput{
			ID:         int(row.ID),
			UserID:     int(row.UserID),
			CardID:     row.CardID,
			DeviceID:   int(row.DeviceID),
			Type:       row.Type,
			Active:     row.Active.Bool,
			AssignedAt: row.AssignedAt.Time.Format("2006-01-02 15:04:05"),
			Total:      int(row.Total),
			// Type and DeviceID not available in CardsListRow
		})
	}
	return result, nil
}

func UpdateCard(input dto.CardOutput) error {
	ctx := context.Background()
	tx, err := database.DB.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %v", err)
	}
	defer tx.Rollback(ctx)

	q := repository_user.New(tx)
	params := repository_user.UpdateCardParams{
		ID:       int32(input.ID),
		CardID:   input.CardID,
		UserID:   int32(input.UserID),
		DeviceID: int32(input.DeviceID),
		Type:     input.Type,
		Active:   pgtype.Bool{Bool: input.Active, Valid: true},
	}
	if err := q.UpdateCard(ctx, params); err != nil {
		return err
	}
	return tx.Commit(ctx)
}

func SoftDeleteCard(cardID int32) error {
	ctx := context.Background()
	tx, err := database.DB.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %v", err)
	}
	defer tx.Rollback(ctx)

	q := repository_user.New(tx)
	if err := q.SoftDeleteCard(ctx, cardID); err != nil {
		return err
	}
	return tx.Commit(ctx)
}

func GetCardByID(cardID int32) (*dto.CardOutput, error) {
	ctx := context.Background()
	q := repository_user.New(database.DB)
	row, err := q.GetCardByID(ctx, cardID)
	if err != nil {
		return nil, err
	}
	return &dto.CardOutput{
		ID:         int(row.ID),
		UserID:     int(row.UserID),
		CardID:     row.CardID,
		DeviceID:   int(row.DeviceID),
		Active:     row.Active.Bool,
		AssignedAt: row.AssignedAt.Time.Format("2006-01-02 15:04:05"),
		Type:       row.Type,
		Deleted:    !row.Active.Bool,
	}, nil
}
