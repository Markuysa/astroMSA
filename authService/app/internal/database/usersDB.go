package database

import (
	"context"
	"errors"
	"github.com/Markuysa/astroMSA/astroService/app/pkg/workers/astroWorker"
	"github.com/Markuysa/astroMSA/authService/app/internal/helpers/hash"
	"github.com/Markuysa/astroMSA/authService/app/internal/helpers/protobuf"
	"github.com/Markuysa/astroMSA/authService/app/internal/model"
	messagesModel "github.com/Markuysa/astroMSA/messageSenderService/app/pkg/model"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
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
	Get(ctx context.Context, eMail string) (*model.User, error)
	GetUsersEmailsWithAllowedNotifications(ctx context.Context) ([]messagesModel.Receiver, error)
	AuthUser(ctx context.Context, email string, password string) (bool, error)
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
	db, err := gorm.Open(postgres.Open("host=postgres user=postgres port=5432 password=islam20011 sslmode=disable"), &gorm.Config{})
	if err != nil {
		return nil
	}
	// создание таблицы
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		return nil
	}
	if err != nil {
		log.Fatal(err)
	}
	return &UsersDB{db: datab}
}
func (db *UsersDB) AuthUser(ctx context.Context, email string, password string) (bool, error) {

	query := `
	select password from users
	where email=$1
`
	var passwordDB string
	err := db.db.QueryRowx(query, email).Scan(&passwordDB)
	if err != nil {
		return false, err
	}
	if ok := hash.CheckPasswordHash(password, passwordDB); ok {
		return true, nil
	}
	return false, errors.New("incorrect password or username")
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
	user.Sign = astroWorker.CalculateSign(protobuf.TimeToInternalDate(user.BirthInfo))
	_, err = db.db.ExecContext(ctx, query,
		user.Email,
		user.BirthInfo,
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
	err := row.Scan(&user.Email, &user.BirthInfo, &user.Sign, &user.Name, &user.Password, &user.Notifications)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUsersEmailsWithAllowedNotifications selects all the users with allowed notifications
// to send the daily predictions through gmail
func (db *UsersDB) GetUsersEmailsWithAllowedNotifications(ctx context.Context) ([]messagesModel.Receiver, error) {

	query := `
		select email, sign
		from users
		where notifications=true
	`
	var users []messagesModel.Receiver
	var userEmail, sign string
	rows, err := db.db.QueryContext(ctx, query)
	for rows.Next() {
		if err := rows.Scan(&userEmail, &sign); err != nil {
			return nil, err
		}
		users = append(users, messagesModel.Receiver{Email: userEmail, Zodiac: sign})
	}
	if err != nil {
		return nil, err
	}

	return users, nil
}
