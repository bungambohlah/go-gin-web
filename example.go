package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	router := gin.Default()
	// disable trusted proxies
	router.SetTrustedProxies(nil)
	// transform string to uppercase in the first letter
	caser := cases.Title(language.Indonesian)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// This handler will match /user/john but will not match /user/ or /user
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello " + caser.String(name),
		})
	})

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := caser.String(name) + " outside of " + action
		c.JSON(http.StatusOK, gin.H{
			"message": message,
		})
	})

	// For each matched request Context will hold the route definition
	router.POST("/user/:name/*action", func(c *gin.Context) {
		fullPath := c.FullPath() == "/user/:name/*action" // true
		c.JSON(http.StatusOK, gin.H{
			"message": fullPath,
		})
	})

	// This handler will add a new router for /user/groups.
	// Exact routes are resolved before param routes, regardless of the order they were defined.
	// Routes starting with /user/groups are never interpreted as /user/:name/... routes
	router.GET("/user/groups", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "The available groups are [...]",
		})
	})
	router.Run() // listen and serve on 0.0.0.0:8080
}
