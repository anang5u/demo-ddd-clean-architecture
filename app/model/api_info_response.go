package model

type InstallmentInfoResponse struct {
	CustomerName   *string `json:"nama_konsumen"`
	ContractNumber *string `json:"nomor_kontrak"`
}
