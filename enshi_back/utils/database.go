package utils

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Pgx connection to database
var Dbx *pgxpool.Pool

func SetupDatabase() error {

	// Url to connect
	url := "postgres://root:JET159sam753@nekiiinkognito.ru:5432/recipes"

	var err error

	// Connecting to database
	Dbx, err = pgxpool.New(context.Background(), url)

	if err != nil {
		fmt.Println("Unable to connect")
		fmt.Println(err)
	} else {
		fmt.Println("Connected successfully!")
	}

	return err
}
