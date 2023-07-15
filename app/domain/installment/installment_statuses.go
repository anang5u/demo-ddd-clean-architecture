package installment

// GetStatusActive
func (m *modInstallment) GetStatusActive() int64 {
	return int64(stsActiveInstallment)
}

// GetStatusPayUnpaid
func (m *modInstallment) GetStatusPayUnpaid() int64 {
	return int64(stsPayUnpaid)
}

// GetStsPayNeedConfirm
func (m *modInstallment) GetStsPayNeedConfirm() int64 {
	return int64(stsPayNeedConfirm)
}

// GetStatusPayPending
func (m *modInstallment) GetStatusPayPending() int64 {
	return int64(stsPayPending)
}

// GetStatusPayPaid
func (m *modInstallment) GetStatusPayPaid() int64 {
	return int64(stsPayPaid)
}

// GetStatusPayFailed
func (m *modInstallment) GetStatusPayFailed() int64 {
	return int64(stsPayFailed)
}

// GetMapStsPay
func (m *modInstallment) GetMapStsPay(iStsPay int64) string {
	if val, ok := stsPayMap[int64(iStsPay)]; ok {
		return val
	}
	return "-"
}
