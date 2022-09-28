package models

import "time"

type User struct {
	Email        string    `gorm:"type:varchar(50);not null;primary_key;unique;"`
	Name         string    `gorm:"type:varchar(20);not null;"`
	Password     string    `gorm:"type:varchar(20);not null"`
	Avatar       string    `gorm:"type:varchar(128);not null"`
	RegisterTime time.Time `gorm:"type:DATETIME"`
}
