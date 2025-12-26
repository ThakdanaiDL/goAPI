package controller

import (
	"strconv"

	"github.com/ThakdanaiDL/goAPI/service"

	"net/http"

	"math/rand"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	Svc service.UserService
}

func NewUserController(s service.UserService) *UserController {
	return &UserController{Svc: s}
}

func (h *UserController) GetAll(c echo.Context) error {
	data, _ := h.Svc.GetAll()
	return c.JSON(http.StatusOK, data)
}

func (h *UserController) Update(c echo.Context) error {
	id := c.QueryParam("id")
	params := c.QueryParams()

	var namePtr, winratePtr, rankPtr *string

	if params.Has("name") {
		v := params.Get("name")
		namePtr = &v
	}

	if params.Has("winrate") {
		v := params.Get("winrate")
		winratePtr = &v
	}

	if params.Has("rank") {
		v := params.Get("rank")
		rankPtr = &v
	}

	data, err := h.Svc.Update(id, namePtr, winratePtr, rankPtr)
	if err != nil {
		return c.JSON(404, echo.Map{"error": "ไม่พบ ID นี้"})
	}

	return c.JSON(200, data)
}

func (h *UserController) Delete(c echo.Context) error {
	id := c.QueryParam("id")
	err := h.Svc.Delete(id)
	if err != nil {
		return c.JSON(404, echo.Map{"error": "ไม่พบข้อมูล"})
	}
	return c.JSON(200, echo.Map{"status": "ลบเรียบร้อย"})
}

func (h *UserController) DeleteAll(c echo.Context) error {
	h.Svc.DeleteAll()
	return c.JSON(200, echo.Map{"status": "ล้างข้อมูลทั้งหมดแล้ว"})
}

func (h *UserController) Send(c echo.Context) error {
	msg := c.QueryParam("name")
	if msg == "" {
		msg = "NameDummy" + strconv.Itoa(rand.Intn(10000))
	}

	err := h.Svc.CreateAndSend(msg)
	if err != nil {
		return c.JSON(500, echo.Map{"error": "ส่งไม่สำเร็จ"})
	}

	return c.JSON(200, echo.Map{"status": "ส่งเข้า Discord แล้ว"})
}

func (h *UserController) FindClosest4(c echo.Context) error {
	users, err := h.Svc.FindClosestCluster4()
	if err != nil {
		return c.JSON(500, echo.Map{"error": err.Error()})
	}

	// แปลงเป็น list ของ map[string]string
	resp := make([]map[string]string, len(users))

	for i, u := range users {
		resp[i] = map[string]string{
			"name":    u.Name,
			"winrate": u.Winrate,
			"rank":    u.Rank,
		}
	}

	return c.JSON(200, resp)
}

func (h *UserController) MakeTeam(c echo.Context) error {
	// players = 4 คนจาก logic closest ที่คุณทำไปแล้ว
	players, err := h.Svc.FindClosestCluster4()
	if err != nil {
		return c.JSON(500, echo.Map{"error": err.Error()})
	}

	teams, err := h.Svc.MakeTeams(players)
	if err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	return c.JSON(200, teams)
}
