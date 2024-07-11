package domain

type Game struct {
	ID            string        `json:"id"`
	Name          string        `json:"name"`
	State         string        `json:"state"`
	BoardSettings BoardSettings `json:"boardSettings"`
	Board         Board         `json:"board"`
}
