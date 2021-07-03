package database

import (
	"github.com/eko9x9/gin-rest-api-sample/database/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Initialize() (*gorm.DB, error) {
	dsn := "host=postgres user=postgres password=postgres dbname=sample_api port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	migrations.Migrate(db)
	return db, err
}
