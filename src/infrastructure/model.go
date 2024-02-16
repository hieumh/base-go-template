package infrastructure

import "github.com/google/uuid"

type User struct {
	Id    uuid.UUID `gorm:"primaryKey"`
	Name  string
	Price float64
	Email string
}
