package service

import (
	models "github.com/ThakdanaiDL/goAPI/models"
	"github.com/ThakdanaiDL/goAPI/repository"
	"github.com/ThakdanaiDL/goAPI/utils"
)

const WINRATE_THRESHOLD = 20

// =====================
// Interface
// =====================
type UserService interface {
	GetAll() ([]models.UserData, error)
	Update(id string, name, rank *string) (models.UserData, error)
	Delete(id string) error
	DeleteAll() error
	CreateAndSend(msg string) error
	UserListing() ([]models.UserSummary, error)
}

// =====================
// Service struct
// =====================
type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{repo: r}
}

// =====================
// Basic CRUD
// =====================
func (s *userService) GetAll() ([]models.UserData, error) {
	return s.repo.GetAll()
}

func (s *userService) Update(id string, name, rank *string) (models.UserData, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return user, err
	}

	if name != nil {
		user.Name = *name
	}

	if rank != nil {
		user.Rank = *rank
	}

	return user, s.repo.Update(user)
}

func (s *userService) Delete(id string) error {
	u, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(u)
}

func (s *userService) DeleteAll() error {
	return s.repo.DeleteAll()
}

func (s *userService) CreateAndSend(newUser string) error {
	if err := utils.Send(newUser); err != nil {
		return err
	}

	return s.repo.Create(models.UserData{
		Name: newUser,
		Rank: "",
	})
}

func (s *userService) UserListing() ([]models.UserSummary, error) {
	return s.repo.GetAll_UserCustom()
}
