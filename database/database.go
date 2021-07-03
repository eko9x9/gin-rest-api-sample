package database

import (
	"fmt"
	"os"

	"github.com/eko9x9/gin-rest-api-sample/database/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Initialize() (*gorm.DB, error) {
	conf := fmt.Sprintf("host=%s user=%s password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	dsn := conf
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	migrations.Migrate(db)
	return db, err
}
