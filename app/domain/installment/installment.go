package installment

import (
	"demo-ddd-clean-architecture/app/model"
	"errors"

	"gorm.io/gorm"
)

type modInstallment struct {
	db *gorm.DB
	tx *gorm.DB
}

// NewInstallment
func NewInstallment() *modInstallment {
	return &modInstallment{}
}

// WithDbConn
func (m *modInstallment) WithDbConn(db *gorm.DB) *modInstallment {
	m.db = db
	return m
}

// WithTx
func (m *modInstallment) WithTx(tx *gorm.DB) *modInstallment {
	m.tx = tx
	return m
}

// IsNewInstallmentExists
func (m *modInstallment) IsNewInstallmentExists(customers *model.Customer) (bool, error) {
	var count int
	if err := m.db.Select("COUNT(id) AS count").Model(&model.Installment{}).Where(&model.Installment{
		CustomerId: customers.Id,
		Base: model.Base{
			Status: stsNewInstallment,
		},
	}).Take(&count).Error; err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}

// Create
func (m *modInstallment) Create(loans *[]model.Installment) error {
	if m.tx == nil {
		return errors.New("Tx while Create installment is required!")
	}

	if err := m.tx.Model(&model.Installment{}).Create(loans).Error; err != nil {
		return err
	}

	return nil
}
