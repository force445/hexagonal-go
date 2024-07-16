package service

import (
	"errors"
	"hexagonal_v2/internal/core/domain"
	"hexagonal_v2/internal/core/port"
	"log"

	"github.com/google/uuid"
)

type airpodService struct {
	repo port.AirpodRepository
}

func NewAirpodService(repo port.AirpodRepository) port.AirpodService {
	return &airpodService{repo: repo}
}

func (a *airpodService) CreateAirpod(airpod *domain.Airpod) error {
	airpod.ID = uuid.New().String()
	return a.repo.CreateAirpod(airpod)
}

func (a *airpodService) GetAirpodByID(id string) (*domain.Airpod, error) {
	return a.repo.GetAirpodByID(id)
}

func (a *airpodService) GetAirpods() ([]*domain.Airpod, error) {
	return a.repo.GetAirpods()
}

func (a *airpodService) GetAirpodByUserID(id string) ([]*domain.Airpod, error) {
	if id == "" {
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

func (a *airpodService) DeleteAirpod(id string) error {
	return a.repo.DeleteAirpod(id)
}
