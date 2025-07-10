package core

import (
	"context"
	database "database/db"
	repository_user "database/repository/users"
	"fmt"
	"shared/common/dto"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func ValidateCard(input dto.ValidateCardInput) (*dto.ValidateCardOutput, error) {
	ctx := context.Background()

	tx, err := database.DB.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %v", err)
	}
	defer tx.Rollback(ctx)

	q := repository_user.New(tx)

	// Get card by card_id
	card, err := q.GetCardByCardID(ctx, int32(input.CardID))
	if err != nil {
		return &dto.ValidateCardOutput{
			Valid:   false,
			Message: "Card not found",
		}, nil
	}

	if !card.Active.Bool {
		return &dto.ValidateCardOutput{
			Valid:   false,
			Message: "Card inactive",
		}, nil
	}

	// Check card type logic
	switch card.Type {
	case "balance":
		// Existing balance logic
		// Get balance row for user
		balanceRow, err := q.GetBalanceByUserID(ctx, pgtype.Int4{Int32: card.UserID, Valid: true})
		if err != nil {
			return &dto.ValidateCardOutput{
				Valid:   false,
				Message: "Balance not found",
			}, nil
		}

		var balance float64
		if balanceRow.Balance.Valid {
			f, _ := balanceRow.Balance.Float64Value()
			balance = f.Float64
		}

		var rideCostValue float64
		if balanceRow.RideCost.Valid {
			f, _ := balanceRow.RideCost.Float64Value()
			rideCostValue = f.Float64
		}
		if balance < rideCostValue {
			return &dto.ValidateCardOutput{
				Valid:   false,
				Message: "Insufficient balance",
			}, nil
		}

		// Properly set numericRideCost from float64 rideCostValue
		var numericRideCost pgtype.Numeric
		if err := numericRideCost.Scan(fmt.Sprintf("%.2f", rideCostValue)); err != nil {
			return nil, fmt.Errorf("failed to set numericRideCost: %v", err)
		}

		// Deduct ride cost from balance
		err = q.DeductBalance(ctx, repository_user.DeductBalanceParams{
			UserID:  pgtype.Int4{Int32: card.UserID, Valid: true},
			Balance: numericRideCost,
		})
		if err != nil {
			return &dto.ValidateCardOutput{
				Valid:   false,
				Message: "Failed to deduct balance",
			}, nil
		}

		// Check if balance is now zero and deactivate card if so
		updatedBalanceRow, err := q.GetBalanceByUserID(ctx, pgtype.Int4{Int32: card.UserID, Valid: true})
		if err == nil {
			var newBalance float64
			if updatedBalanceRow.Balance.Valid {
				f, _ := updatedBalanceRow.Balance.Float64Value()
				newBalance = f.Float64
			}
			if newBalance == 0 {
				// Deactivate the card
				_, err := tx.Exec(ctx, "UPDATE cards SET active = false WHERE id = $1", card.ID)
				if err != nil {
					return &dto.ValidateCardOutput{
						Valid:   false,
						Message: "Failed to deactivate card",
					}, nil
				}
			}
		}

		// Insert charge record
		err = q.InsertCharge(ctx, repository_user.InsertChargeParams{
			UserID:      pgtype.Int4{Int32: card.UserID, Valid: true},
			Amount:      numericRideCost,
			Description: pgtype.Text{String: "Ride fare deducted", Valid: true},
		})
		if err != nil {
			return nil, fmt.Errorf("failed to insert charge: %v", err)
		}

		if err := tx.Commit(ctx); err != nil {
			return nil, fmt.Errorf("failed to commit transaction: %v", err)
		}

		return &dto.ValidateCardOutput{
			Valid:    true,
			UserID:   int(card.UserID),
			UserName: card.UserName,
			Balance:  balance - rideCostValue,
			Message:  "Charge applied, ride allowed",
		}, nil
	case "activation":
		// Activation card logic: check card_activations for valid period
		var isActive bool
		row := tx.QueryRow(ctx, "SELECT activation_start, activation_end FROM card_activations WHERE card_id = $1 ORDER BY activation_end DESC LIMIT 1", card.ID)
		var activationStart, activationEnd pgtype.Date
		err := row.Scan(&activationStart, &activationEnd)
		if err != nil {
			return &dto.ValidateCardOutput{
				Valid:   false,
				Message: "No activation period found",
			}, nil
		}
		// Use Go's time.Now() for current date
		now := time.Now()
		if activationStart.Valid && activationEnd.Valid {
			if !activationStart.Time.After(now) && !activationEnd.Time.Before(now) {
				isActive = true
			}
		}
		if !isActive {
			return &dto.ValidateCardOutput{
				Valid:   false,
				Message: "Card not active for this period",
			}, nil
		}
		if err := tx.Commit(ctx); err != nil {
			return nil, fmt.Errorf("failed to commit transaction: %v", err)
		}
		return &dto.ValidateCardOutput{
			Valid:    true,
			UserID:   int(card.UserID),
			UserName: card.UserName,
			Balance:  0, // Unlimited rides, no balance
			Message:  "Activation valid, ride allowed",
		}, nil
	}

	return nil, fmt.Errorf("unknown card type")
}
