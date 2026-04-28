package main

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"time"

	"vinyl-store-api/middleware"
	"vinyl-store-api/models"
	"vinyl-store-api/store"

	"github.com/gin-gonic/gin"
)

func generateToken() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Vinyl Store API is running",
		})
	})

	router.GET("/login", func(c *gin.Context) {
		username, password, hasAuth := c.Request.BasicAuth()

		if !hasAuth {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Username and password are required",
			})
			return
		}

		expectedPassword, exists := store.Users[username]

		if !exists || expectedPassword != password {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid username or password",
			})
			return
		}

		token := generateToken()

		store.Mutex.Lock()
		store.Tokens[token] = username
		store.Mutex.Unlock()

		c.JSON(http.StatusOK, gin.H{
			"message": "Hi " + username + ", welcome to the Store System",
			"token":   token,
		})
	})

	protected := router.Group("/")
	protected.Use(middleware.AuthRequired())

	protected.GET("/logout", func(c *gin.Context) {
		username := c.GetString("username")
		token := c.GetString("token")

		store.Mutex.Lock()
		delete(store.Tokens, token)
		store.Mutex.Unlock()

		c.JSON(http.StatusOK, gin.H{
			"message": "Bye " + username + ", your token has been revoked",
		})
	})

	protected.GET("/status", func(c *gin.Context) {
		username := c.GetString("username")

		currentTime := time.Now().Format("2006-01-02 15:04:05")

		c.JSON(http.StatusOK, gin.H{
			"message": "Hi " + username + ", the DPIP System is Up and Running",
			"time":    currentTime,
		})
	})

	protected.GET("/albums", func(c *gin.Context) {
		c.JSON(http.StatusOK, store.Albums)
	})

	protected.GET("/albums/:id", func(c *gin.Context) {
		id := c.Param("id")

		for _, album := range store.Albums {
			if album.ID == id {
				c.JSON(http.StatusOK, album)
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{
			"error": "Album not found",
		})
	})

	protected.POST("/post-album", func(c *gin.Context) {
		var newAlbum models.Album

		err := c.BindJSON(&newAlbum)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid JSON format",
			})
			return
		}

		if newAlbum.ID == "" || newAlbum.Title == "" || newAlbum.Artist == "" || newAlbum.Price <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "All fields are required and price must be greater than 0",
			})
			return
		}

		for _, album := range store.Albums {
			if album.ID == newAlbum.ID {
				c.JSON(http.StatusConflict, gin.H{
					"error": "Album ID already exists",
				})
				return
			}
		}

		store.Albums = append(store.Albums, newAlbum)

		c.JSON(http.StatusCreated, newAlbum)
	})

	router.Run(":8080")
}
