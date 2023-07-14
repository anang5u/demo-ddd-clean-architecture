package identityrules

import (
	"demo-ddd-clean-architecture/app/helper"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type identityRules struct {
	db *gorm.DB
}

// NewIdentityRules
func NewIdentityRules() *identityRules {
	return &identityRules{}
}

// WithDbConn
func (m *identityRules) WithDbConn(db *gorm.DB) *identityRules {
	m.db = db
	return m
}

// GetContractNumberCharacterLength
func (m *identityRules) GetContractNumberCharacterLength() int {
	return contractNumberCharacterLength
}

// GenerateContractNumber
func (m *identityRules) GenerateContractNumber(length ...int) (string, error) {
	var characterLength int

	if len(length) > 0 {
		characterLength = length[0]
	} else {
		characterLength = contractNumberCharacterLength
	}

	// Format YYMM+randomNumber
	return fmt.Sprintf("%s-%s", time.Now().Format("0601"), helper.RndNumeric((characterLength - 4))), nil
}
