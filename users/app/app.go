package app

import (
	"fmt"
	"log"
	"os"
	"strings"
	"users/domain"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	r = gin.Default()
)

func StartApp() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	userPort := os.Getenv("USERPORT")

	fmt.Println(userPort)

	if strings.TrimSpace(userPort) == "" {
		log.Fatal("userPort doesn't exist!!!")
	}

	Routers()
	domain.ConnDB()
	r.Run(userPort)
}
