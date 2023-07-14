package installment

import (
	"demo-ddd-clean-architecture/app/model"

	"gorm.io/gorm"
)

type InstallmentRepository interface {
	WithDbConn(db *gorm.DB) *modInstallment
	WithTx(tx *gorm.DB) *modInstallment

	IsNewInstallmentExists(customers *model.Customer) (bool, error)
	Create(loans *[]model.Installment) error

	// Statuses
	GetStatusNew() int
}
