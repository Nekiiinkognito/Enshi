package main

import (
	"context"
	db_repo "enshi/db/go_queries"
	"enshi/db_connection"
	"enshi/env"
	"enshi/global"
	"enshi/routes"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	if err := env.LoadEnv("utils/secret.env"); err != nil {
		fmt.Println(err.Error())
		return
	}

	if err := db_connection.SetupDatabase(); err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db_connection.Dbx.Close()
	defer db_connection.Dbx_connection.Close(context.Background())

	router := gin.Default()

	f, err := os.OpenFile("gin.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	gin.DefaultWriter = io.MultiWriter(f)

	if err := routes.SetupRotes(router); err != nil {
		fmt.Println(err.Error())
		return
	}

	// Test Transaction
	tx, _ := db_connection.Dbx.Begin(context.Background())
	defer tx.Rollback(context.Background())

	repo := db_repo.New(tx)

	users, _ := repo.GetAllUsers(context.Background())

	for _, user := range users {
		fmt.Printf("%v\n", user.Username)
	}

	if err := tx.Commit(context.Background()); err != nil {
		return
	}

	router.Run(global.GetGinWorkPath())

	fmt.Printf("Hey!, %v", "you")
}
