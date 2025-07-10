package core

import (
	"context"
	repository_users "database/repository/users"
)

func CreateAuthUser(ctx context.Context, q *repository_users.Queries, username, passwordHash, role string) (*repository_users.CreateAuthUserRow, error) {
	user, err := q.CreateAuthUser(ctx, repository_users.CreateAuthUserParams{
		Username:     username,
		PasswordHash: passwordHash,
		Role:         role,
	})
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetAuthUserByUsername(ctx context.Context, q *repository_users.Queries, username string) (*repository_users.AuthUser, error) {
	user, err := q.GetAuthUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetAuthUserByID(ctx context.Context, q *repository_users.Queries, id int32) (*repository_users.GetAuthUserByIDRow, error) {
	user, err := q.GetAuthUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func ListAuthUsers(ctx context.Context, q *repository_users.Queries) ([]repository_users.ListAuthUsersRow, error) {
	return q.ListAuthUsers(ctx)
}
