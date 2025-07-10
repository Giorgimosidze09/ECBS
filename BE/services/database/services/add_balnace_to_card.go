package core

import (
	"context"
	database "database/db"
	repository_user "database/repository/users"
	"fmt"
	"log"
	"shared/common/dto"
	"time"
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
	total, err := q.AddBalanceToCard(ctx, pramas)
	log.Printf("balance list: %v", total)
	if err != nil {
		return err
	}

	// Fetch card info
	cardRow, err := q.GetCardByID(ctx, int32(input.CardID))
	if err != nil {
		return fmt.Errorf("failed to fetch card: %v", err)
	}

	// Subscription activation logic
	if input.Amount >= 15 {
		// Activate subscription for 30 days
		activationStart := time.Now()
		activationEnd := activationStart.AddDate(0, 0, 30)
		activation := dto.CardActivation{
			CardID:          int(cardRow.ID),
			ActivationStart: activationStart,
			ActivationEnd:   activationEnd,
		}
		_, err := AddCardActivation(activation)
		if err != nil {
			return fmt.Errorf("failed to activate subscription: %v", err)
		}
		// Change card type to 'activation'
		updateParams := repository_user.UpdateCardParams{
			ID:       cardRow.ID,
			CardID:   cardRow.CardID,
			UserID:   cardRow.UserID,
			DeviceID: cardRow.DeviceID,
			Type:     "activation",
			Active:   cardRow.Active,
		}
		if err := q.UpdateCard(ctx, updateParams); err != nil {
			return fmt.Errorf("failed to update card type to activation: %v", err)
		}
	} else if input.Amount < 15 {
		// Change card type to 'balance'
		updateParams := repository_user.UpdateCardParams{
			ID:       cardRow.ID,
			CardID:   cardRow.CardID,
			UserID:   cardRow.UserID,
			DeviceID: cardRow.DeviceID,
			Type:     "balance",
			Active:   cardRow.Active,
		}
		if err := q.UpdateCard(ctx, updateParams); err != nil {
			return fmt.Errorf("failed to update card type to balance: %v", err)
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}
