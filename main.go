package main

import (
	"context"
	"database/sql"
	"log"
	
	_ "github.com/lib/pq"

	"github.com/kovalyov-valentin/sqlc-example/postgres"
)

func main() {
	connection, err := sql.Open("postgres", "user=postgres password=newPassword dbname=sqlc sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	db := postgres.New(connection)

	user, err := db.CreateUser(context.Background(), postgres.CreateUserParams{
		Firstname: "Valentin",
		Lastname: "Kovalyov",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(user)
}