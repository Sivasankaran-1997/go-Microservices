package app

import (
	"buy/domain"
	"fmt"
	"log"
	"os"
	"strings"

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

	buyPort := os.Getenv("BUYPORT")

	fmt.Println(buyPort)

	if strings.TrimSpace(buyPort) == "" {
		log.Fatal("Buy Port doesn't exist!!!")
	}

	Routers()
	domain.ConnDB()
	r.Run(buyPort)
}
