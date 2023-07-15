package transaction

var (
	mapMessage = map[string]string{
		"error_others":         "Transaksi anda tidak dapat di proses. Silahkan hubungi CS kami!",
		"error_mismatch_token": "Transaksi tidak dapat di proses. Pastikan token yg anda input sudah benar!",
	}
)

func (m *modTrx) GetMessage(key string) string {
	message := ""
	if val, ok := mapMessage[key]; ok {
		message = val
	}

	return message
}
