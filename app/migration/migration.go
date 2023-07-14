package migration

import "demo-ddd-clean-architecture/app/model"

// ModelMigrations
var ModelMigrations []interface{} = []interface{}{
	model.Customer{},
	model.LoanApplication{},
}
