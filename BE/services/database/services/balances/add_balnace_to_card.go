package core

import (
	"context"
	database "database/db"
	repository_balances "database/repository/balances"
	repository_cards "database/repository/cards"
	repository_paybox "database/repository/paybox"
	cardsActivation "database/services/cards"
	"fmt"
	"log"
	"shared/common/dto"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func AddBalanceToCard(input dto.PayboxTopupRequest) error {
	ctx := context.Background()

	tx, err := database.DB.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %v", err)
	}
	defer tx.Rollback(ctx)

	q := repository_paybox.New(tx)
	balances := repository_balances.New(tx)
	cards := repository_cards.New(tx)

	// 1. Check for duplicate transaction
	exists, err := q.CheckPayboxTransactionExists(ctx, pgtype.Text{String: input.TransactionID, Valid: true})
	if err != nil {
		return fmt.Errorf("failed to check for duplicate transaction: %v", err)
	}
	if exists {
		log.Printf("transaction already processed: %s", input.TransactionID)
		return nil
	}

	// 2. Add balance to card
	params := AddBalanceToCardParams(input)
	total, err := balances.AddBalanceToCard(ctx, params)
	if err != nil {
		return fmt.Errorf("failed to add balance: %v", err)
	}
	log.Printf("balance updated: card_id=%d, new_total=%v", input.CardID, total)

	// 3. Fetch card info
	cardRow, err := cards.GetCardByID(ctx, int32(input.CardID))
	if err != nil {
		return fmt.Errorf("failed to fetch card: %v", err)
	}

	// 4. Handle subscription logic
	if input.Amount >= 15 {
		activationStart := time.Now()
		activationEnd := activationStart.AddDate(0, 0, 30)
		activation := dto.CardActivation{
			CardID:          int(cardRow.ID),
			ActivationStart: activationStart,
			ActivationEnd:   activationEnd,
		}
		_, err := cardsActivation.AddCardActivation(activation)
		if err != nil {
			return fmt.Errorf("failed to activate subscription: %v", err)
		}
		// Change card type to 'activation'
		updateParams := repository_cards.UpdateCardParams{
			ID:       cardRow.ID,
			CardID:   cardRow.CardID,
			UserID:   cardRow.UserID,
			DeviceID: cardRow.DeviceID,
			Type:     "activation",
			Active:   cardRow.Active,
		}
		if err := cards.UpdateCard(ctx, updateParams); err != nil {
			return fmt.Errorf("failed to update card to activation: %v", err)
		}
	} else {
		// Change card type to 'balance'
		updateParams := repository_cards.UpdateCardParams{
			ID:       cardRow.ID,
			CardID:   cardRow.CardID,
			UserID:   cardRow.UserID,
			DeviceID: cardRow.DeviceID,
			Type:     "balance",
			Active:   cardRow.Active,
		}
		if err := cards.UpdateCard(ctx, updateParams); err != nil {
			return fmt.Errorf("failed to update card to balance: %v", err)
		}
	}

	// 5. Log paybox transaction
	logParams := repository_paybox.CreatePayboxTransactionParams{
		TransactionID: pgtype.Text{String: input.TransactionID, Valid: true},
		CardID:        int32(input.CardID),
		Amount:        input.Amount,
		Source:        pgtype.Text{String: "paybox", Valid: true},
		CreatedAt:     pgtype.Timestamp{Time: time.Now(), Valid: true},
	}
	if err := q.CreatePayboxTransaction(ctx, logParams); err != nil {
		return fmt.Errorf("failed to log paybox transaction: %v", err)
	}

	// 6. Commit
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	log.Printf("paybox top-up successful: card_id=%d, amount=%.2f", input.CardID, input.Amount)
	return nil
}
