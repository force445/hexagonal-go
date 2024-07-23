package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int64          `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeleteAt  gorm.DeletedAt `gorm:"index" json:"-"`
	Username  string         `json:"username"`
	Airpods   []Airpod       `gorm:"foreignKey:UserID" json:"airpods"`
}

type Airpod struct {
	ID         int64          `gorm:"primaryKey" json:"id"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeleteAt   gorm.DeletedAt `gorm:"index" json:"-"`
	ModelName  string         `json:"model"`
	UserID     int64          `json:"user_id"`
	LocationID int            `json:"location_id"`
	Location   *Location      `gorm:"foreignKey:LocationID" json:"location"`
}

type Location struct {
	gorm.Model
	ID int64 `gorm:"primaryKey" json:"id"`
	X  int   `json:"x"`
	Y  int   `json:"y"`
}
