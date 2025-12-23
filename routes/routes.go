package routes

import (
	"github.com/ThakdanaiDL/goAPI/controller"

	"github.com/labstack/echo/v4"
)

func Register(e *echo.Echo, ctrl *controller.MessageController) {

	// default page
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Welcome to My Go API on Render!")
	})

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, echo.Map{
			"status":  "UP",
			"message": "Echo API is running",
		})
	})

	e.GET("/history", ctrl.GetAll)
	e.GET("/update", ctrl.Update)
	e.GET("/delete", ctrl.Delete)
	e.GET("/delete-all", ctrl.DeleteAll)
	e.GET("/send", ctrl.Send)
}
