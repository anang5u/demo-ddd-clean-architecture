package installment

import (
	"demo-ddd-clean-architecture/app/model"
	"errors"

	"github.com/google/uuid"
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

// IsActiveInstallmentExists
func (m *modInstallment) IsActiveInstallmentExists(customers *model.Customer) (bool, error) {
	var count int
	if err := m.db.Select("COUNT(id) AS count").Model(&model.Installment{}).Where(&model.Installment{
		CustomerId: customers.Id,
		Base: model.Base{
			Status: stsActiveInstallment,
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

// UpdateStatusPayTo
func (m *modInstallment) UpdateStatusPayTo(installmentId uuid.UUID, iStsPay int64) error {
	if m.tx == nil {
		return errors.New("Tx while UpdateStatusPayTo installment, tx is required!")
	}

	if err := m.tx.Model(&model.Installment{}).Where("id=?", installmentId).Update("payment_status", iStsPay).Error; err != nil {
		return err
	}

	return nil
}

// GetInfo
func (m *modInstallment) GetInfo() (*[]model.Installment, error) {
	result := []model.Installment{}
	if err := m.db.Model(&model.Installment{}).Select("contract_number, customer_id").
		Group("contract_number, customer_id").
		Preload("Customer").
		Where("payment_status=?", stsPayUnpaid).Find(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

// GetInstallment
func (m *modInstallment) GetInstallment(contractNumber string) (*model.Installment, error) {
	result := model.Installment{}
	if err := m.db.Model(&model.Installment{}).
		Preload("Customer").
		Order("sort").
		Where("contract_number=? AND payment_status=?", contractNumber, stsPayUnpaid).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

// GetRemainingTenor
func (m *modInstallment) GetRemainingTenor(contractNumber string) (int64, error) {
	var count int64
	if err := m.db.Model(&model.Installment{}).Select("COUNT(id) AS count").
		Where("contract_number=? AND payment_status=?", contractNumber, stsPayUnpaid).First(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// TakeById
func (m *modInstallment) TakeById(installmentId *uuid.UUID) (*model.Installment, error) {
	result := model.Installment{}
	if err := m.db.Model(&model.Installment{}).
		Preload("Customer").
		Preload("Transaction").
		Where("id=?", installmentId).Take(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}
