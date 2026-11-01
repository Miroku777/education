package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	_ "lesson22/docs"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const Port = ":8080"

const (
	Timeout = 2 * time.Second // Connection timeout
	DB      = "./auth.db"     // Path to DB
)

// User model info
// @Description User account information
type User struct {
	Email    string `json:"email" example:"user@example.com"`
	Password string `json:"password" example:"secret"`
}

// ContactRequest model info
// @Description Contact form payload
type ContactRequest struct {
	Email   string `json:"email" example:"user@example.com"`
	Message string `json:"message" example:"Hello, world!"`
}

// IndexHandler - index handler
// @Summary      Base Index page
// @Description  Returns HTML Index page
// @Tags         Base
// @Produce      html
// @Success      200  {string}  string  "Success html render"
// @Router       / [get]
func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"name": "INDEX",
	})
}

// ApiUsersHandler - get user by id
// @Summary      Get user by id
// @Description  API handler for getting user info by id
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        id    path      string  true   "User ID"
// @Param        name  query     string  false  "User name filter"
// @Success      200   {object}  map[string]string  "Success getting user info"
// @Failure      400   {object}  map[string]string  "Bad request"
// @Router       /api/users/{id} [get]
func ApiUsersHandler(c *gin.Context) {
	id := c.Param("id")
	name := c.Query("name")
	c.JSON(http.StatusOK, gin.H{
		"id":       id,
		"name":     name,
		"password": "123456",
	})
}

// ContactsPostHandler - submit contact form
// @Summary      Submit contact form
// @Description  Accepts contact form data and returns status
// @Tags         contacts
// @Accept       json
// @Produce      json
// @Param        message  body      ContactRequest  true  "Contact payload"
// @Success      201      {object}  map[string]string  "Created"
// @Failure      400      {object}  map[string]string  "Bad request"
// @Router       /contacts [post]
func ContactsPostHandler(c *gin.Context) {
	var Message struct {
		Email   string
		Message string
	}
	if err := c.ShouldBindJSON(&Message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println(Message)
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("NEW REQ:", c.Request.Method, c.Request.URL, time.Now().Format("2006/01/02 15:04"))
		c.Next()
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("HELLO!!!")
		c.Next()
	}
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.Use(LoggerMiddleware())

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/", IndexHandler)
	r.POST("/contacts", ContactsPostHandler)
	r.GET("/api/users/:id", ApiUsersHandler)

	if err := r.Run(Port); err != nil {
		log.Println(err)
		return
	}
}
