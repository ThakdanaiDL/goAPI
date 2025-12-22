package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Health Check Route
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status":  "UP",
			"message": "Echo API is running on Cloud Run!",
		})
	})

	e.GET("/trigger", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status":  "triggerEndpoint",
			"message": "trigger!!!",
		})
	})

	// สำคัญ: Cloud Run จะกำหนด Port ผ่าน Environment Variable ชื่อ PORT
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	//ddd
	e.Logger.Fatal(e.Start(":" + port))
}
