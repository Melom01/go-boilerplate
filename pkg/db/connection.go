package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Melom01/go-boilerplate/pkg/config"
	"github.com/Melom01/go-boilerplate/pkg/domain"
)

func ConnectDatabase(cfg config.Configuration) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s user=%s dbname=%s port=%s password=%s",
		cfg.DbHost, cfg.DbUser, cfg.DbName, cfg.DbPort, cfg.DbPassword,
	)

	db, err := gorm.Open(
		postgres.Open(psqlInfo), &gorm.Config{
			SkipDefaultTransaction: true,
		},
	)

	if dbMigrationErr := db.AutoMigrate(&domain.Book{}); dbMigrationErr != nil {
		log.Fatalf("Could not migrate database: %v", dbMigrationErr)
	}

	return db, err
}
