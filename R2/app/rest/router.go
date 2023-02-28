package rest

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	server := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	server.Use(sessions.Sessions("mysession", store))

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Login and logout routes
	server.POST("/login", login)
	server.GET("/logout", logout)

	// writer routes
	spiral := server.Group("/spiral")
	{
		spiral.Use(AuthRequired())
		spiral.POST("", SpiralFibonacci)
	}

	return server
}
