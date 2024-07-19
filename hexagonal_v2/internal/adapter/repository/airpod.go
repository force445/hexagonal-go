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
	Db.AutoMigrate(&domain.Airpod{})
	return &airpodRepository{Db: Db}
}
func (a *airpodRepository) CreateAirpod(airpod *domain.Airpod) error {
	return a.Db.Create(airpod).Error
}

func (a *airpodRepository) GetAirpodByID(id int64) (*domain.Airpod, error) {
	var airpod domain.Airpod
	if err := a.Db.Where("id = ?", id).First(&airpod).Error; err != nil {
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

func (a *airpodRepository) GetAirpodByUserID(id int64) ([]*domain.Airpod, error) {
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

func (a *airpodRepository) DeleteAirpod(id int64) error {
	return a.Db.Delete(&domain.Airpod{}, id).Error
}
