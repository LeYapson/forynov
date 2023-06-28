package database

import (
	"log"
	"main/models"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("forum.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Impossible de se connecter à la base de données! \n", err.Error())
		os.Exit(2)
	}

	log.Println("connecté à la base de données de manière parfaitement parfaite")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Immigration des données en cours")
	// add migration
	db.AutoMigrate(&models.User{}, &models.Moderator{}, &models.UserProfile{}, &models.Subject{}, &models.Message{})

	Database = DbInstance{Db: db}
}
