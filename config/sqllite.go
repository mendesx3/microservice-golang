package config

import (
	"github.com/mendesx3/microservice-golang/repository/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
)

func InitalizeSQLlite() (*gorm.DB, error) {
	dbPath := "./database/main.database"
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	log.Printf("SQLLITE DB PATH: %v\n", dbPath)
	_, err = os.Stat(dbPath)
	if os.IsNotExist(err) {
		err = os.MkdirAll("./database", os.ModePerm)
		if err != nil {
			log.Printf("Error on creating SQLLITE %v\n", err)
			return nil, err
		}

		file, err := os.Create(dbPath)
		if err != nil {
			log.Printf("Error on creating SQLLITE %v\n", err)
			return nil, err
		}
		file.Close()
	}
	log.Printf("SQLLITE DB: %v\n", db)

	if err != nil {
		log.Printf("Error on connecting to SQLLITE %v\n", err)
		return nil, err
	}
	err = db.AutoMigrate(&schemas.Person{})

	if err != nil {
		log.Printf("Error on migrating SQLLITE %v\n", err)
		return nil, err
	}
	log.Printf("db migrated")

	return db, nil
}
