package global

import (
	"fmt"
	"os"
)

var PathForCookies = "/"
var DomainForCookies = "127.0.0.1"

const SecureForCookies = false
const HttpOnlyForCookies = false

// Change to 0.0.0.0 when docker this
const GinWorkPath = "127.0.0.1:9876"

func GetGinWorkPath() string {

	if os.Getenv("DOMAIN") != "" {
		DomainForCookies = os.Getenv("DOMAIN")
		PathForCookies = "/api/v1/"
		fmt.Println("DomainForCookies is", DomainForCookies)
	}

	if os.Getenv("ENV") == "docker" {
		fmt.Println("GinWorkPath is docker 0.0.0.0:9876")
		return "0.0.0.0:9876"
	}

	fmt.Println("GinWorkPath is local 127.0.0.1:9876")
	return GinWorkPath
}
