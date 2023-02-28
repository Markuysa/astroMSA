package database

import (
	"authService/app/model"
	"context"
	"database/sql"
	"time"
)

type UsersDB struct {
	db *sql.DB
}

func Add(ctx context.Context, user model.User) error {

	return nil
}

func Get(ctx context.Context, id int64) (model.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()
	return model.User{}, nil
}
