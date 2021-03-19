package rest

import (
	"discount-service/transport/controllers"
	"github.com/labstack/echo"
)

func (ac *RContext) RegisterRoutes(e *echo.Echo) {
	cc := controllers.ControllerContext{ApplicationContext: &ac.ApplicationContext}
	v1 := e.Group("/v1")

	v1.GET("/hello", cc.Hello)
}
