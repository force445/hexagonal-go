package service

import (
	"hexagonal_v2/internal/core/domain"
	"hexagonal_v2/internal/core/port"
)

type AirpodService struct {
	airpod port.AirpodService
}

func NewAirpodService(airpod port.AirpodService) *AirpodService {
	return &AirpodService{airpod: airpod}
}

func (a *AirpodService) CreateAirpod(airpod *domain.Airpod) error {
	return a.airpod.CreateAirpod(airpod)
}

func (a *AirpodService) GetAirpodByID(id int) (*domain.Airpod, error) {
	return a.airpod.GetAirpodByID(id)
}

func (a *AirpodService) GetAirpods() ([]*domain.Airpod, error) {
	return a.airpod.GetAirpods()
}

func (a *AirpodService) UpdateAirpod(airpod *domain.Airpod) error {
	return a.airpod.UpdateAirpod(airpod)
}

func (a *AirpodService) DeleteAirpod(id int) error {
	return a.airpod.DeleteAirpod(id)
}
