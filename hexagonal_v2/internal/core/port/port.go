package port

import "hexagonal_v2/internal/core/domain"

type AirpodService interface {
	CreateAirpod(airpod *domain.Airpod) error
	GetAirpodByID(id int) (*domain.Airpod, error)
	GetAirpods() ([]*domain.Airpod, error)
	UpdateAirpod(airpod *domain.Airpod) error
	DeleteAirpod(id int) error
}

type UserService interface {
	CreateUser(user *domain.User) error
	GetUserByID(id int) (*domain.User, error)
	GetUsers() ([]*domain.User, error)
	UpdateUser(user *domain.User) error
	DeleteUser(id int) error
}

type LocationService interface {
	CreateLocation(location *domain.Location) error
	GetLocationByID(id int) (*domain.Location, error)
	GetLocations() ([]*domain.Location, error)
	UpdateLocation(location *domain.Location) error
	DeleteLocation(id int) error
}
