package migration

import "demo-ddd-clean-architecture/app/model"

var (
	customer        model.Customer
	loanApplication model.LoanApplication
)

// DataSeeds
func DataSeeds() []interface{} {
	return []interface{}{
		customer.Seed(),
		loanApplication.Seed(),
	}
}
