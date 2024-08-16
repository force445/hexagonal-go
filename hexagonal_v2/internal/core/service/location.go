package service

import (
	"hexagonal_v2/internal/core/domain"
	"hexagonal_v2/internal/core/port"
)

type locationService struct {
	location port.LocationService
}

func NewLocationService(location port.LocationService) port.LocationService {
	return &locationService{location: location}
}

func (l *locationService) CreateLocation(location *domain.Location) error {
	return l.location.CreateLocation(location)
}

func (l *locationService) GetLocationByID(id int64) (*domain.Location, error) {
	return l.location.GetLocationByID(id)
}

func (l *locationService) GetLocations() ([]*domain.Location, error) {
	return l.location.GetLocations()
}

func (l *locationService) UpdateLocation(location *domain.Location) error {
	return l.location.UpdateLocation(location)
}

func (l *locationService) DeleteLocation(id int64) error {
	return l.location.DeleteLocation(id)
}
