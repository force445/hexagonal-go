package port

import "hexagonal_v2/internal/core/domain"

type AirpodService interface {
	CreateAirpod(airpod *domain.Airpod) error
	GetAirpodByID(id string) (*domain.Airpod, error)
	GetAirpods() ([]*domain.Airpod, error)
	GetAirpodByUserID(id string) ([]*domain.Airpod, error)
	UpdateAirpod(airpod *domain.Airpod) error
	DeleteAirpod(id string) error
}

type UserService interface {
	CreateUser(user *domain.User) error
	GetUserByID(id string) (*domain.User, error)
	GetUsers() ([]*domain.User, error)
	UpdateUser(user *domain.User) error
	DeleteUser(id string) error
}

type LocationService interface {
	CreateLocation(location *domain.Location) error
	GetLocationByID(id string) (*domain.Location, error)
	GetLocations() ([]*domain.Location, error)
	UpdateLocation(location *domain.Location) error
	DeleteLocation(id string) error
}

type AirpodRepository interface {
	CreateAirpod(airpod *domain.Airpod) error
	GetAirpodByID(id string) (*domain.Airpod, error)
	GetAirpods() ([]*domain.Airpod, error)
	GetAirpodByUserID(id string) ([]*domain.Airpod, error)
	UpdateAirpod(airpod *domain.Airpod) error
	DeleteAirpod(id string) error
}

type UserRepository interface {
	CreateUser(user *domain.User) error
	GetUserByID(id string) (*domain.User, error)
	GetUsers() ([]*domain.User, error)
	UpdateUser(user *domain.User) error
	DeleteUser(id string) error
}

type LocationRepository interface {
	CreateLocation(location *domain.Location) error
	GetLocationByID(id string) (*domain.Location, error)
	GetLocations() ([]*domain.Location, error)
	UpdateLocation(location *domain.Location) error
	DeleteLocation(id string) error
}
