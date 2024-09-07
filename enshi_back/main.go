package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	var err error

	if err = godotenv.Load("secret.env"); err != nil {
		fmt.Printf("%v", err)
		return
	}
	var test1 string
	var exists bool
	if test1, exists = os.LookupEnv("TEST1"); !exists {
		fmt.Printf("%v", err)
		return
	}

	fmt.Printf("Hey!, %v", test1)
}
