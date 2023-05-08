package database

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID       uint `gorm:"primaryKey"`
	Quantity uint
	Date     time.Time
}

func (d *Order) BeforeCreate(tx *gorm.DB) error {
	d.Date = time.Now()
	return nil
}
