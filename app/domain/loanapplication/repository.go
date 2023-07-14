package loanapplication

import (
	"demo-ddd-clean-architecture/app/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LoanApplicationRepository interface {
	WithDbConn(db *gorm.DB) *loanApp
	WithTx(tx *gorm.DB) *loanApp

	IsApprovedExists(custId uuid.UUID) (bool, error)
	GetApprovedFirst(custId uuid.UUID) (*model.LoanApplication, error)
	GetApprovedLoans() (*[]model.LoanApplication, error)
	CalculateTenor(loans *model.LoanApplication) (int, error)
	UpdateToStatusDone(IDs *[]uuid.UUID) error
}
