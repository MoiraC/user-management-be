package main

import (
	"log"
	"net/http"

	"ayse.com/user-management/models"
	"github.com/gin-gonic/gin"
)

func main() {
	err := models.ConnectDatabase()
	checkErr(err)

	r := gin.Default()

	// API Blogs
	router := r.Group("/users")
	{
		router.POST("/create", postUser)
		router.GET("/", readUser)
		router.POST("/update/:id", updateUser)
		router.DELETE("/delete/:id", deleteUser)
	}

	// By default it serves on :8080
	r.Run()
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func readUser(c *gin.Context) {
	users, err := models.GetUsers()
	checkErr(err)

	if users == nil {
		c.JSON(404, gin.H{"error": "No Records Found"})
		return
	} else {
		c.JSON(200, gin.H{"data": users})
	}
}

func updateUser(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Record Updated!"})
}
func deleteUser(c *gin.Context) {
	id := c.Param("id")

	success, err := models.RemoveUser(id)

	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Record Deleted!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}

func postUser(c *gin.Context) {

	var json models.User

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	success, err := models.AddUser(json)

	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}
