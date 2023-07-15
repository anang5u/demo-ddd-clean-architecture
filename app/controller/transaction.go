package controller

import (
	"demo-ddd-clean-architecture/app/model"
	"demo-ddd-clean-architecture/app/services/transaction"
	"errors"

	"github.com/gofiber/fiber/v2"
)

var trans = transaction.Apply()

// Info Tagihan
func Info(c *fiber.Ctx) error {
	res, err := trans.GetInstallmentInfo()
	if err != nil {
		return SendError(c, err)
	}

	return SendSuccess(c, res)
}

// Cek Tagihan
func Inquiry(c *fiber.Ctx) error {
	contractNumber := c.Params("contract_number")

	res, err := trans.Inquiry(contractNumber)
	if err != nil {
		return SendError(c, err)
	}

	return SendSuccess(c, res)
}

// Payment
func Payment(c *fiber.Ctx) error {
	req := model.PaymentRequest{}

	if err := c.BodyParser(&req); err != nil {
		trans.Common.Log.Error("Error body parse: ", err.Error())
		return SendError(c, errors.New("Error lainnya"))
	}

	res, err := trans.Payment(&req)
	if err != nil {
		return SendError(c, err)
	}

	return SendSuccess(c, res)
}

// Detail Pembayaran
func PaymentDetail(c *fiber.Ctx) error {
	installmentId := c.Params("installment_id")

	res, err := trans.PaymentDetail(installmentId)
	if err != nil {
		return SendError(c, err)
	}

	return SendSuccess(c, res)
}

// Payment Confirm with token
func Token(c *fiber.Ctx) error {
	req := model.PaymentTokenRequest{}

	if err := c.BodyParser(&req); err != nil {
		trans.Common.Log.Error("Error body parse POST Token: ", err.Error())
		return SendError(c, errors.New("Error lainnya"))
	}

	res, err := trans.Token(&req)
	if err != nil {
		return SendError(c, err)
	}

	return SendSuccess(c, res)
}
