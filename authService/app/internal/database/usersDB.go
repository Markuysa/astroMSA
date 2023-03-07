package database

import (
	"astroService/app/pkg/workers/astroWorker"
	"authService/app/internal/model"
	"context"
	"errors"
	"github.com/jmoiron/sqlx"
	"log"
)

var (
	AddingUserError  = errors.New("Error with addimg the user")
	GettingUserError = errors.New("Error with getting the user")
)

type UsersDB struct {
	db *sqlx.DB
}

func New(ctx context.Context) *UsersDB {
	datab, err := sqlx.ConnectContext(ctx,
		"postgres",
		"host=localhost port=5432 user=postgres password=islam20011 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return &UsersDB{db: datab}
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
		return AddingUserError
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
		return nil, GettingUserError
	}
	return &user, nil
}
