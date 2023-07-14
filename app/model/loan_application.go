package model

import (
	"demo-ddd-clean-architecture/app/helper"
	"time"

	"github.com/google/uuid"
)

type LoanApplication struct {
	Base
	Limit        float64    `json:"limit" gorm:"type:float;"` // Nominal limit pengajuan peminjaman
	AssetName    string     `json:"asset_name" gorm:"type:varchar(255);not null;"`
	ApprovedDate *time.Time `json:"approved_date" gorm:"type:datetime;default:null"` // Tanggal Pengajuan di setujui
	ApprovedBy   *uuid.UUID `json:"approved_by" gorm:"type:varchar(36)"`
	CustomerId   uuid.UUID  `json:"customer_id,omitempty" gorm:"type:varchar(36)"`
	Customer     *Customer  `json:"customer,omitempty" gorm:"foreignKey:CustomerId;references:Id"`
}

func (m *LoanApplication) TableName() string {
	return `loan_application`
}

// Seed
func (m *LoanApplication) Seed() *[]LoanApplication {
	loan := []LoanApplication{}

	now := time.Now()
	approvedBy := helper.UuidNew()

	loan = append(loan, LoanApplication{
		Base: Base{
			Id:        helper.UuidMustParse("91907819-cf4a-4a16-9e2d-b11a100ac1ea"),
			Sort:      1,
			Status:    1, // New
			CreatedAt: &now,
		},
		CustomerId: helper.UuidMustParse("71e11445-8f94-493d-bd19-d7e43a1e576c"),
		Limit:      10000000,
		AssetName:  "Mobil Daihatsu Xenia",
	})

	loan = append(loan, LoanApplication{
		Base: Base{
			Id:        helper.UuidMustParse("3644a03c-b69c-4854-828e-4178fdac987d"),
			Sort:      1,
			Status:    2, // Approved
			CreatedAt: &now,
		},
		CustomerId:   helper.UuidMustParse("71e11445-8f94-493d-bd19-d7e43a1e576c"),
		Limit:        500000,
		ApprovedDate: &now,
		ApprovedBy:   &approvedBy,
		AssetName:    "LED TV",
	})

	loan = append(loan, LoanApplication{
		Base: Base{
			Id:        helper.UuidMustParse("a33f2062-a637-454d-bbcc-2e2f3f2c4c8d"),
			Sort:      2,
			Status:    2, // Approved
			CreatedAt: &now,
		},
		CustomerId:   helper.UuidMustParse("7a49c137-84d4-433f-810c-3188ddef783f"),
		Limit:        2000000,
		ApprovedDate: &now,
		ApprovedBy:   &approvedBy,
		AssetName:    "Laptop",
	})

	return &loan
}
