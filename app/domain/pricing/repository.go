package pricing

type PricingRepository interface {
	GetOptionAdminFee() float64
	GetOptionInterestAmt() float64
}
