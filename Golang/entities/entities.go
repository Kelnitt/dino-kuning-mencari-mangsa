package entities

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type SampleTabler struct {
	// Sample Tabler Model
	ID        uint           `gorm:"primaryKey;autoIncrement;not null;column:sample_id" json:"sample_id"`
	Sample    string         `gorm:"type:varchar(255);not null;column:sample_name" json:"sample"`
	Quantity  int            `gorm:"not null;default:0;column:quantity" json:"quantity"`
	Price     int            `gorm:"not null;default:0;column:price" json:"price"`
	CreatedAt time.Time      `gorm:"autoCreateTime;not null;column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime;not null;column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (s *SampleTabler) BeforeSave(tx *gorm.DB) (err error) {
	// Maintain Data Quality
	if s.Quantity < 0 {
		log.Printf("Warning : Minimum Quantity is 0, But Your Quantity is %d", s.Quantity)
	}
	if s.Price < 0 {
		log.Printf("Warning : Minimum Price is 0, But Your Price is %d", s.Price)
	}
	return nil
}