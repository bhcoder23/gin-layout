// Package components initializes shared infrastructure components.
package components

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InitLog() error {
	var logger *zap.Logger
	var err error
	if gin.Mode() == gin.DebugMode {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}
	if err != nil {
		return err
	}
	zap.ReplaceGlobals(logger)
	return nil
}
