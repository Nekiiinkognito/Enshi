package utils

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Pgx connection to database
var Dbx *pgxpool.Pool

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
	url := fmt.Sprintf("postgres://%v:%v@nekiiinkognito.ru:5432/enshi_db", bd_user, bd_pass)

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
