package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func lookupEnv(dest *string, envVar string) error {
	if v, exists := os.LookupEnv(envVar); !exists {
		return fmt.Errorf("%v not found in local env", envVar)
	} else {
		*dest = v
		return nil
	}
}

func main() {
	var bd_pass, bd_user string
	var err error

	if err = godotenv.Load("secret.env"); err != nil {
		fmt.Printf("%v", err)
		return
	}

	if err := lookupEnv(&bd_pass, "BD_PASSWORD"); err != nil {
		fmt.Printf("%v", err)
		return
	}
	if err := lookupEnv(&bd_user, "BD_USER"); err != nil {
		fmt.Printf("%v", err)
		return
	}

	Dbx, err := pgxpool.New(context.Background(), fmt.Sprintf("postgres://%v:%v@nekiiinkognito.ru:5432/enshi_db", bd_user, bd_pass))
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	print(Dbx)

	fmt.Printf("Hey!, %v", "you")
}
