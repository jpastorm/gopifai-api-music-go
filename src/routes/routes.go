package routes

import (
	"github.com/labstack/echo"
	"wopifai/src/controllers"
)

func SetUpRoutes(e *echo.Echo) {
	explorer := e.Group("/explorer")
	explorer.GET("",controllers.GetDir)
	explorer.GET("/cover",controllers.GetCover)
	explorer.GET("/track",controllers.GetTrack)
}
