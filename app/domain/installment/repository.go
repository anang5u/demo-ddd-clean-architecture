package installment

import (
	"demo-ddd-clean-architecture/app/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InstallmentRepository interface {
	WithDbConn(db *gorm.DB) *modInstallment
	WithTx(tx *gorm.DB) *modInstallment

	IsActiveInstallmentExists(customers *model.Customer) (bool, error)
	Create(loans *[]model.Installment) error
	UpdateStatusPayTo(installmentId uuid.UUID, iStsPay int64) error
	GetInfo() (*[]model.Installment, error)
	GetInstallment(contractNumber string) (*model.Installment, error)
	GetRemainingTenor(contractNumber string) (int64, error)
	TakeById(installmentId *uuid.UUID) (*model.Installment, error)

	// Statuses
	GetStatusActive() int64

	// status payment
	GetStatusPayUnpaid() int64
	GetStsPayNeedConfirm() int64
	GetStatusPayPending() int64
	GetStatusPayPaid() int64
	GetStatusPayFailed() int64
	GetMapStsPay(iStsPay int64) string
}
