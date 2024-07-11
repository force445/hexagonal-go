package ports

import "hexagonal/internal/core/domain"

type GameRepository interface {
	Get(id string) (domain.Game, error)
	Save(game domain.Game) error
}
