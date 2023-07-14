package loanapplication

import (
	"demo-ddd-clean-architecture/app/model"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type loanApp struct {
	db *gorm.DB
	tx *gorm.DB
}

// NewLoanApplication
func NewLoanApplication() *loanApp {
	return &loanApp{}
}

// WithDbConn
func (m *loanApp) WithDbConn(db *gorm.DB) *loanApp {
	m.db = db
	return m
}

// WithTx
func (m *loanApp) WithTx(tx *gorm.DB) *loanApp {
	m.tx = tx
	return m
}

// IsApprovedExists
func (m *loanApp) IsApprovedExists(custId uuid.UUID) (bool, error) {
	var count int
	if err := m.db.Select("COUNT(id) AS count").Model(&model.LoanApplication{}).Where(&model.LoanApplication{
		CustomerId: custId,
		Base: model.Base{
			Status: stsApproved,
		},
	}).Take(&count).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return false, err
	}
	if count > 0 {
		return true, nil
	}

	return false, nil
}

// GetApprovedFirst
func (m *loanApp) GetApprovedFirst(custId uuid.UUID) (*model.LoanApplication, error) {
	res := model.LoanApplication{}
	if err := m.db.Where("customer_id=? AND status=?", custId, stsApproved).Order("created_at DESC").First(&res).Error; err != nil {
		return nil, err
	}

	return &res, nil
}

// GetApprovedLoans
func (m *loanApp) GetApprovedLoans() (*[]model.LoanApplication, error) {
	res := []model.LoanApplication{}
	if err := m.db.Preload("Customer").Where("status=?", stsApproved).Order("created_at").Find(&res).Error; err != nil {
		return nil, err
	}

	return &res, nil
}

// CalculateTenor
func (m *loanApp) CalculateTenor(loans *model.LoanApplication) (int, error) {
	// Disini seharusnya di kalkulasi tenor berdasarkan limit pengajuan peminjaman
	// ## TODO

	if loans == nil {
		return 0, errors.New(errEmptyLoanApplication)
	}

	if loans.Limit == 500000 {
		return 3, nil

	} else if loans.Limit == 2000000 {
		return 4, nil
	}

	return 0, errors.New(errEmptyLoanApplication)
}

// UpdateToStatusDone
func (m *loanApp) UpdateToStatusDone(IDs *[]uuid.UUID) error {
	for _, id := range *IDs {
		if err := m.tx.Model(&model.LoanApplication{}).Where("id = ?", id).Update("status", stsDone).Error; err != nil {
			return err
		}
	}
	return nil
}
