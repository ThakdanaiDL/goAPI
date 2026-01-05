// package main

// import (
// 	"log"
// 	"net/http"
// 	"os"

// 	"bytes"
// 	"encoding/json"

// 	"github.com/labstack/echo/v4"
// 	"github.com/labstack/echo/v4/middleware"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// type MessageLog struct {
// 	gorm.Model
// 	Content string `json:"content"`
// 	Status  string `json:"status"`
// }

// var db *gorm.DB

// func main() {

// 	dsn := os.Getenv("DATABASE_URL")

// 	dialector := postgres.New(postgres.Config{
// 		DSN:                  dsn,
// 		PreferSimpleProtocol: true, // ⭐ สำคัญที่สุด
// 	})

// 	var err error
// 	db, err = gorm.Open(dialector, &gorm.Config{
// 		PrepareStmt: false, // ปิด GORM prepare
// 	})
// 	if err != nil {
// 		log.Fatal("Failed to connect to database: ", err)
// 	}

// 	// Auto migrate
// 	if err := db.AutoMigrate(&MessageLog{}); err != nil {
// 		log.Fatal("AutoMigrate failed: ", err)
// 	}

// 	e := echo.New()

// 	// Middleware
// 	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
// 		AllowOrigins: []string{"*"}, // หรือใส่ URL เฉพาะของหน้าเว็บคุณเพื่อความปลอดภัย
// 		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
// 	}))
// 	e.Use(middleware.Logger())
// 	e.Use(middleware.Recover())

// 	//read allกดกกด
// 	e.GET("/history", func(c echo.Context) error {
// 		var logs []MessageLog
// 		db.Find(&logs)
// 		return c.JSON(http.StatusOK, logs)
// 	})

// 	// 2. UPDATE - แก้ไขข้อความตาม ID
// 	// วิธีเรียก: /update?id=1&msg=ข้อความใหม่
// 	e.GET("/update", func(c echo.Context) error {
// 		id := c.QueryParam("id")
// 		newMsg := c.QueryParam("msg")

// 		var log MessageLog
// 		if err := db.First(&log, id).Error; err != nil {
// 			return c.JSON(http.StatusNotFound, map[string]string{"error": "ไม่พบ ID นี้"})
// 		}

// 		log.Content = newMsg
// 		db.Save(&log) // บันทึกการแก้ไข

// 		return c.JSON(http.StatusOK, map[string]interface{}{
// 			"message": "อัปเดตเรียบร้อย!",
// 			"data":    log,
// 		})
// 	})

// 	// 3. DELETE - ลบข้อมูลตาม ID
// 	// วิธีเรียก: /delete?id=1
// 	e.GET("/delete", func(c echo.Context) error {
// 		id := c.QueryParam("id")

// 		// ค้นหาดูก่อนว่ามีไหม
// 		var log MessageLog
// 		if err := db.First(&log, id).Error; err != nil {
// 			return c.JSON(http.StatusNotFound, map[string]string{"error": "ไม่พบข้อมูลที่ต้องการลบ"})
// 		}

// 		// ลบข้อมูล (GORM จะทำ Soft Delete ถ้ามี gorm.Model)
// 		db.Delete(&log)

// 		return c.JSON(http.StatusOK, map[string]string{"status": "ลบ ID " + id + " เรียบร้อยแล้ว"})
// 	})

// 	// 4. DELETE ALL - ล้างฐานข้อมูล (Option)
// 	e.GET("/delete-all", func(c echo.Context) error {
// 		// ลบทุกอย่างในตาราง
// 		db.Exec("DELETE FROM message_logs")
// 		return c.JSON(http.StatusOK, map[string]string{"status": "ล้างข้อมูลทั้งหมดแล้ว"})
// 	})

// 	e.GET("/send", func(c echo.Context) error {
// 		msg := c.QueryParam("msg")
// 		if msg == "" {
// 			msg = "สวัสดีจาก Go API บน Railway! 🚀"
// 		}

// 		err := sendDiscordNotify(msg)
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "ส่งไม่สำเร็จ"})
// 		}

// 		db.Create(&MessageLog{Content: msg, Status: "Sent"})

// 		return c.JSON(http.StatusOK, map[string]string{"status": "ส่งเข้า Discord แล้ว!"})
// 	})

// 	e.GET("/health", func(c echo.Context) error {
// 		return c.JSON(http.StatusOK, map[string]string{
// 			"status":  "UP",
// 			"message": "Echo API is running on Cloud Run!",
// 		})
// 	})

// 	e.GET("/", func(c echo.Context) error {
// 		return c.String(http.StatusOK, "Welcome to My Go API on Render!")
// 	})

// 	// สำคัญ: Cloud Run จะกำหนด Port ผ่าน Environment Variable ชื่อ PORT
// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = "8080"
// 	}

// 	//ddd
// 	e.Logger.Fatal(e.Start(":" + port))
// }

// func sendDiscordNotify(message string) error {
// 	webhookURL := os.Getenv("DISCORD_WEBHOOK_URL") //แนะนำให้ใช้ เพื่อความปลอดภัย

// 	payload := map[string]string{
// 		"content": message,
// 	}
// 	jsonPayload, _ := json.Marshal(payload)

// 	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonPayload))
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	return nil
// }

package main

import (
	"net/http"
	"os"

	"github.com/ThakdanaiDL/goAPI/config"
	"github.com/ThakdanaiDL/goAPI/controller"
	models "github.com/ThakdanaiDL/goAPI/models"
	"github.com/ThakdanaiDL/goAPI/repository"
	"github.com/ThakdanaiDL/goAPI/routes"
	"github.com/ThakdanaiDL/goAPI/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	db := config.ConnectDatabase()
	db.AutoMigrate(&models.MessageLog{}, &models.UserData{})

	repo := repository.NewMessageRepository(db)
	svc := service.NewMessageService(repo)
	ctrl := controller.NewMessageController(svc)

	User_repo := repository.NewUserRepository(db)
	User_svc := service.NewUserService(User_repo)
	User_ctrl := controller.NewUserController(User_svc)

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://godash.onrender.com"},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodOptions, // แนะนำให้ใส่
		},
		AllowHeaders: []string{
			"Content-Type",
			"Authorization",
		},
		AllowCredentials: true, // ไม่จำเป็นก็ได้ แต่ถ้าคุณมี cookie หรือ credential ให้เปิด
	}))

	e.HEAD("/", func(c echo.Context) error {
		return c.NoContent(200)
	})
	e.HEAD("/health", func(c echo.Context) error {
		return c.NoContent(200)
	})

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.Register(e, ctrl)

	routes.UserRegister(e, User_ctrl)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8089"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
