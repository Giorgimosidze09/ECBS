package database_client

import (
	"shared/common/dto"
	"shared/common/utils"
)

func (c DatabaseClient) SyncAuthorizedAccess(input dto.AuthorizedInput) ([]dto.AuthorizedAccessSyncDTO, error) {
	data, err := utils.Encode(input)
	if err != nil {
		return nil, err
	}

	response := c.Publisher.Request("authorized", data)
	if response.Error != nil {
		return nil, response.Error
	}

	return utils.Decode[[]dto.AuthorizedAccessSyncDTO](response.Data)
}

func (c DatabaseClient) SyncAccessLogs(input dto.SyncAccessLogInput) error {
	data, err := utils.Encode(input)
	if err != nil {
		return err
	}

	response := c.Publisher.Request("access_logs", data)
	if response.Error != nil {
		return response.Error
	}

	return nil
}
