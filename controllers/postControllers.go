package controllers

import (
	"net/http"

	"github.com/aayy07/go-gin-gorm/initializers"
	"github.com/aayy07/go-gin-gorm/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	//Get Data of a Request body
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	//Create a POST
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	//Return it
	c.JSON(http.StatusOK, gin.H{
		"Post": post,
	})
}

func PostsIndex(c *gin.Context) {
	//Get the POST
	var posts []models.Post
	initializers.DB.Find(&posts)

	//respond with them
	c.JSON(http.StatusOK, gin.H{
		"Posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	//Get ID off URl
	id := c.Param("id")
	//Get the post
	var post models.Post
	initializers.DB.First(&post, id)

	//respond with them
	c.JSON(http.StatusOK, gin.H{
		"Post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	//Get ID off url
	id := c.Param("id")

	//Get data off request body
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)
	//Find the post we're updating
	var post models.Post
	initializers.DB.First(&post, id)
	//Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	//Respond with them
	c.JSON(http.StatusOK, gin.H{
		"Post": post,
	})
}

func PostsDelete(c *gin.Context) {
	//get the ID off URL
	id := c.Param("id")
	//delete post
	initializers.DB.Delete(&models.Post{}, id)
	//respond
	c.Status(200)
}
