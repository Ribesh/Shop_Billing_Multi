package database

import (
	"fmt"
	"log"

	"github.com/Ribesh/gorm/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	USER   = "ribeshshr"
	PASS   = "P@ssw0rd01"
	HOST   = "localhost"
	DBNAME = "gorm-main"
)

func Connect() *gorm.DB {

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", USER, PASS, HOST, DBNAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to DB")
		return nil
	}

	//Remove Comment to Drop Table

	// db.Migrator().DropTable(&models.Customer{})
	// db.Debug().Migrator().DropTable("customer")

	//Remove Comment to Drop Table
	// db.Debug().Migrator().DropTable(&models.Product{})
	// db.Debug().Migrator().DropTable("product")

	db.Debug().AutoMigrate(&models.Product{}, &models.Customer{})

	// db.Debug().Migrator().CreateConstraint(&models.Customer{}, "Product")
	// db.Debug().Migrator().CreateConstraint(&models.Customer{}, "fk_users_credit_cards")

	//db.Debug().AutoMigrate(&models.Customer{}, &models.Product{})

	return db

}
