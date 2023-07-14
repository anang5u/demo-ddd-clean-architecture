package installment

const (

	// Installment Status
	stsNewInstallment    = 1 // status=1
	stsActiveInstallment = 2 // status=2 => cicilan berjalan
	stsDoneInstallment   = 3 // status=3 => cicilan selesai, tidak lagi mempunai tenor
)
