package service

import (
	"hexagonal_v2/internal/core/domain"
	"hexagonal_v2/internal/core/port"

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

func (a *airpodService) UpdateAirpod(airpod *domain.Airpod) error {
	return a.repo.UpdateAirpod(airpod)
}

func (a *airpodService) DeleteAirpod(id string) error {
	return a.repo.DeleteAirpod(id)
}
