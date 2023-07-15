package installment

var (
	// Installment Status
	stsActiveInstallment = 1

	// Pyment Status
	stsPayUnpaid      = 0
	stsPayNeedConfirm = 1 // menunggu validasi token
	stsPayPending     = 2 // menunggu callback dri thirdparty payment gateway
	stsPayPaid        = 3 // pembayaran sukses
	stsPayFailed      = 4 // pembayaran gagal

	stsPayMap = map[int64]string{
		0: "Belum Bayar",
		1: "Need Confirm",
		2: "Pembayaran Pending",
		3: "Pembayaran Sukses",
		4: "Pembayaran Gagal",
	}
)
