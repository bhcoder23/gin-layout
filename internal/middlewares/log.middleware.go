package middlewares

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		stopTime := time.Since(startTime)
		spendTime := int(stopTime.Milliseconds())

		hostName, err := os.Hostname()
		if err != nil {
			hostName = "Unknown"
		}

		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		userAgent := c.Request.UserAgent()
		dataSize := c.Writer.Size()
		dataSize = max(dataSize, 0)
		method := c.Request.Method
		url := c.Request.RequestURI
		raw := c.Request.URL.RawQuery

		fields := []zap.Field{
			zap.String("hostName", hostName),
			zap.Int("spendTime", spendTime),
			zap.String("path", url),
			zap.String("query", raw),
			zap.String("method", method),
			zap.Int("status", statusCode),
			zap.String("clientIp", clientIP),
			zap.Int("dataSize", dataSize),
			zap.String("userAgent", userAgent),
			zap.String("trace_id", c.GetString(`trace_id`)),
		}

		if len(c.Errors) > 0 {
			fields = append(fields, zap.Strings("errors", c.Errors.Errors()))
		}

		switch {
		case statusCode >= 500:
			// L returns the global Logger, which can be reconfigured with ReplaceGlobals. It's safe for concurrent use.
			zap.L().Error("HTTP Server Error", fields...)
		case statusCode >= 400:
			zap.L().Warn("HTTP Client Error", fields...)
		default:
			zap.L().Info("HTTP Request", fields...)
		}
	}
}
