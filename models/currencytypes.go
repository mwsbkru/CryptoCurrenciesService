package models

type CurrencyRate struct {
	Code string
	RateTo string
	Rate float64
	Timestamp int
}