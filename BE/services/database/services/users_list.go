package core

import (
	"context"
	database "database/db"
	repository_user "database/repository/users"
	"fmt"
	"log"
	"shared/common/dto"

	"github.com/jackc/pgx/v5/pgtype"
)

func UsersList(input dto.UsersListInput) ([]*dto.UsersListOutput, error) {
	ctx := context.Background()

	tx, err := database.DB.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %v", err)
	}
	defer tx.Rollback(ctx)

	q := repository_user.New(tx)

	pramas := ConvertUserListInput(input)
	total, err := q.UsersList(context.Background(), pramas)
	log.Printf("Users List: %v", total)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	var result []*dto.UsersListOutput
	for _, row := range total {
		result = append(result, ConvertUsersListOutput(row))
	}
	return result, nil
}

func UpdateUser(input dto.UserOutput) error {
	ctx := context.Background()
	tx, err := database.DB.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %v", err)
	}
	defer tx.Rollback(ctx)

	q := repository_user.New(tx)
	params := repository_user.UpdateUserParams{
		ID:    input.ID,
		Name:  input.Name,
		Email: pgtype.Text{String: input.Email, Valid: true},
		Phone: pgtype.Text{String: input.Phone, Valid: true},
	}
	if err := q.UpdateUser(ctx, params); err != nil {
		return err
	}
	return tx.Commit(ctx)
}

func SoftDeleteUser(userID int32) error {
	ctx := context.Background()
	tx, err := database.DB.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %v", err)
	}
	defer tx.Rollback(ctx)

	q := repository_user.New(tx)
	if err := q.SoftDeleteUser(ctx, userID); err != nil {
		return err
	}
	return tx.Commit(ctx)
}

func GetUserByID(userID int32) (*dto.UserOutput, error) {
	ctx := context.Background()
	q := repository_user.New(database.DB)
	row, err := q.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &dto.UserOutput{
		ID:      row.ID,
		Name:    row.Name,
		Email:   row.Email.String,
		Phone:   row.Phone.String,
		Deleted: row.Deleted.Bool,
	}, nil
}
