package core

import (
	"context"
	database "database/db"
	repository_devices "database/repository/devices"
	"fmt"
	"log"
	"shared/common/dto"

	"github.com/jackc/pgx/v5/pgtype"
)

func DevicesList(input dto.UsersListInput) ([]*dto.DevicesOutput, error) {
	ctx := context.Background()

	tx, err := database.DB.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %v", err)
	}
	defer tx.Rollback(ctx)

	q := repository_devices.New(tx)

	pramas := ConvertDevicesListInput(input)
	total, err := q.DeviceList(ctx, pramas)
	log.Printf("Devices list: %v", total)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	var result []*dto.DevicesOutput
	for _, row := range total {
		result = append(result, ConvertDevicesListOutput(row))
	}
	return result, nil
}

func UpdateDevice(input dto.DeviceOutput) error {
	ctx := context.Background()
	tx, err := database.DB.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %v", err)
	}
	defer tx.Rollback(ctx)

	q := repository_devices.New(tx)
	params := repository_devices.UpdateDeviceParams{
		ID:       int32(input.ID),
		DeviceID: input.DeviceID,
		Location: pgtype.Text{String: input.Location, Valid: true},
		Active:   pgtype.Bool{Bool: input.Active, Valid: true},
	}
	if err := q.UpdateDevice(ctx, params); err != nil {
		return err
	}
	return tx.Commit(ctx)
}

func SoftDeleteDevice(deviceID int32) error {
	ctx := context.Background()
	tx, err := database.DB.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %v", err)
	}
	defer tx.Rollback(ctx)

	q := repository_devices.New(tx)
	if err := q.SoftDeleteDevice(ctx, deviceID); err != nil {
		return err
	}
	return tx.Commit(ctx)
}

func GetDeviceByID(deviceID int32) (*dto.DeviceOutput, error) {
	ctx := context.Background()
	q := repository_devices.New(database.DB)
	row, err := q.GetDeviceByID(ctx, deviceID)
	if err != nil {
		return nil, err
	}
	return &dto.DeviceOutput{
		ID:        int(row.ID),
		DeviceID:  row.DeviceID,
		Location:  row.Location.String,
		Active:    row.Active.Bool,
		Installed: row.InstalledAt.Time.Format("2006-01-02 15:04:05"),
	}, nil
}
