package config

import (
	"fmt"
	"gorm.io/gorm"
	"log"
)

var (
	db *gorm.DB
)

func Init() error {
	var err error
	db, err = InitalizeSQLlite()
	if err != nil {
		return fmt.Errorf("failed to initialize SQLite: %v", err)
	}
	return nil
}

func GetDB() *gorm.DB {
	if db == nil {
		log.Fatal("Database is not initialized. Please call config.Init() first.")
	}
	return db
}
