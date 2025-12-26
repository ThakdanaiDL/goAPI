package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	// dsn := os.Getenv("DATABASE_URL")
	dsn := "postgresql://renderpostgresdb_i3ir_user:9PV4VeS4HJfl1Cu4qiVu7XsNqTIwftGz@dpg-d55ojlpr0fns73d19o50-a.singapore-postgres.render.com/renderpostgresdb_i3ir" //for test

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
