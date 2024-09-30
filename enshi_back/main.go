package main

import (
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

	router := gin.Default()
	if err := utils.SetupRotes(router); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Hey!, %v", "you")
}
