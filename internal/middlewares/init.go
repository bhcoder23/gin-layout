// Package middlewares defines Gin middleware setup.
package middlewares

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func InitMiddlewares(r *gin.Engine) {
	r.Use(RateLimitMiddleware())
	r.Use(requestid.New())
	r.Use(TraceIDMiddleware())
	r.Use(LoggerMiddleware())
}
