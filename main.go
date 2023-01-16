package main

//import "fmt"
import (
	"db_crud/controllers"
	"db_crud/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	r := gin.Default()

	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/post/:id", controllers.PostShow)
	r.PUT("/post/:id", controllers.PostUpdate)
	r.DELETE("/post/:id", controllers.PostDelete)

	r.Run()
}
