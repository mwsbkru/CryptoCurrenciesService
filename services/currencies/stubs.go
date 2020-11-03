package currencies

import (
	"../../models"
)

type Struct StubRemoteDataProvider {}

func (provider StubRemoteDataProvider) GetCurrencyRate(code string) models.CurrencyRate, err {
	rate := models.CurrencyRate{
		Code = "BTC",
		RateTo = "USD",
		Rate = 12000,
		timestamp = 1234567
	}

	return rate, nil
}