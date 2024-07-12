package domain

type User struct {
	ID       int       `json:"id"`
	Username string    `json:"username"`
	Airpod   []*Airpod `json:"airpod"`
}

type Airpod struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Model      string    `json:"model"`
	UserID     int       `json:"user_id"`
	User       *User     `json:"user"`
	LocationID int       `json:"location_id"`
	Location   *Location `json:"location"`
}

type Location struct {
	ID int `json:"id"`
	X  int `json:"x"`
	Y  int `json:"y"`
}
