package database

import (
	"context"
	"errors"
	"github.com/Markuysa/astroMSA/astroService/app/pkg/workers/astroWorker"
	"github.com/Markuysa/astroMSA/authService/app/internal/helpers/hash"
	"github.com/Markuysa/astroMSA/authService/app/internal/helpers/protobuf"
	"github.com/Markuysa/astroMSA/authService/app/internal/model"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

var (
	AddingUserErr        = errors.New("error with adding the user")
	GettingUserErr       = errors.New("error with getting the user")
	UserAlreadyExistsErr = errors.New("user already exists")
	PasswordHashErr      = errors.New("password hash err")
)

// UsersRepository - Interface of users database
type UsersRepository interface {
	Add(ctx context.Context, user *model.User) error
	Get(ctx context.Context, id int64) (*model.User, error)
	GetUsersEmailsWithAllowedNotifications(ctx context.Context) ([]string, error)
}

// UsersDB - structure, that implements a UsersRepository interface
type UsersDB struct {
	db *sqlx.DB
}

// New method creates an object of users repository
func New(ctx context.Context) *UsersDB {

	datab, err := sqlx.ConnectContext(ctx,
		"postgres",
		"host=postgres user=postgres port=5432 password=islam20011 sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}
	return &UsersDB{db: datab}
}

// Add method creates user object and saves him/her in the database
func (db *UsersDB) Add(ctx context.Context, user *model.User) error {
	if _, err := db.Get(ctx, user.Email); err == nil {
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
		notifications
	)values (
	    $1,$2,$3,$4,$5,$6
	)
	`
	user.Sign = astroWorker.CalculateSign(user.BirthInfo)
	_, err = db.db.ExecContext(ctx, query,
		user.Email,
		protobuf.DateToTime(user.BirthInfo),
		user.Sign,
		user.Name,
		password,
		user.Notifications,
	)
	if err != nil {
		return AddingUserErr
	}
	return nil
}

// Get method selects a user from the database with
// some id
func (db *UsersDB) Get(ctx context.Context, eMail string) (*model.User, error) {

	query := `
		select email,
			   birth_info,
			   sign,
			   name,
			   password,
			   notifications
		from users
		where email=$1
	`
	var user model.User
	row := db.db.QueryRowxContext(ctx, query, eMail)
	var birthDate time.Time
	err := row.Scan(&user.Email, &birthDate, &user.Sign, &user.Name, &user.Password, &user.Notifications)
	user.BirthInfo = protobuf.TimeToInternalDate(birthDate)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUsersEmailsWithAllowedNotifications selects all the users with allowed notifications
// to send the daily predictions through gmail
func (db *UsersDB) GetUsersEmailsWithAllowedNotifications(ctx context.Context) ([]string, error) {

	query := `
		select email
		from users
		where notifications=true
	`
	var users []string
	var userEmail string
	rows, err := db.db.QueryContext(ctx, query)
	for rows.Next() {
		if err := rows.Scan(&userEmail); err != nil {
			return nil, err
		}
		users = append(users, userEmail)
	}
	if err != nil {
		return nil, err
	}

	return users, nil
}
