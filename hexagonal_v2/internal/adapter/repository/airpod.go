package repository

import (
	"hexagonal_v2/internal/core/domain"
	"hexagonal_v2/internal/core/port"

	"gorm.io/gorm"
)

type airpodRepository struct {
	Db *gorm.DB
}

func NewAirpodRepository(Db *gorm.DB) port.AirpodRepository {
	return &airpodRepository{Db: Db}
}

func (a *airpodRepository) CreateAirpod(airpod *domain.Airpod) error {
	return a.Db.Create(airpod).Error
}

func (a *airpodRepository) GetAirpodByID(id string) (*domain.Airpod, error) {
	var airpod domain.Airpod
	err := a.Db.First(&airpod, id).Error
	if err != nil {
		return nil, err
	}
	return &airpod, nil
}

func (a *airpodRepository) GetAirpods() ([]*domain.Airpod, error) {
	var airpods []*domain.Airpod
	err := a.Db.Find(&airpods).Error
	if err != nil {
		return nil, err
	}
	return airpods, nil
}

func (a *airpodRepository) GetAirpodByUserID(id string) ([]*domain.Airpod, error) {
	var airpods []*domain.Airpod
	err := a.Db.Where("user_id = ?", id).Find(&airpods).Error
	if err != nil {
		return nil, err
	}
	return airpods, nil
}

func (a *airpodRepository) UpdateAirpod(airpod *domain.Airpod) error {
	return a.Db.Save(airpod).Error
}

func (a *airpodRepository) DeleteAirpod(id string) error {
	return a.Db.Delete(&domain.Airpod{}, id).Error
}
