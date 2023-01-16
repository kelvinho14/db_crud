package controllers

import (
	"db_crud/initializers"
	"db_crud/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// Get data of req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)
	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Respond with them
	c.JSON(200, gin.H{
		"post": posts,
	})
}

func PostShow(c *gin.Context) {
	// Get ID or url
	id := c.Param("id")

	// Get the posts
	var post models.Post
	initializers.DB.First(&post, id)

	// Respond with them
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	// Get ID or url
	id := c.Param("id")

	// Get the data of req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Find the post we wanna update
	var post models.Post
	initializers.DB.First(&post, id)

	// Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// Respond
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	// Get ID or url
	id := c.Param("id")

	// Delete the post
	initializers.DB.Delete(&models.Post{}, id)

	// Respond
	c.JSON(200, gin.H{
		"message": "Deleted",
	})
}
