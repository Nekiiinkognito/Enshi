package utils

import (
	"context"
	db_repo "enshi/db/go_queries"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Pgx connection to database
var Dbx *pgxpool.Pool
var Dbx_connection *pgx.Conn
var Sqlc_db = db_repo.New(Dbx)

func SetupDatabase() error {

	var bd_pass, bd_user string
	var err error

	if err := LookupEnv(&bd_pass, "BD_PASSWORD"); err != nil {
		fmt.Printf("%v", err)
		return err
	}
	if err := LookupEnv(&bd_user, "BD_USER"); err != nil {
		fmt.Printf("%v", err)
		return err
	}

	// Url to connect
	url := fmt.Sprintf("postgres://%v:%v@nekiiinkognito.ru:5432/postgres", bd_user, bd_pass)

	// Connecting to database
	Dbx, err = pgxpool.New(context.Background(), url)
	Dbx_connection, err = pgx.Connect(context.Background(), url)

	if err != nil {
		fmt.Println("Unable to connect")
		fmt.Println(err)
	} else {
		fmt.Println("Connected successfully!")
	}

	return err
}
