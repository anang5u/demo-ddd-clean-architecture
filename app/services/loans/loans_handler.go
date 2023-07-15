package loans

import (
	"demo-ddd-clean-architecture/app/helper"
	"demo-ddd-clean-architecture/app/model"
	"math"

	"github.com/google/uuid"
)

// GenerateInstallment
func (s *LoanService) GenerateInstallment() {
	// mencegah duplicate generate
	s.Lock()
	defer s.Unlock()

	loanCustomers, err := s.Loans.GetApprovedLoans()
	if err != nil {
		s.Common.Log.Error("GenerateInstallment", err.Error())
		return
	}

	var loanIDs []uuid.UUID
	var installments []model.Installment

	for _, loan := range *loanCustomers {
		// cek, apakah cicilan sudah di generate
		isExists, err := s.Installment.IsActiveInstallmentExists(loan.Customer)
		if err != nil {
			s.Common.Log.Error("IsNewInstallmentExists check Error, cust_id: ", loan.Customer.Id, err.Error())
			continue
		}

		// skip jika sudah di generate
		if isExists {
			continue
		}

		// contract number
		contractNumber, err := s.IdentityRule.GenerateContractNumber()
		if err != nil {
			s.Common.Log.Error("GenerateContractNumber Error, cust_id: ", loan.Customer.Id, err.Error())
			continue
		}

		// tenor
		tenor, err := s.Loans.CalculateTenor(&loan)
		if err != nil {
			s.Common.Log.Error("CalculateTenor Error, cust_id: ", loan.Customer.Id, err.Error())
			continue
		}

		loanIDs = append(loanIDs, loan.Id)
		createdAt := helper.GetTimeNow()
		admFee := s.Pricing.GetOptionAdminFee()                  // admin fee
		installmentAmt := math.Ceil(loan.Limit / float64(tenor)) // intallmentAmount
		interestAmt := math.Ceil((loan.Limit * s.Pricing.GetOptionInterestAmt() / float64(100)) / float64(tenor))

		// list cicilan untuk masing2 tenor
		for i := 1; i <= tenor; i++ {
			installments = append(installments, model.Installment{
				Base: model.Base{
					Id:        helper.UuidNew(),
					Sort:      int64(i),
					CreatedAt: &createdAt,
					Status:    int(s.Installment.GetStatusActive()),
				},
				ContractNumber: contractNumber,
				OtrAmt:         loan.Limit, // untuk saat ini di isi dengan nominal jumlah pengajuan ??
				AdminFee:       admFee,
				InstallmentAmt: installmentAmt,
				InterestAmt:    interestAmt,
				TotalAmt:       (admFee + installmentAmt + interestAmt),
				AssetName:      loan.AssetName,
				CustomerId:     loan.CustomerId,
			})
		}

	}

	if len(installments) > 0 {
		tx := s.Common.Db.TxBegin()
		defer s.Common.Db.TxRecover(tx)

		// create installment
		if err := s.Installment.WithTx(tx).Create(&installments); err != nil {
			s.Common.Log.Error("Error while Create installment: ", err.Error())
			s.Common.Db.TxRollback(tx)
			return
		}

		// Update pengajuan menjadi status done
		if err := s.Loans.WithTx(tx).UpdateToStatusDone(&loanIDs); err != nil {
			s.Common.Log.Error("Error while UpdateToStatusDone loans: ", err.Error())
			s.Common.Db.TxRollback(tx)
			return
		}

		// Commit
		if err := tx.Commit().Error; err != nil {
			s.Common.Log.Error("Error while Commit GenerateInstallment: ", err.Error())
		}

	}
}
