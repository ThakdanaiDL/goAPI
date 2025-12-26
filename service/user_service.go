package service

import (
	"errors"
	"math"
	"sort"
	"strconv"

	models "github.com/ThakdanaiDL/goAPI/models"
	"github.com/ThakdanaiDL/goAPI/repository"
	"github.com/ThakdanaiDL/goAPI/utils"
)

type UserService interface {
	GetAll() ([]models.UserData, error)
	Update(id string, name, winrate, rank *string) (models.UserData, error)
	Delete(id string) error
	DeleteAll() error
	CreateAndSend(msg string) error
	FindClosestCluster4() ([]models.UserData, error)
	MakeTeams(players []models.UserData) (map[string][]models.UserData, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) GetAll() ([]models.UserData, error) {
	return s.repo.GetAll()
}

func (s *userService) Update(id string, name, winrate, rank *string) (models.UserData, error) {
	Userdata, err := s.repo.GetByID(id)
	if err != nil {
		return Userdata, err
	}
	if name != nil {
		Userdata.Name = *name
	}

	if winrate != nil {
		Userdata.Winrate = *winrate
	}
	if rank != nil {
		Userdata.Rank = *rank

	}
	err = s.repo.Update(Userdata)

	return Userdata, err
}

func (s *userService) Delete(id string) error {
	log, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(log)
}

func (s *userService) DeleteAll() error {
	return s.repo.DeleteAll()
}

func (s *userService) CreateAndSend(msg string) error {
	if err := utils.Send(msg); err != nil {
		return err
	}
	return s.repo.Create(models.UserData{
		Name:    msg,
		Rank:    "",
		Winrate: "",
	})
}

func (s *userService) FindClosestCluster4() ([]models.UserData, error) {
	users, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	if len(users) < 4 {
		return users, nil
	}

	// แปลง winrate เป็นตัวเลข
	type uw struct {
		models.UserData
		W int
	}

	list := make([]uw, 0, len(users))

	for _, u := range users {
		w, _ := strconv.Atoi(u.Winrate)
		list = append(list, uw{UserData: u, W: w})
	}

	// sort จาก winrate น้อย → มาก
	sort.Slice(list, func(i, j int) bool {
		return list[i].W < list[j].W
	})

	// sliding window หา cluster 4 คนที่ใกล้สุด
	bestStart := 0
	bestRange := math.MaxInt32

	for i := 0; i <= len(list)-4; i++ {
		diff := list[i+3].W - list[i].W
		if diff < bestRange {
			bestRange = diff
			bestStart = i
		}
	}

	// return 4 คนที่ winrate ใกล้ที่สุด
	result := []models.UserData{
		list[bestStart].UserData,
		list[bestStart+1].UserData,
		list[bestStart+2].UserData,
		list[bestStart+3].UserData,
	}

	return result, nil
}

func (s *userService) MakeTeams(players []models.UserData) (map[string][]models.UserData, error) {
	if len(players) != 4 {
		return nil, errors.New("ต้องมีผู้เล่น 4 คนพอดี")
	}

	// แปลง winrate เป็นตัวเลข
	w := make([]int, 4)
	for i, p := range players {
		w[i], _ = strconv.Atoi(p.Winrate)
	}

	// ความเป็นไปได้ของทีมทั้งหมด
	combos := []struct {
		A [2]int
		B [2]int
	}{
		{A: [2]int{0, 1}, B: [2]int{2, 3}}, // A+B vs C+D
		{A: [2]int{0, 2}, B: [2]int{1, 3}}, // A+C vs B+D
		{A: [2]int{0, 3}, B: [2]int{1, 2}}, // A+D vs B+C
	}

	bestDiff := math.MaxInt32
	bestCombo := 0

	for i, c := range combos {
		sumA := w[c.A[0]] + w[c.A[1]]
		sumB := w[c.B[0]] + w[c.B[1]]
		diff := abs(sumA - sumB)

		if diff < bestDiff {
			bestDiff = diff
			bestCombo = i
		}
	}

	// เลือกคู่ที่ดีที่สุด
	c := combos[bestCombo]

	teamA := []models.UserData{
		players[c.A[0]],
		players[c.A[1]],
	}

	teamB := []models.UserData{
		players[c.B[0]],
		players[c.B[1]],
	}

	return map[string][]models.UserData{
		"teamA": teamA,
		"teamB": teamB,
	}, nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
