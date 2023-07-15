package transaction

import (
	"demo-ddd-clean-architecture/app/model"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type modTrx struct {
	db *gorm.DB
	tx *gorm.DB
}

// NewTransaction
func NewTransaction(db *gorm.DB) *modTrx {
	return &modTrx{
		db: db,
	}
}

// WithDbConn
func (m *modTrx) WithDbConn(db *gorm.DB) *modTrx {
	m.db = db
	return m
}

// WithTx
func (m *modTrx) WithTx(tx *gorm.DB) *modTrx {
	m.tx = tx
	return m
}

// Create
func (m *modTrx) Create(trans *[]model.Transaction) error {
	if m.tx == nil {
		return errors.New("Tx while Create transaction is required!")
	}

	if err := m.tx.Model(&model.Transaction{}).Create(trans).Error; err != nil {
		return err
	}

	return nil
}

// UpdateToStatusPay
func (m *modTrx) UpdateToStatusPay(transId *uuid.UUID, iStatus *int64, desc ...*string) error {
	if m.tx == nil {
		return errors.New("Tx while UpdateToStatusPay transaction is required!")
	}

	var longDesc *string
	if len(desc) > 0 {
		longDesc = desc[0]
	}

	if err := m.tx.Model(&model.Transaction{}).Where("id=?", transId).Updates(&model.Transaction{
		PaymentStatus:   iStatus,
		LongDescription: longDesc,
	}).Error; err != nil {
		return err
	}

	return nil
}

// Delete
func (m *modTrx) Delete(transId *uuid.UUID, desc ...*string) error {
	if m.tx == nil {
		return errors.New("Tx while UpdateToStatusPay transaction is required!")
	}

	var longDesc *string
	if len(desc) > 0 {
		longDesc = desc[0]
	}

	if err := m.tx.Where("id=?", transId).Updates(&model.Transaction{
		LongDescription: longDesc,
	}).Error; err != nil {
		return err
	}

	// Soft deleted
	if err := m.tx.Where("id=?", transId).Delete(&model.Transaction{}).Error; err != nil {
		return err
	}

	return nil
}
