package main

import (
	"github.com/aayy07/go-gin-gorm/initializers"
	"github.com/aayy07/go-gin-gorm/models"
)

func init() {
	initializers.ConnectToDB()
	initializers.LoadEnvVariables()

}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
