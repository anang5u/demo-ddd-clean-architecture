package model

type PaymentRequest struct {
	IntallmentID     string `json:"id_transaksi"`
	ShortDescription string `json:"keterangan"`
}

type PaymentTokenRequest struct {
	IntallmentID string `json:"id_transaksi"`
	Token        string `json:"token"`
}
