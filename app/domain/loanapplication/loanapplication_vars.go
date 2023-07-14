package loanapplication

// Statuses
const (
	stsNew      = 1 // status=1
	stsApproved = 2 // status=2
	stsRejected = 3 // status=3
	stsDone     = 4 // cicilan sudah di generate
)

// Error
const (
	errEmptyLoanApplication = "Pengajuan tidak ditemukan. Info lebih lanjut silahkan kontak CS kami!"
)
