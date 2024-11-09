package env

import (
	"enshi/auth"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LookupEnv(dest *string, envVar string) error {
	if v, exists := os.LookupEnv(envVar); !exists {
		return fmt.Errorf("%v not found in local env", envVar)
	} else {
		*dest = v
		return nil
	}
}

func LoadEnv(path string) error {

	if err := godotenv.Load(path); err != nil {
		fmt.Printf("%v\n", err)
		return err
	}

	if err := LookupEnv(&auth.SecretKey, "SECRET_KEY"); err != nil {
		fmt.Printf("%v\n", err)
		return err
	}

	fmt.Println("Env loaded")
	return nil
}
