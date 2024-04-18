package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/Anuolu-2020/Expense-Calculator-App/models"
)

func Init() *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  false,         // Disable color
		},
	)

	// Get go environment
	Go_ENV := os.Getenv("GO_ENV")

	var DATABASE_URL string

	if Go_ENV == "development" {
		DATABASE_URL = os.Getenv("DATABASE_URL")
	} else {
		DATABASE_URL = os.Getenv("LOCAL_DB_URL")
	}

	db, err := gorm.Open(postgres.Open(DATABASE_URL), &gorm.Config{Logger: newLogger})
	if err != nil {
		fmt.Println("Error connecting to Database")
		log.Fatal(err)
	}

	log.Println("DB migrating......")
	db.AutoMigrate(&models.User{}, &models.Report{})
	log.Println("DB migrated")

	log.Println("Database connected.....")

	return db
}
