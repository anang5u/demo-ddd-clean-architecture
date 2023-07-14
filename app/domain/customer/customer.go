package customer

import (
	"gorm.io/gorm"
)

type modCustomer struct {
	db *gorm.DB
}

// NewCustomer
func NewCustomer(db *gorm.DB) *modCustomer {
	return &modCustomer{
		db: db,
	}
}
