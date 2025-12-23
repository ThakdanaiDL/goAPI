package repository

import (
	models "github.com/ThakdanaiDL/goAPI/models"

	"gorm.io/gorm"
)

type MessageRepository interface {
	GetAll() ([]models.MessageLog, error)
	GetByID(id string) (models.MessageLog, error)
	Create(log models.MessageLog) error
	Update(log models.MessageLog) error
	Delete(log models.MessageLog) error
	DeleteAll() error
}

type messageRepo struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) MessageRepository {
	return &messageRepo{db}
}

func (r *messageRepo) GetAll() ([]models.MessageLog, error) {
	var logs []models.MessageLog
	return logs, r.db.Find(&logs).Error
}

func (r *messageRepo) GetByID(id string) (models.MessageLog, error) {
	var log models.MessageLog
	return log, r.db.First(&log, id).Error
}

func (r *messageRepo) Create(log models.MessageLog) error {
	return r.db.Create(&log).Error
}

func (r *messageRepo) Update(log models.MessageLog) error {
	return r.db.Save(&log).Error
}

func (r *messageRepo) Delete(log models.MessageLog) error {
	return r.db.Delete(&log).Error
}

func (r *messageRepo) DeleteAll() error {
	return r.db.Exec("DELETE FROM message_logs").Error
}
