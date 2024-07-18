package repository

import (
	"hexagonal_v2/internal/core/domain"
	"hexagonal_v2/internal/core/port"

	"gorm.io/gorm"
)

type locationRepository struct {
	Db *gorm.DB
}

func NewLocationRepository(Db *gorm.DB) port.LocationRepository {
	return &locationRepository{Db: Db}
}

func (l *locationRepository) CreateLocation(location *domain.Location) error {
	return l.Db.Create(location).Error
}

func (l *locationRepository) GetLocationByID(id string) (*domain.Location, error) {
	var location domain.Location
	err := l.Db.First(&location, id).Error
	if err != nil {
		return nil, err
	}
	return &location, nil
}

func (l *locationRepository) GetLocations() ([]*domain.Location, error) {
	var locations []*domain.Location
	err := l.Db.Find(&locations).Error
	if err != nil {
		return nil, err
	}
	return locations, nil
}

func (l *locationRepository) UpdateLocation(location *domain.Location) error {
	return l.Db.Save(location).Error
}

func (l *locationRepository) DeleteLocation(id string) error {
	return l.Db.Delete(&domain.Location{}, id).Error
}
