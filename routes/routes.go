package routes

import (
	"github.com/ThakdanaiDL/goAPI/controller"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, ctrl *controller.MessageController) {
	e.GET("/history", ctrl.GetAll)
	e.GET("/update", ctrl.Update)
	e.GET("/delete", ctrl.Delete)
}
