package main

import (
	"demo-ddd-clean-architecture/app/services/common"
)

func main() {
	common.Apply().Db.AutoMigrate()
}
