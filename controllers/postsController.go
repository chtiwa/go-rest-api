package controllers

import (
	"net/http"

	"github.com/chtiwa/gin_gorm/initializers"
	"github.com/chtiwa/gin_gorm/models"
	"github.com/gin-gonic/gin"
)

type Body struct {
	Title string
	Body  string
}

func PostsCreate(c *gin.Context) {
	var json Body

	if err := c.ShouldBind(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := models.Post{Title: json.Title, Body: json.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	id := c.Param("id")
	var json Body

	c.Bind(&json)
	var post models.Post

	initializers.DB.First(&post, id)
	initializers.DB.Model(&post).Updates(models.Post{
		Title: json.Title,
		Body:  json.Body,
	})

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Where("ID = ?", id).Delete(&models.Post{})
	c.JSON(200, gin.H{
		"message": "The post was deleted successfully",
	})
}
