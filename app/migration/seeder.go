package migration

import "demo-ddd-clean-architecture/app/model"

var (
	customer model.Customer
)

// DataSeeds
func DataSeeds() []interface{} {
	return []interface{}{
		customer.Seed(),
	}
}
