package users

import "time"

type User struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Name         string `json:"name" gorm:"not null;type:varchar(100)"`
	Email        string `json:"email" gorm:"not null;type:varchar(100)"`
	Address      string `json:"address" gorm:"not null;type:varchar(200)"`
	HashPassword string `json:"-" gorm:"not null"`
	Role         string `json:"role" gorm:"type:varchar(10);default:'user'"`

	CreatedAt time.Time `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
}
