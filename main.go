package main

import (
	"demo-ddd-clean-architecture/app/services/common"
	"demo-ddd-clean-architecture/app/services/loans"
)

func main() {
	cmn := common.Apply(".env")
	cmn.Db.AutoMigrate()

	loanSevice := loans.Apply(cmn.Db.GetDbConn())
	loanSevice.GenerateInstallment()
}
