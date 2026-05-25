package routers

import (
	"github.com/bhcoder23/gin-layout/internal/controllers"
	"github.com/bhcoder23/gin-layout/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func init() {
	RegisterRoute(func(e *gin.Engine) {
		v1 := e.Group("/api/v1")
		v1.Use(middlewares.AuthMiddleware())
		goods := v1.Group("/goods")
		{
			goods.GET("", controllers.GoodsPageList)
			goods.GET("/:id", controllers.GoodsOne)
			goods.POST("", controllers.GoodsAdd)
			goods.PUT("/:id", controllers.GoodsUpdate)
			goods.DELETE("/:id", controllers.GoodsDel)
		}
	})
}
