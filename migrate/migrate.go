package main

import (
	"github.com/chtiwa/gin_gorm/initializers"
	"github.com/chtiwa/gin_gorm/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
