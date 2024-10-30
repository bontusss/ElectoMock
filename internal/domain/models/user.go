package models

import "time"

type User struct {
	ID         uint   `gorm:"primaryKey"`
	Email      string `gorm:"uniqueIndex;not null"`
	Password   string
	Name       string
	Provider   string
	ProviderID string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
