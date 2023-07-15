package model

import (
	"github.com/google/uuid"
)

type TrxResponse struct {
	InstallmentId    *uuid.UUID `json:"id_transaksi"`
	CreatedAt        *string    `json:"tgl_transaksi,omitempty"`
	CustomerName     *string    `json:"nama_konsumen"`
	ContractNumber   *string    `json:"nomor_kontrak"`
	OtrAmt           *float64   `json:"nominal_otr"`
	Tenor            *int64     `json:"cicilan_ke"`
	Remaining        *int64     `json:"sisa_tenor"`
	AdminFee         *float64   `json:"biaya_admin"`
	InstalmentAmt    *float64   `json:"cicilan"`
	InterestAmt      *float64   `json:"margin"`
	TotalAmt         *float64   `json:"total"`
	Currency         *string    `json:"mata_uang"`
	StsPay           *string    `json:"status_pembayaran"`
	AssetName        *string    `json:"asset_pembiayaan,omitempty"`
	ShortDescription *string    `json:"keterangan_1,omitempty"`
	LongDescription  *string    `json:"keterangan_2,omitempty"`
	Token            *string    `json:"token_pembayaran,omitempty"`
}
