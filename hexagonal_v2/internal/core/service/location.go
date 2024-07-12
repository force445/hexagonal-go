package service

import (
	"hexagonal_v2/internal/core/domain"
	"hexagonal_v2/internal/core/port"
)

type LocationService struct {
	location port.LocationService
}

func NewLocationService(location port.LocationService) *LocationService {
	return &LocationService{location: location}
}

func (l *LocationService) CreateLocation(location *domain.Location) error {
	return l.location.CreateLocation(location)
}

func (l *LocationService) GetLocationByID(id int) (*domain.Location, error) {
	return l.location.GetLocationByID(id)
}

func (l *LocationService) GetLocations() ([]*domain.Location, error) {
	return l.location.GetLocations()
}

func (l *LocationService) UpdateLocation(location *domain.Location) error {
	return l.location.UpdateLocation(location)
}

func (l *LocationService) DeleteLocation(id int) error {
	return l.location.DeleteLocation(id)
}
