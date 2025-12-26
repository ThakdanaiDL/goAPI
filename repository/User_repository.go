package repository

import (
	models "github.com/ThakdanaiDL/goAPI/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll() ([]models.UserData, error)
	GetByID(id string) (models.UserData, error)
	Create(data models.UserData) error
	Update(data models.UserData) error
	Delete(data models.UserData) error
	DeleteAll() error
	// Mathmaking() error
}

type userRepo struct {
	db *gorm.DB
}

//******************* factory ************************//

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db}
}

//******************* impl ************************//

// func (r *userRepo) Mathmaking() ([]models.UserData, error) {
// 	var userdata []models.UserData
// 	r.db.Find(&userdata).Error

// 	// return userdata, r.db.Find(&userdata).Error
// }

func (r *userRepo) GetAll() ([]models.UserData, error) {
	var logs []models.UserData
	return logs, r.db.Find(&logs).Error
}

func (r *userRepo) GetByID(id string) (models.UserData, error) {
	var data models.UserData
	return data, r.db.First(&data, id).Error
}

func (r *userRepo) Create(data models.UserData) error {
	return r.db.Create(&data).Error
}

func (r *userRepo) Update(data models.UserData) error {
	return r.db.Save(&data).Error
}

func (r *userRepo) Delete(data models.UserData) error {
	return r.db.Delete(&data).Error
}

func (r *userRepo) DeleteAll() error {
	return r.db.Exec("DELETE FROM user_data").Error
}
