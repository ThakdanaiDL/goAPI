package config

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")

	dialector := postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // ⭐ สำคัญที่สุด
	})

	var err error

	db, err := gorm.Open(dialector, &gorm.Config{
		PrepareStmt: false,
	})
	if err != nil {
		log.Fatal("❌ Failed to connect database:", err)
	}

	return db
}
