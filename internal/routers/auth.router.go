package routers

import (
	"github.com/bhcoder23/gin-layout/internal/controllers"
	"github.com/bhcoder23/gin-layout/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func init() {
	RegisterRoute(func(e *gin.Engine) {
		api := e.Group("/api/v1")
		auth := api.Group("/auth")

		auth.POST("/login", controllers.Login)
		auth.POST("/reg", controllers.Register)

		auth.Use(middlewares.AuthMiddleware())
		{
			auth.GET("/me", controllers.Me)
		}
	})
}
