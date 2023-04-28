package app

import (
	"log"

	"github.com/vaibhavvikas/stay-sorted/database"
)

func Bootstrap() {
	log.Println("Started Application!")
	LoadEnv()
	db := database.DatabaseConnection()
	router := Router(db)
	router.Run(":5000")
}
