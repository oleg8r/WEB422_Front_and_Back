package api

import (
	"booksApi/controllers"
	"github.com/gin-gonic/gin"
)

func StartRouter() {
	router := gin.Default()
	router.Use(CORS())

	router.POST("/create", controllers.Create)
	router.GET("/read", controllers.Read)
	router.POST("/update", controllers.Update)
	router.DELETE("/delete", controllers.Delete)

	router.Run(":3000")
}

func CORS() gin.HandlerFunc {
	// TO allow CORS
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
