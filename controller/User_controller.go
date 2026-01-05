package controller

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/ThakdanaiDL/goAPI/service"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	Svc service.UserService
}

func NewUserController(s service.UserService) *UserController {
	return &UserController{Svc: s}
}

// =====================
// Basic CRUD
// =====================
func (h *UserController) GetAll(c echo.Context) error {
	data, err := h.Svc.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, data)
}

func (h *UserController) Update(c echo.Context) error {
	id := c.QueryParam("id")
	params := c.QueryParams()

	var namePtr, rankPtr *string

	if params.Has("name") {
		v := params.Get("name")
		namePtr = &v
	}

	if params.Has("rank") {
		v := params.Get("rank")
		rankPtr = &v
	}

	data, err := h.Svc.Update(id, namePtr, rankPtr)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "ไม่พบ ID นี้"})
	}
	return c.JSON(http.StatusOK, data)
}

func (h *UserController) Delete(c echo.Context) error {
	id := c.QueryParam("id")
	if err := h.Svc.Delete(id); err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "ไม่พบข้อมูล"})
	}
	return c.JSON(http.StatusOK, echo.Map{"status": "ลบเรียบร้อย"})
}

func (h *UserController) DeleteAll(c echo.Context) error {
	h.Svc.DeleteAll()
	return c.JSON(http.StatusOK, echo.Map{"status": "ล้างข้อมูลทั้งหมดแล้ว"})
}

func (h *UserController) AddUser(c echo.Context) error {
	name := c.QueryParam("name")
	if name == "" {
		name = "NameDummy" + strconv.Itoa(rand.Intn(10000))
	}

	if err := h.Svc.CreateAndSend(name); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "สร้างผู้เล่นไม่สำเร็จ"})
	}
	return c.JSON(http.StatusOK, echo.Map{"status": "เพิ่มผู้เล่นแล้ว"})
}

func (h *UserController) UserListing(c echo.Context) error {
	data, err := h.Svc.UserListing()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, data)
}
