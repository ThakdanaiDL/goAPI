package controller

import (
	"github.com/ThakdanaiDL/goAPI/service"

	"net/http"

	"github.com/labstack/echo/v4"
)

type MessageController struct {
	Svc service.MessageService
}

func NewMessageController(s service.MessageService) *MessageController {
	return &MessageController{Svc: s}
}

func (h *MessageController) GetAll(c echo.Context) error {
	data, _ := h.Svc.GetAll()
	return c.JSON(http.StatusOK, data)
}

func (h *MessageController) Update(c echo.Context) error {
	id := c.QueryParam("id")
	msg := c.QueryParam("msg")

	data, err := h.Svc.Update(id, msg)
	if err != nil {
		return c.JSON(404, echo.Map{"error": "ไม่พบ ID นี้"})
	}

	return c.JSON(200, data)
}

func (h *MessageController) Delete(c echo.Context) error {
	id := c.QueryParam("id")
	err := h.Svc.Delete(id)
	if err != nil {
		return c.JSON(404, echo.Map{"error": "ไม่พบข้อมูล"})
	}

	return c.JSON(200, echo.Map{"status": "ลบเรียบร้อย"})
}
