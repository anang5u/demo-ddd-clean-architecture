package transaction

import (
	"demo-ddd-clean-architecture/app/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TrxRepository interface {
	WithDbConn(db *gorm.DB) *modTrx
	WithTx(tx *gorm.DB) *modTrx
	Create(trans *[]model.Transaction) error
	UpdateToStatusPay(transId *uuid.UUID, iStatus *int64, longDesc ...*string) error
	Delete(transId *uuid.UUID, longDesc ...*string) error

	// Messaging
	GetMessage(key string) string
}
