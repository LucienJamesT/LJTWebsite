package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Serve static files
	router.Static("/assets", "./assets")

	// Handler for the root path
	router.LoadHTMLFiles(
		"templates/index.html",
		"templates/aboutMe.html",
		"templates/resume.html",
		"templates/projects.html",
		"templates/home.html",
	)

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET("/Home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", nil)
	})

	router.GET("/About-Me", func(c *gin.Context) {
		c.HTML(http.StatusOK, "aboutMe.html", nil)
	})

	router.GET("/Resume", func(c *gin.Context) {
		c.HTML(http.StatusOK, "resume.html", nil)
	})

	router.GET("/Projects", func(c *gin.Context) {
		c.HTML(http.StatusOK, "projects.html", nil)
	})

	router.Run("0.0.0.0:8080")
}
