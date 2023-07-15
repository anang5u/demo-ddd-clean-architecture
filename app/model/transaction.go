package model

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	Base
	InstallmentId    *uuid.UUID `json:"installment_id" gorm:"type:varchar(36)"`
	ContractNumber   *string    `json:"contract_number" gorm:"type:varchar(25);index;not null;"` // Nomor Kontrak untuk setiap transaksi konsumen
	OtrAmt           *float64   `json:"otr_amt" gorm:"type:float;"`                              // Angka On The Road transaksi barang baik itu White Godds, Motor atau Mobil konsumen
	AdminFee         *float64   `json:"admin_fee" gorm:"type:float;"`                            // Angka admin transaksi barang
	InstallmentAmt   *float64   `json:"installment_amt" gorm:"type:float;not null"`              // Angka jumlah cicilan transaksi
	InterestAmt      *float64   `json:"interest_amt" gorm:"type:float"`                          // Angka bunga yang ditagihkan setiap transaksi
	TotalAmt         *float64   `json:"total_amt" gorm:"type:float"`
	AssetName        *string    `json:"asset_name" gorm:"type:varchar(255);not null;"`
	PaymentStatus    *int64     `json:"payment_status" gorm:"type:smallint;default:0"`
	PaymetExpiredAt  *time.Time `json:"payment_expired" gorm:"type:datetime"`
	Refnum           *string    `json:"refnum" gorm:"type:varchar(80)"`
	CustomerId       *uuid.UUID `json:"customer_id" gorm:"type:varchar(36)"`
	FullName         *string    `json:"full_name" gorm:"type:varchar(255);"`
	LegalName        *string    `json:"legal_name" gorm:"type:varchar(255);"`
	IdCardNumber     *string    `json:"id_card_number" gorm:"type:varchar(20);"`
	Token            *string    `json:"token" gorm:"type:varchar(255);"`
	ShortDescription *string    `json:"short_desc" gorm:"type:varchar(80);"`
	LongDescription  *string    `json:"long_desc" gorm:"type:varchar(255);"`
}

func (m *Transaction) TableName() string {
	return `transaction`
}
