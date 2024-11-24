package routers

import "github.com/gin-gonic/gin"

// Routers is a function to define all the routes
func Routers(r *gin.Engine) {
	prefix := r.Group("/api/v1/")
	authRouter(prefix)
	userRouter(prefix)

	// health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})
}
