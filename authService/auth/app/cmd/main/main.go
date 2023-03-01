package main

import (
	db "authService/app/internal/database"
	"authService/app/internal/model"
	"context"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	datab, err := sqlx.ConnectContext(ctx,
		"postgres",
		"host=localhost port=5432 user=postgres password=islam20011 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	usersDatabase := db.New(datab)
	usersDatabase.Add(ctx, model.User{
		Name:      "artur",
		Email:     "aa@gmail.ru",
		Sign:      "taurus",
		BirthInfo: time.Now(),
	})
	get, err := usersDatabase.Get(ctx, 1)
	if err != nil {
		return
	}
	log.Println(get)
	//database, err := gorm.Open(postgres.Open("host=localhost port=5432 user=postgres password=islam20011"))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//err = database.AutoMigrate(&model.User{})
	//if err != nil {
	//	return
	//}

}
