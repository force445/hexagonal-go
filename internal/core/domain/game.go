package domain

type Game struct {
	ID            string        `json:"id"`
	Name          string        `json:"name"`
	State         string        `json:"state"`
	BoardSettings BoardSettings `json:"boardSettings"`
	Board         Board         `json:"board"`
}

func NewGame(id string, name string, size uint, bombs uint) Game {
	return Game{
		ID:            id,
		Name:          name,
		State:         "playing",
		BoardSettings: BoardSettings{Size: size, Bombs: bombs},
		Board:         NewBoard(size, bombs),
	}
}
