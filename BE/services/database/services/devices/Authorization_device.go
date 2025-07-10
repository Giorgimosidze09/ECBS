package core

import (
	"context"
	database "database/db"
	repository_devices "database/repository/devices"
	"fmt"
	"shared/common/dto"
)

func Authorization(input dto.AuthorizedInput) ([]dto.AuthorizedAccessSyncDTO, error) {
	ctx := context.Background()

	tx, err := database.DB.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %v", err)
	}
	defer tx.Rollback(ctx)

	q := repository_devices.New(tx)

	rows, err := q.GetAuthorizedAccessByDeviceUniqueID(ctx, input.DeviceID)
	if err != nil {
		return nil, fmt.Errorf("failed to query authorized access: %v", err)
	}

	var result []dto.AuthorizedAccessSyncDTO

	for _, row := range rows {
		r := dto.AuthorizedAccessSyncDTO{
			CardID:   row.CardID,
			UserID:   int(row.UserID),
			UserName: row.UserName,
			Type:     row.Type,
			Active:   row.Active.Bool,
		}

		// Add PinCode if present
		if row.PinCode.Valid {
			pin := row.PinCode.String
			r.PinCode = &pin
		}

		if row.Balance.Valid {
			f, _ := row.Balance.Float64Value()
			r.Balance = &f.Float64
		}
		if row.RideCost.Valid {
			f, _ := row.RideCost.Float64Value()
			r.RideCost = &f.Float64
		}
		if row.ActivationStart.Valid {
			s := row.ActivationStart.Time.Format("2006-01-02")
			r.ActivationStart = &s
		}
		if row.ActivationEnd.Valid {
			s := row.ActivationEnd.Time.Format("2006-01-02")
			r.ActivationEnd = &s
		}

		result = append(result, r)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	return result, nil
}
