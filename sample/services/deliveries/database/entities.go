package database

import (
	"time"

	"gorm.io/gorm"
)

type Delivery struct {
	ID      uint `gorm:"primaryKey"`
	Address string
	OrderID uint
	Date    time.Time
}

func (d *Delivery) BeforeCreate(tx *gorm.DB) error {
	d.Date = time.Now()
	return nil
}
