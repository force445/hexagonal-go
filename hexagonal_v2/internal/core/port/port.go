package port

import "hexagonal_v2/internal/core/domain"

type AirpodService interface {
	CreateAirpod(airpod *domain.Airpod) error
	GetAirpodByID(id int64) (*domain.Airpod, error)
	GetAirpods() ([]*domain.Airpod, error)
	GetAirpodByUserID(id int64) ([]*domain.Airpod, error)
	UpdateAirpod(airpod *domain.Airpod) error
	DeleteAirpod(id int64) error
}

type UserService interface {
	CreateUser(user *domain.User) error
	GetUserByID(id int64) (*domain.User, error)
	GetUsers() ([]*domain.User, error)
	UpdateUser(user *domain.User) error
	DeleteUser(id int64) error
}

type LocationService interface {
	CreateLocation(location *domain.Location) error
	GetLocationByID(id int64) (*domain.Location, error)
	GetLocations() ([]*domain.Location, error)
	UpdateLocation(location *domain.Location) error
	DeleteLocation(id int64) error
}

type AirpodRepository interface {
	CreateAirpod(airpod *domain.Airpod) error
	GetAirpodByID(id int64) (*domain.Airpod, error)
	GetAirpods() ([]*domain.Airpod, error)
	GetAirpodByUserID(id int64) ([]*domain.Airpod, error)
	UpdateAirpod(airpod *domain.Airpod) error
	DeleteAirpod(id int64) error
}

type UserRepository interface {
	CreateUser(user *domain.User) error
	GetUserByID(id int64) (*domain.User, error)
	GetUsers() ([]*domain.User, error)
	UpdateUser(user *domain.User) error
	DeleteUser(id int64) error
}

type LocationRepository interface {
	CreateLocation(location *domain.Location) error
	GetLocationByID(id int64) (*domain.Location, error)
	GetLocations() ([]*domain.Location, error)
	UpdateLocation(location *domain.Location) error
	DeleteLocation(id int64) error
}
