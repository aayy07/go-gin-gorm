package initializers

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error

	// dsn := os.Getenv("DSN")
	dsn := "host=satao.db.elephantsql.com user=cgswxlgr password=2ORg0JFnH7fKSxtavG-egu8sQVUK7xm4 dbname=cgswxlgr port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
}
