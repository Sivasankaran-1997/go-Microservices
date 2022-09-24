package app

import (
	"fmt"
	"log"
	"os"
	"strings"
	"ticket/domain"

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

	ticketPort := os.Getenv("TICKETPORT")

	fmt.Println(ticketPort)

	if strings.TrimSpace(ticketPort) == "" {
		log.Fatal("ticketPort doesn't exist!!!")
	}

	Routers()
	domain.ConnDB()
	r.Run(ticketPort)
}
