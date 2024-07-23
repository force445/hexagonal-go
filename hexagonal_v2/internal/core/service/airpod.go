package service

import (
	"errors"
	"hexagonal_v2/internal/core/domain"
	"hexagonal_v2/internal/core/port"
	"log"
)

type airpodService struct {
	repo     port.AirpodRepository
	userrepo port.UserRepository
}

func NewAirpodService(repo port.AirpodRepository, userrepo port.UserRepository) port.AirpodService {
	return &airpodService{repo: repo, userrepo: userrepo}
}

func (a *airpodService) CreateAirpod(airpod *domain.Airpod) error {
	a.repo.CreateAirpod(airpod)
	user, err := a.userrepo.GetUserByID(airpod.UserID)
	user.Airpods = append(user.Airpods, *airpod)
	a.userrepo.UpdateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (a *airpodService) GetAirpodByID(id int64) (*domain.Airpod, error) {
	if id == 0 {
		return nil, errors.New("id is required")
	}

	airpod, err := a.repo.GetAirpodByID(id)
	if err != nil {
		return nil, err
	}

	return airpod, nil
}

func (a *airpodService) GetAirpods() ([]*domain.Airpod, error) {
	return a.repo.GetAirpods()
}

func (a *airpodService) GetAirpodByUserID(id int64) ([]*domain.Airpod, error) {
	if id == 0 {
		return nil, errors.New("user id is required")
	}

	log.Printf("Fetching airpods for user id: %s", id)

	airpods, err := a.repo.GetAirpodByUserID(id)
	if err != nil {
		return nil, err
	}

	return airpods, nil
}

func (a *airpodService) UpdateAirpod(airpod *domain.Airpod) error {
	return a.repo.UpdateAirpod(airpod)
}

func (a *airpodService) DeleteAirpod(id int64) error {
	return a.repo.DeleteAirpod(id)
}
