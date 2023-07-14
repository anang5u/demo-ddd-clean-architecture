package installment

import (
	"demo-ddd-clean-architecture/app/model"
	"errors"
	"log"

	"gorm.io/gorm"
)

type modInstallment struct {
	db *gorm.DB
}

// NewInstallment
func NewInstallment() *modInstallment {
	return &modInstallment{}
}

// WitdDbConn
func (m *modInstallment) WithDbConn(db *gorm.DB) *modInstallment {
	m.db = db
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
	}).Take(&count).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}

// GetLimit
/*
func (m *modInstallment) GetLimit(customerId uuid.UUID) float64 {

}
*/

// Generate
func (m *modInstallment) Generate(customers *[]model.Customer) error {
	m.db.Transaction(func(tx *gorm.DB) error {
		for _, cust := range *customers {
			log.Println(cust)
			/*
				// do some database operations in the transaction (use 'tx' from this point, not 'db')
				if err := tx.Create(&Animal{Name: "Giraffe"}).Error; err != nil {
					// return any error will rollback
					return err
				}

				if err := tx.Create(&Animal{Name: "Lion"}).Error; err != nil {
					return err
				}
			*/
		}
		// return nil will commit the whole transaction
		return nil
	})

	return nil
}
