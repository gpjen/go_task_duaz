package products

import "time"

type Product struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"not null;type:varchar(100)"`
	Price       int    `json:"price" gorm:"not null"`
	Stock       int    `json:"stock" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`

	CreatedAt time.Time `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
}
