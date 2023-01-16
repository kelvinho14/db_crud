package main

import (
	"db_crud/initializers"
	"db_crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
