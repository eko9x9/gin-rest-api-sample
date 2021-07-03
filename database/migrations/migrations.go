package migrations

import (
	"fmt"

	"github.com/eko9x9/gin-rest-api-sample/app/models"
	"gorm.io/gorm"
)

var model models.Model

func Migrate(db *gorm.DB) {

	if db.Migrator().HasTable(model.User) {
		db.Migrator().DropTable(model.User)
	}
	defer RunMigrate(db)

}

func RunMigrate(db *gorm.DB) {
	db.AutoMigrate(model.User)
	fmt.Println("Auto Migration has been finished")
}
