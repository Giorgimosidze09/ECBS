package core

import (
	"context"
	database "database/db"
	repository_user "database/repository/users"
	"fmt"
	"shared/common/dto"

	"github.com/jackc/pgx/v5/pgtype"
)

func SyncAccessLogs(input dto.SyncAccessLogInput) error {
	ctx := context.Background()

	tx, err := database.DB.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin tx: %v", err)
	}
	defer tx.Rollback(ctx)

	q := repository_user.New(tx)

	for _, log := range input.Logs {
		if !log.Success {
			continue // skip failed attempts
		}

		// Convert log.CardID (string) to int32 for GetCardByCardID

		card, err := q.GetCardByItsCardID(ctx, log.CardID)
		if err != nil || !card.Active.Bool {
			continue
		}

		if card.Type == "balance" {
			// Deduct ride cost and insert charge
			balanceRow, err := q.GetBalanceByUserID(ctx, pgtype.Int4{Int32: card.UserID, Valid: true})
			if err != nil || !balanceRow.RideCost.Valid {
				continue
			}

			rideCost, _ := balanceRow.RideCost.Float64Value()

			var rideCostNumeric pgtype.Numeric
			if err := rideCostNumeric.Scan(fmt.Sprintf("%.2f", rideCost.Float64)); err != nil {
				continue
			}

			// Deduct
			err = q.DeductBalance(ctx, repository_user.DeductBalanceParams{
				UserID:  pgtype.Int4{Int32: card.UserID, Valid: true},
				Balance: rideCostNumeric,
			})
			if err != nil {
				continue
			}

			// Insert charge
			_ = q.InsertCharge(ctx, repository_user.InsertChargeParams{
				UserID:      pgtype.Int4{Int32: card.UserID, Valid: true},
				Amount:      rideCostNumeric,
				Description: pgtype.Text{String: "Offline ride sync", Valid: true},
			})

			// Optional: deactivate card if new balance == 0
			newBalRow, _ := q.GetBalanceByUserID(ctx, pgtype.Int4{Int32: card.UserID, Valid: true})
			if newBalRow.Balance.Valid {
				balF, _ := newBalRow.Balance.Float64Value()
				if balF.Float64 == 0 {
					tx.Exec(ctx, "UPDATE cards SET active = false WHERE id = $1", card.ID)
				}
			}
		} else if card.Type == "activation" {
			// For now, do nothing — access already allowed
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit tx: %v", err)
	}

	return nil
}
