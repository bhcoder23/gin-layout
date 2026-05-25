// Package routers registers HTTP routes.
package routers

import (
	"github.com/bhcoder23/gin-layout/internal/middlewares"
	"github.com/gin-gonic/gin"
)

type Router func(*gin.Engine)

var routers = []Router{}

func InitRouter() *gin.Engine {
	r := gin.Default()
	middlewares.InitMiddlewares(r)
	for _, route := range routers {
		route(r)
	}
	return r
}

func RegisterRoute(r ...Router) {
	routers = append(routers, r...)
}
