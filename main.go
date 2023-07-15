package main

import (
	"demo-ddd-clean-architecture/app/routes"
	"demo-ddd-clean-architecture/app/services/common"
	"demo-ddd-clean-architecture/app/services/loans"
	"log"

	"github.com/gofiber/fiber/v2"
)

var cmn *common.CommonService

func init() {
	cmn = common.Apply(".env")
	cmn.Db.AutoMigrate()

	/*
	* Auto Generate Installment - Cicilan pambayaran
	* Hanya untuk demo cek cicilan/tagihan dan pembayaran
	* Di env production (real world) mungkin di generate pada saat pengajuan di setujui/approved
	* atau sesuai dg busines requirement tertentu :-)
	 */
	loanSevice := loans.Apply(cmn.Db.GetDbConn())
	loanSevice.GenerateInstallment()

}

func main() {
	app := fiber.New()

	routes.Handle(app)
	if err := app.Listen(":" + cmn.Config.Get("APP_PORT")); err != nil {
		log.Panic(err)
	}
}
