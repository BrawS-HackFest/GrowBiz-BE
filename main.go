package main

import (
	"HackFest/database"
	"HackFest/handler/route"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Failed to load env file")
	}
	db := database.InitDb()
	database.MigrateDb(db)

	r := gin.Default()
	route.Route(r)
	r.Run()
}
