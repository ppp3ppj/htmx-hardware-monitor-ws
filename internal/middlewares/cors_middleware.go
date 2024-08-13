package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CORSMiddleware(allowOrigins []string) echo.MiddlewareFunc {
    allowOrigins = append(allowOrigins, "*://localhost:*")
    return middleware.CORSWithConfig(middleware.CORSConfig{
        Skipper: middleware.DefaultSkipper,
        AllowOrigins: allowOrigins,
        AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.PATCH, echo.DELETE},
        AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
    })
}
