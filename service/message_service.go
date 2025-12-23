package service

import (
	models "github.com/ThakdanaiDL/goAPI/models"
	"github.com/ThakdanaiDL/goAPI/repository"
)

type MessageService interface {
	GetAll() ([]models.MessageLog, error)
	Update(id string, msg string) (models.MessageLog, error)
	Delete(id string) error
	Create(msg string) error
}

type messageService struct {
	repo repository.MessageRepository
}

func NewMessageService(r repository.MessageRepository) MessageService {
	return &messageService{repo: r}
}

func (s *messageService) GetAll() ([]models.MessageLog, error) {
	return s.repo.GetAll()
}

func (s *messageService) Update(id string, msg string) (models.MessageLog, error) {
	log, err := s.repo.GetByID(id)
	if err != nil {
		return log, err
	}

	log.Content = msg
	err = s.repo.Update(log)
	return log, err
}

func (s *messageService) Delete(id string) error {
	log, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(log)
}

func (s *messageService) Create(msg string) error {
	return s.repo.Create(models.MessageLog{
		Content: msg,
		Status:  "Sent",
	})
}
