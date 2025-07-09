package core

import (
	"context"
	database "database/db"
	repository_user "database/repository/users"
	"fmt"
	"log"
	"shared/common/dto"
)

func CreateUser(input dto.CreateUsersInput) (*dto.UserOutput, error) {
	ctx := context.Background()

	tx, err := database.DB.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %v", err)
	}
	defer tx.Rollback(ctx)

	q := repository_user.New(tx)

	dbParams := CreateUserParams(input)

	user, err := q.CreateUser(context.Background(), dbParams)

	log.Printf("Created user: %v", user)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return ConvertOutput(user), nil
}
