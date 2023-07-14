package model

import (
	"github.com/google/uuid"
)

type Installment struct {
	Base
	ContractNumber string    `json:"full_name" gorm:"type:varchar(25);index;not null;"` // Nomor Kontrak untuk setiap transaksi konsumen
	OtrAmt         float64   `json:"otr_amt" gorm:"type:float;"`                        // Angka On The Road transaksi barang baik itu White Godds, Motor atau Mobil konsumen
	AdminFee       float64   `json:"admin_fee" gorm:"type:float;"`                      // Angka admin transaksi barang
	InstallmentAmt float64   `json:"installment_amt" gorm:"type:float;not null"`        // Angka jumlah cicilan transaksi
	InterestAmt    float64   `json:"interest_amt" gorm:"type:float"`                    // Angka bunga yang ditagihkan setiap transaksi
	TotalAmt       float64   `json:"total_amt" gorm:"type:float"`
	AssetName      string    `json:"asset_name" gorm:"type:varchar(255);not null;"`
	PaymentStatus  int       `json:"payment_status" gorm:"type:smallint;default:0"`
	CustomerId     uuid.UUID `json:"customer_id,omitempty" gorm:"type:varchar(36)" format:"uuid"`
	Customer       *Customer `json:"customer,omitempty" gorm:"foreignKey:CustomerId;references:Id"`
}

func (m *Installment) TableName() string {
	return `installment`
}
