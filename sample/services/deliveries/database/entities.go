package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Deliveries struct {
	Id      string `gorm:"type:varchar(40);not null"`
	Address string
	OrderId string `gorm:"type:varchar(40);not null"`
	Date    time.Time
}

func (d *Deliveries) BeforeCreate(tx *gorm.DB) error {
	d.Id = uuid.New().String()
	d.Date = time.Now()
	return nil
}
