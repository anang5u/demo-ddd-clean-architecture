package transaction

import (
	"demo-ddd-clean-architecture/app/helper"
	"demo-ddd-clean-architecture/app/model"
	"errors"
	"fmt"
)

// GetInstallmentInfo
func (s *TrxService) GetInstallmentInfo() (*[]model.InstallmentInfoResponse, error) {
	res, err := s.Installment.GetInfo()
	if err != nil {
		s.Common.Log.Error("Error while GetInstallmentInfo: ", err.Error())
		return nil, errors.New("Tidak ada info konsumen terdaftar!")
	}

	data := []model.InstallmentInfoResponse{}
	for _, info := range *res {
		maskedName := s.IdentityRule.CreateMask(info.Customer.FullName)
		contractNumber := info.ContractNumber

		data = append(data, model.InstallmentInfoResponse{
			CustomerName:   &maskedName,
			ContractNumber: &contractNumber,
		})
	}

	return &data, nil
}

// Inquiry
func (s *TrxService) Inquiry(contractNumber string) (*model.TrxResponse, error) {
	res, err := s.Installment.GetInstallment(contractNumber)
	if err != nil {
		s.Common.Log.Error("Error while GetInstallment: ", err.Error())
		return nil, errors.New("Tidak ada info tagihan/res!")
	}

	return s.generateTrxResponse(res), nil
}

// Payment
func (s *TrxService) Payment(req *model.PaymentRequest) (*model.TrxResponse, error) {
	if req == nil {
		return nil, errors.New("parameter request tidak valid!")
	}

	istallmentId := helper.UuidMustParse(req.IntallmentID)

	// prevent duplicate transaction
	s.Lock()
	defer s.Unlock()

	res, err := s.Installment.TakeById(&istallmentId)
	if err != nil {
		return nil, errors.New("tagihan/res tidak ditemukan!")
	}

	// Data transaksi sudah tersedia dan belum dilakukan pembayaran
	iStatusPay := s.Installment.GetStatusPayUnpaid()
	if res.Transaction != nil && *res.Transaction.PaymentStatus == iStatusPay {
		return s.PaymentDetail(res.Id.String())
	}

	token := s.IdentityRule.GenerateToken()
	tokenHash := s.IdentityRule.MakeHash(fmt.Sprintf("%s%s", req.IntallmentID, token))
	paymentExpiredAt := s.IdentityRule.MakeExpiredAtHours()
	trxDate := helper.GetTimeNow()

	trans := []model.Transaction{}
	trans = append(trans, model.Transaction{
		Base: model.Base{
			Id:        helper.UuidNew(),
			CreatedAt: &trxDate,
		},
		InstallmentId:    &res.Id,
		ContractNumber:   &res.ContractNumber,
		OtrAmt:           &res.OtrAmt,
		AdminFee:         &res.AdminFee,
		InstallmentAmt:   &res.InstallmentAmt,
		InterestAmt:      &res.InterestAmt,
		TotalAmt:         &res.TotalAmt,
		AssetName:        &res.AssetName,
		PaymentStatus:    helper.Int64Ptr(s.Installment.GetStsPayNeedConfirm()),
		PaymetExpiredAt:  &paymentExpiredAt,
		CustomerId:       helper.UuidPtr(res.CustomerId),
		FullName:         &res.Customer.FullName,
		LegalName:        &res.Customer.LegalName,
		IdCardNumber:     &res.Customer.IdCardNumber,
		Token:            &tokenHash,
		ShortDescription: &req.ShortDescription,
	})

	// begin transaction
	tx := s.Common.Db.TxBegin()
	defer s.Common.Db.TxRecover(tx)

	// create transaction
	if err := s.Trx.WithTx(tx).Create(&trans); err != nil {
		s.Common.Log.Error("Error while Create transaction: ", err.Error())
		s.Common.Db.TxRollback(tx)
		return nil, errors.New(s.Trx.GetMessage("error_others"))
	}

	// Update Installment to Need Confirm
	// Menunggu pengiriman TOKEN sebagai konfirmasi pembayaran
	paymentStatus := s.Installment.GetStsPayNeedConfirm()
	if err := s.Installment.WithTx(tx).UpdateStatusPayTo(res.Id, paymentStatus); err != nil {
		s.Common.Log.Error("Error while UpdateStatusPayTo on installment: ", err.Error())
		s.Common.Db.TxRollback(tx)
		return nil, errors.New(s.Trx.GetMessage("error_others"))
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		s.Common.Log.Error("Error while Commit create transaction: ", err.Error())
		return nil, errors.New(s.Trx.GetMessage("error_others"))
	}

	data := s.generateTrxResponse(res)
	data.CreatedAt = helper.StrPtr(trxDate.Format("2006-01-02 15:04:05 MST"))
	data.ShortDescription = &req.ShortDescription
	data.LongDescription = helper.StrPtr(fmt.Sprintf("Segera lakukan pembayaran sebelum %s", paymentExpiredAt.Format("2006-01-02 15:04:05 MST")))
	data.StsPay = helper.StrPtr(s.Installment.GetMapStsPay(paymentStatus))

	descToken := fmt.Sprintf("Ini adalah token %s, dan tidak seharusnya dikirim kan ke sini (hanya untuk demo!!!). Biasanya dikirim melalui SMS atau Email", token)
	data.Token = &descToken

	return data, nil
}

// Detail Pembayaran
// PaymentDetail
func (s *TrxService) PaymentDetail(sInstallmentId string) (*model.TrxResponse, error) {
	istallmentId := helper.UuidMustParse(sInstallmentId)
	res, err := s.Installment.TakeById(&istallmentId)
	if err != nil {
		return nil, errors.New("Data pembayaran tidak ditemukan!")
	}

	data := s.generateTrxResponse(res)
	data.CreatedAt = helper.StrPtr(res.Transaction.CreatedAt.Format("2006-01-02 15:04:05 MST"))
	data.ShortDescription = res.Transaction.ShortDescription
	data.AssetName = &res.AssetName

	return data, nil
}

// Token
func (s *TrxService) Token(req *model.PaymentTokenRequest) (*model.TrxResponse, error) {
	if req == nil {
		return nil, errors.New("parameter request tidak valid!")
	}

	istallmentId := helper.UuidMustParse(req.IntallmentID)

	// prevent duplicate transaction
	s.Lock()
	defer s.Unlock()

	res, err := s.Installment.TakeById(&istallmentId)
	if err != nil || res.Transaction == nil {
		return nil, errors.New("Tidak ditemukan data pembayaran. Silahkan kontak CS kami!")
	}

	// begin transaction
	tx := s.Common.Db.TxBegin()
	defer s.Common.Db.TxRecover(tx)

	// compare token
	tokenSavedPlain := fmt.Sprintf("%s%s", req.IntallmentID, req.Token)
	isOK := s.IdentityRule.CompareHash(*res.Transaction.Token, tokenSavedPlain)
	if isOK == false {
		iStsPayUnpaid := s.Installment.GetStatusPayUnpaid()
		longDescription := "Mismatch Token"

		// Rollback transaksi agar bisa dilakukan pembayaran ulang
		if err := s.Trx.WithTx(tx).Delete(&res.Transaction.Id, &longDescription); err != nil {
			return nil, errors.New(s.Trx.GetMessage("error_others"))
		}

		// Rollback installment agar bisa dilakukan pembayaran ulang
		if err := s.Installment.WithTx(tx).UpdateStatusPayTo(res.Id, iStsPayUnpaid); err != nil {
			s.Common.Db.TxRollback(tx)
			return nil, errors.New(s.Trx.GetMessage("error_others"))
		}

		// Commit transaction
		if err := tx.Commit().Error; err != nil {
			s.Common.Log.Error("Error while Commit Rollback transaction: ", err.Error())
			s.Common.Db.TxRollback(tx)
			return nil, errors.New(s.Trx.GetMessage("error_others"))
		}

		return nil, errors.New(s.Trx.GetMessage("error_mismatch_token"))
	}

	iStsPayPending := s.Installment.GetStatusPayPending()
	if err := s.Trx.WithTx(tx).UpdateToStatusPay(&res.Transaction.Id, &iStsPayPending); err != nil {
		s.Common.Db.TxRollback(tx)
		return nil, errors.New(s.Trx.GetMessage("error_others"))
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		s.Common.Log.Error("Error while Commit transaction: ", err.Error())
		s.Common.Db.TxRollback(tx)
		return nil, errors.New(s.Trx.GetMessage("error_others"))
	}

	data := s.generateTrxResponse(res)
	data.CreatedAt = helper.StrPtr(res.Transaction.UpdatedAt.Format("2006-01-02 15:04:05 MST"))
	data.ShortDescription = res.Transaction.ShortDescription
	data.AssetName = &res.AssetName
	data.StsPay = helper.StrPtr(s.Installment.GetMapStsPay(iStsPayPending))

	return data, nil
}

// generateTrxResponse
func (s *TrxService) generateTrxResponse(res *model.Installment) *model.TrxResponse {
	remainingTenor, _ := s.Installment.GetRemainingTenor(res.ContractNumber)

	return &model.TrxResponse{
		InstallmentId:  &res.Id,
		CustomerName:   helper.StrPtr(s.IdentityRule.CreateMask(res.Customer.FullName)),
		ContractNumber: &res.ContractNumber,
		AdminFee:       &res.AdminFee,
		Currency:       helper.StrPtr(s.IdentityRule.GetCurrency()),
		InstalmentAmt:  &res.InstallmentAmt,
		InterestAmt:    &res.InterestAmt,
		OtrAmt:         &res.OtrAmt,
		Tenor:          &res.Base.Sort,
		Remaining:      &remainingTenor,
		TotalAmt:       &res.TotalAmt,
		StsPay:         helper.StrPtr(s.Installment.GetMapStsPay(res.PaymentStatus)),
	}
}
