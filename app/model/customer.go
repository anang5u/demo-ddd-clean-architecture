package model

import (
	"demo-ddd-clean-architecture/app/helper"
	"time"
)

type Customer struct {
	Base
	IdCardNumber string  `json:"id_card_number" gorm:"type:varchar(16);index;not null;"` // Nomor KTP Konsumen
	FullName     string  `json:"full_name" gorm:"type:varchar(255);not null;"`           // Nama Lengkap Konsumen
	LegalName    string  `json:"legal_name" gorm:"type:varchar(255);not null;"`          // Nama Konsumen Di KTP
	PlaceOfBirth string  `json:"place_of_birth" gorm:"type:varchar(100)"`                // Tempat lahir konsumen sesuai KTP
	DateOfBirth  string  `json:"date_of_birth" gorm:"type:varchar(20);not null;"`        // Tanggal lahir konsumen sesuai KTP
	Salary       float64 `json:"salary" gorm:"type:float"`                               // Gaji Konsumen
	IdCardPhoto  string  `json:"id_card_photo" gorm:"type:varchar(100)"`                 // Foto KTP Konsumen
	SelfiePhoto  string  `json:"product_name" gorm:"type:varchar(100)"`                  // Foto Selfie Konsumen
}

func (m *Customer) TableName() string {
	return `customer`
}

// Seed
func (m *Customer) Seed() *[]Customer {
	customers := []Customer{}

	customers = append(customers, Customer{
		Base: Base{
			Id:        helper.UuidMustParse("71e11445-8f94-493d-bd19-d7e43a1e576c"),
			Sort:      1,
			CreatedAt: time.Now(),
		},
		IdCardNumber: "3211111111111111",
		FullName:     "Budi",
		LegalName:    "Budi",
		PlaceOfBirth: "Palembang",
		DateOfBirth:  "1987-11-22",
		Salary:       3000000,
		IdCardPhoto:  "budi-ktp-11223344.jpg",
		SelfiePhoto:  "budi-selfie-11223355.jpg",
	})

	customers = append(customers, Customer{
		Base: Base{
			Id:        helper.UuidMustParse("7a49c137-84d4-433f-810c-3188ddef783f"),
			Sort:      2,
			CreatedAt: time.Now(),
		},
		IdCardNumber: "3222222222222222",
		FullName:     "Annisa",
		LegalName:    "Annisa",
		PlaceOfBirth: "Bandung",
		DateOfBirth:  "1992-01-27",
		Salary:       12000000,
		IdCardPhoto:  "annisa-ktp-11223366.jpg",
		SelfiePhoto:  "budi-selfie-11223377.jpg",
	})

	return &customers
}
