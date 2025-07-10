package core

import (
	"context"
	repository_auth_users "database/repository/auth_users"
)

func CreateAuthUser(ctx context.Context, q *repository_auth_users.Queries, username, passwordHash, role string) (*repository_auth_users.CreateAuthUserRow, error) {
	user, err := q.CreateAuthUser(ctx, repository_auth_users.CreateAuthUserParams{
		Username:     username,
		PasswordHash: passwordHash,
		Role:         role,
	})
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetAuthUserByUsername(ctx context.Context, q *repository_auth_users.Queries, username string) (*repository_auth_users.AuthUser, error) {
	user, err := q.GetAuthUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetAuthUserByID(ctx context.Context, q *repository_auth_users.Queries, id int32) (*repository_auth_users.GetAuthUserByIDRow, error) {
	user, err := q.GetAuthUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func ListAuthUsers(ctx context.Context, q *repository_auth_users.Queries) ([]repository_auth_users.ListAuthUsersRow, error) {
	return q.ListAuthUsers(ctx)
}
