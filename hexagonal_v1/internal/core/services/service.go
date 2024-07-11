package services

import (
	"errors"
	"hexagonal/internal/core/domain"
	"hexagonal/internal/core/ports"
	"hexagonal/pkg/uidgen"
)

type service struct {
	gameRepository ports.GameRepository
	uidGen         uidgen.UIDGen
}

func New(gameRepository ports.GameRepository, uidGen uidgen.UIDGen) *service {
	return &service{
		gameRepository: gameRepository,
		uidGen:         uidGen,
	}
}

func (srv *service) Get(id string) (domain.Game, error) {
	game, err := srv.gameRepository.Get(id)
	if err != nil {
		return domain.Game{}, errors.New("get game from repository has failed")
	}

	return game, nil
}

// func (srv *service) Create(name string, size uint, bombs uint) (domain.Game, error) {
// 	if bombs >= size*size {
// 		return domain.Game{}, errors.New("the number of bombs is invalid")
// 	}

// 	game := domain.NewGame(srv.uidGen.New(), name, size, bombs)

// 	if err := srv.gameRepository.Save(game); err != nil {
// 		return domain.Game{}, errors.New("create game into repository has failed")
// 	}

// 	return game, nil
// }
