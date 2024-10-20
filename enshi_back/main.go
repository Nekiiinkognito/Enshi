package main

import (
	"context"
	db_repo "enshi/db/go_queries"
	"enshi/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	if err := utils.LoadEnv("utils/secret.env"); err != nil {
		fmt.Println(err.Error())
		return
	}

	if err := utils.SetupDatabase(); err != nil {
		fmt.Println(err.Error())
		return
	}
	defer utils.Dbx.Close()
	defer utils.Dbx_connection.Close(context.Background())

	router := gin.Default()
	if err := utils.SetupRotes(router); err != nil {
		fmt.Println(err.Error())
		return
	}

	// Transaction
	tx, _ := utils.Dbx_connection.Begin(context.Background())
	defer tx.Rollback(context.Background())

	repo := db_repo.New(tx)

	users, _ := repo.GetAllUsers(context.Background())

	for _, user := range users {
		fmt.Printf("%v\n", user.Username)
	}

	if err := tx.Commit(context.Background()); err != nil {
		return
	}

	fmt.Printf("Hey!, %v", "you")
}
