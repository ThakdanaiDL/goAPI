package repository

import (
	models "github.com/ThakdanaiDL/goAPI/models"

	"gorm.io/gorm"
)

// =====================
// Interface
// =====================
type UserRepository interface {
	GetAll() ([]models.UserData, error)
	GetByID(id string) (models.UserData, error)
	Create(data models.UserData) error
	Update(data models.UserData) error
	Delete(data models.UserData) error
	DeleteAll() error
	//************************* custom function *************//
	GetAll_UserCustom() ([]models.UserSummary, error)
}

// =====================
// Struct
// =====================
type userRepo struct {
	db *gorm.DB
}

// =====================
// Factory
// =====================
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db: db}
}

// =====================
// Implementation
// =====================

// GetAll
func (r *userRepo) GetAll() ([]models.UserData, error) {
	var users []models.UserData
	return users, r.db.Find(&users).Error
}

func (r *userRepo) GetAll_UserCustom() ([]models.UserSummary, error) {
	var users []models.UserSummary
	err := r.db.
		Model(&models.UserData{}).
		Select("id, name, winrate, rank").
		Scan(&users).Error

	return users, err
}

// GetByID
func (r *userRepo) GetByID(id string) (models.UserData, error) {
	var user models.UserData
	return user, r.db.First(&user, id).Error
}

// Create
func (r *userRepo) Create(data models.UserData) error {
	return r.db.Create(&data).Error
}

// Update
// ใช้กับทั้ง update ข้อมูลทั่วไป + update LastPlayedRound
func (r *userRepo) Update(data models.UserData) error {
	return r.db.Save(&data).Error
}

// Delete
func (r *userRepo) Delete(data models.UserData) error {
	return r.db.Delete(&data).Error
}

// DeleteAll
func (r *userRepo) DeleteAll() error {
	return r.db.Exec("DELETE FROM user_data").Error
}
