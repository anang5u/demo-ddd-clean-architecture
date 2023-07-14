package pricing

import (
	"gorm.io/gorm"
)

type modPricing struct {
	db *gorm.DB
}

// NewPricing
func NewPricing(db *gorm.DB) *modPricing {
	return &modPricing{
		db: db,
	}
}

// GetOptionAdminFee
func (m *modPricing) GetOptionAdminFee() float64 {
	//bisa diambil dri salah satu tabel sesuai requrement bisnis
	return defaultAdmFee
}

// GetOptionInterestAmt
func (m *modPricing) GetOptionInterestAmt() float64 {
	//bisa diambil dri salah satu tabel sesuai requrement bisnis
	return defaultInterestAmt
}
