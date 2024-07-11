package domain

type BoardSettings struct {
	Size  uint `json:"size"`
	Bombs uint `json:"bombs"`
}

type Board [][]string

func NewBoard(size uint, bombs uint) Board {
	board := make(Board, size)
	for i := range board {
		board[i] = make([]string, size)
	}

	return board
}
