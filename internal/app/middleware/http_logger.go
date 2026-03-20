package middleware

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/quangdtptit/go-cli/pkg/logger"
	"go.uber.org/zap"
)

func LoggerMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			req := c.Request()
			res := c.Response()
			// traceId := uuid.New().String()
			//l := logger.Logger.With(
			//	zap.String("traceId", traceId),
			//)

			logger.Logger.Info("[LogMiddleware] request",
				zap.String("method", req.Method),
				zap.String("path", req.URL.Path),
				zap.String("ip", c.RealIP()),
			)

			c.SetResponse(res)
			err := next(c)

			latency := time.Since(start)

			logger.Logger.Info("[LogMiddleware] response",
				zap.String("method", req.Method),
				zap.String("path", req.URL.Path),
				zap.Int("status", res.Status),
				zap.Duration("latency", latency),
			)
			return err
		}
	}
}
