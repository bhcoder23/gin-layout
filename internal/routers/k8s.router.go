package routers

import (
	"github.com/bhcoder23/gin-layout/internal/services"
	"github.com/gin-gonic/gin"
)

func init() {
	RegisterRoute(func(e *gin.Engine) {
		e.GET("/healthz", func(ctx *gin.Context) {
			if services.IsReady() {
				ctx.JSON(200, gin.H{"status": "ok"})
			} else {
				ctx.JSON(503, gin.H{"status": "not ready"})
			}
		})
	})
}
