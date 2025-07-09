package core

import (
	"context"
	database "database/db"
	repository_user "database/repository/users"
	"fmt"
	"log"
	"shared/common/dto"
)

func CreateDevices(input dto.DevicesInput) (*dto.DevicesOutput, error) {
	ctx := context.Background()

	tx, err := database.DB.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %v", err)
	}
	defer tx.Rollback(ctx)

	q := repository_user.New(tx)

	dbParams := CreateDeviceParams(input)

	devices, err := q.CreateDevice(context.Background(), dbParams)

	log.Printf("Created devices: %v", devices)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return ConvertDeviceOutput(devices), nil
}
