package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	Id        uuid.UUID `json:"id,omitempty" gorm:"primaryKey;type:varchar(36)"`
	Sort      int64     `json:"sort,omitempty" gorm:"default:0"`
	Status    int       `json:"status,omitempty" gorm:"type:smallint;default:1"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
