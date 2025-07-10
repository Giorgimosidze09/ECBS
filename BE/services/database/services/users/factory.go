package core

import (
	"shared/common/dto"

	repository_user "database/repository/users"

	"github.com/jackc/pgx/v5/pgtype"
)

func CreateUserParams(input dto.CreateUsersInput) repository_user.CreateUserParams {
	return repository_user.CreateUserParams{
		Name:  input.Name,
		Email: pgtype.Text{String: input.Email, Valid: true},
		Phone: pgtype.Text{String: input.Phone, Valid: true},
	}
}

func ConvertOutput(input repository_user.CreateUserRow) *dto.UserOutput {
	return &dto.UserOutput{
		ID:    input.ID,
		Name:  input.Name,
		Email: input.Email.String,
		Phone: input.Phone.String,
	}
}

func ConvertUserListInput(input dto.UsersListInput) repository_user.UsersListParams {
	return repository_user.UsersListParams{
		Limit:  int32(input.Limit),
		Offset: int32(input.Offset),
	}
}

func ConvertUsersListOutput(input repository_user.UsersListRow) *dto.UsersListOutput {

	return &dto.UsersListOutput{
		ID:           input.ID,
		Name:         input.Name,
		Email:        input.Email.String,
		Phone:        input.Phone.String,
		CardCount:    input.CardCount,
		TotalBalance: float64(input.TotalBalance),
		Total:        int(input.Total),
	}
}
