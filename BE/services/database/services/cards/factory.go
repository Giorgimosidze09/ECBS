package core

import (
	"shared/common/dto"

	repository_cards "database/repository/cards"

	"github.com/jackc/pgx/v5/pgtype"
)

func CreateCardParams(input dto.AssignCardInput) repository_cards.CreateCardParams {
	return repository_cards.CreateCardParams{
		CardID:   input.CardID,
		UserID:   int32(input.UserID),
		DeviceID: int32(input.DeviceID),
		Active:   pgtype.Bool{Bool: true, Valid: true},
		Type:     input.Type,
	}
}

func ConvertCardOutput(input repository_cards.CreateCardRow) *dto.CardOutput {
	return &dto.CardOutput{
		ID:       int(input.ID),
		UserID:   int(input.UserID),
		CardID:   input.CardID,
		DeviceID: int(input.DeviceID),
		Active:   input.Active.Bool,
	}
}

func ConvertCardsListInput(input dto.UsersListInput) repository_cards.CardsListParams {
	return repository_cards.CardsListParams{
		Limit:  int32(input.Limit),
		Offset: int32(input.Offset),
	}
}

func ConvertCardsListOutput(input repository_cards.CardsListRow) *dto.CardOutput {
	return &dto.CardOutput{
		ID:         int(input.ID),
		UserID:     int(input.UserID),
		CardID:     input.CardID,
		Active:     input.Active.Bool,
		AssignedAt: input.AssignedAt.Time.Format("2006-01-02 15:04:05"),
		Total:      int(input.Total),
	}
}

func CreateCardActivationParams(input dto.CardActivation) repository_cards.CreateCardActivationParams {
	return repository_cards.CreateCardActivationParams{
		CardID:          int32(input.CardID),
		ActivationStart: pgtype.Date{Time: input.ActivationStart, Valid: true},
		ActivationEnd:   pgtype.Date{Time: input.ActivationEnd, Valid: true},
	}
}

func ConvertCardActivationOutput(input repository_cards.CardActivation) *dto.CardActivation {
	return &dto.CardActivation{
		ID:              int(input.ID),
		CardID:          int(input.CardID),
		ActivationStart: input.ActivationStart.Time,
		ActivationEnd:   input.ActivationEnd.Time,
	}
}
