package core

import (
	"shared/common/dto"

	repository_devices "database/repository/devices"

	"github.com/jackc/pgx/v5/pgtype"
)

func CreateDeviceParams(input dto.DevicesInput) repository_devices.CreateDeviceParams {
	return repository_devices.CreateDeviceParams{
		DeviceID: input.DeviceID,
		Location: pgtype.Text{String: input.Location, Valid: true},
	}
}

func ConvertDeviceOutput(input repository_devices.Device) *dto.DevicesOutput {
	return &dto.DevicesOutput{
		ID:          int(input.ID),
		DeviceID:    input.DeviceID,
		Location:    input.Location.String,
		InstalledAt: input.InstalledAt.Time.Format("2006-01-02 15:04:05"),
		Active:      input.Active.Bool,
	}
}

func ConvertDevicesListInput(input dto.UsersListInput) repository_devices.DeviceListParams {
	return repository_devices.DeviceListParams{
		Limit:  int32(input.Limit),
		Offset: int32(input.Offset),
	}
}

func ConvertDevicesListOutput(input repository_devices.DeviceListRow) *dto.DevicesOutput {
	return &dto.DevicesOutput{
		ID:          int(input.ID),
		DeviceID:    input.DeviceID,
		Location:    input.Location.String,
		InstalledAt: input.InstalledAt.Time.Format("2006-01-02 15:04:05"),
		Active:      input.Active.Bool,
	}
}
