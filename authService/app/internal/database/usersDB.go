package database

import (
	"authService/app/internal/helpers/hash"
	"authService/app/internal/model"
	"context"
	"errors"
	"github.com/Markuysa/astroMSA/astroService/app/pkg/workers/astroWorker"
	"github.com/jmoiron/sqlx"
	"log"
)

var (
	AddingUserErr        = errors.New("error with adding the user")
	GettingUserErr       = errors.New("error with getting the user")
	UserAlreadyExistsErr = errors.New("user already exists")
	PasswordHashErr      = errors.New("password hash err")
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
	if _, err := db.Get(ctx, int64(user.ID)); err == nil {
		return UserAlreadyExistsErr
	}
	password, err := hash.HashPassword(user.Password)
	if err != nil {
		return PasswordHashErr
	}
	query := `
	insert into users(
	    email,
	    birth_info,
	    sign,
		name,
		password,
	    created_at
	)values (
	    $1,$2,$3,$4,$5,$6 
	)
	`
	user.Sign = astroWorker.CalculateSign(user.BirthInfo)
	_, err = db.db.ExecContext(ctx, query,
		user.Email,
		user.BirthInfo,
		user.Sign,
		user.Name,
		password,
		user.CreatedAt,
	)
	if err != nil {
		return AddingUserErr
	}
	return nil
}

func (db *UsersDB) Get(ctx context.Context, id int64) (*model.User, error) {

	query := `
		select email,
			   birth_info,
			   sign,
			   name,
			   password,
			   created_at
		from users
		where id=$1
	`

	var user model.User
	row := db.db.QueryRowxContext(ctx, query, id)
	err := row.Scan(&user.Email, &user.BirthInfo, &user.Sign, &user.Name, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, GettingUserErr
	}
	return &user, nil
}

func (db *UsersDB) GetUsersEmailsWithAllowedNotifications(ctx context.Context) ([]string, error) {

	query := `
		select email
		from users
		where notifications=true
	`
	var users []string

	rows, err := db.db.QueryContext(ctx, query)
	log.Println(rows)
	if err != nil {
		return nil, err
	}
	err = rows.Scan(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
