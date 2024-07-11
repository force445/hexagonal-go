package services

import (
	"errors"
	"hexagonal/internal/core/domain"
	"hexagonal/internal/core/ports"
)

type service struct {
	gameRepository ports.GameRepository
}

func New(gameRepository ports.GameRepository) *service {
	return &service{
		gameRepository: gameRepository,
	}
}

func (srv *service) Get(id string) (domain.Game, error) {
	game, err := srv.gameRepository.Get(id)
	if err != nil {
		return domain.Game{}, errors.New("get game from repository has failed")
	}

	return game, nil
}
