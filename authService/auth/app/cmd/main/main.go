package main

import (
	db "authService/app/internal/database"
	"authService/app/model"
	"context"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	ctx := context.Background()
	db.Add(ctx, model.User{})
	database, err := gorm.Open(postgres.Open("host=localhost port=5432 user=postgres password=islam20011"))
	if err != nil {
		log.Fatal(err)
	}
	err = database.AutoMigrate(&model.User{})
	if err != nil {
		return
	}

}
