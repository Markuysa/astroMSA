package database

import (
	"authService/app/internal/model"
	"authService/app/pkg/workers/astroWorker"
	"context"
	"errors"
	"github.com/jmoiron/sqlx"
)

var (
	addingUserError = errors.New("Error with addimg the user")
)

type UsersDB struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *UsersDB {
	return &UsersDB{db: db}
}

func (db *UsersDB) Add(ctx context.Context, user model.User) error {
	query := `
	insert into users(
	    email,
	    birth_info,
	    sign,
		name
	)values (
	    $1,$2,$3,$4  
	)
	`
	user.Sign = astroWorker.CalculateSign(user.BirthInfo)
	_, err := db.db.ExecContext(ctx, query,
		user.Email,
		user.BirthInfo,
		user.Sign,
		user.Name)
	if err != nil {
		return addingUserError
	}
	return nil
}

func (db *UsersDB) Get(ctx context.Context, id int64) (*model.User, error) {

	query := `
		select email,
			   birth_info,
			   sign,
			   name
		from users
		where id=$1
	`
	var user model.User
	row := db.db.QueryRowxContext(ctx, query, id)
	err := row.Scan(&user.Email, &user.BirthInfo, &user.Sign, &user.Name)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
