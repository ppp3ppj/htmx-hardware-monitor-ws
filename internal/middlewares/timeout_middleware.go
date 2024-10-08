package middlewares

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func TimeOutMiddleware(timeout time.Duration) echo.MiddlewareFunc {
    return middleware.TimeoutWithConfig(middleware.TimeoutConfig{
        Skipper: middleware.DefaultSkipper,
        ErrorMessage: "Error: Request timeout.",
        Timeout: timeout * time.Second,
    })
}
