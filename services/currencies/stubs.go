package currencies

import (
	"../../models"
)

type StubRemoteDataProvider struct {}

func (provider StubRemoteDataProvider) GetCurrencyRate(code string) models.CurrencyRate, err {
	rates := make(map[string]models.CurrencyRate)

	rates["btc"] := models.CurrencyRate{
		Code = "BTC",
		RateTo = "USD",
		Rate = 20000,
		timestamp = 100
	}

	rates["eth"] := models.CurrencyRate{
		Code = "ETH",
		RateTo = "USD",
		Rate = 500,
		timestamp = 100
	}

	rates["xmr"] := models.CurrencyRate{
		Code = "XMR",
		RateTo = "USD",
		Rate = 150,
		timestamp = 100
	}

	if val, ok := rates[code]; ok {
		return rates[code]
	}

	return nil, fmt.Errorf("No remote data for code: %s", code)
}

//////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////

type StubCacheDataValidator struct {}

func (validator StubCacheDataValidator) IsValid(currency models.CurrencyRate) bool {
	currentTimeStamp := 100
	expirePeriod := 20

	return currentTimeStamp - currency.Timestamp < expirePeriod
}

//////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////

type StubCacheDataProvider struct {
	cacheValidator ICacheDataValidator
}

func (provider *StubCacheDataProvider) SetCacheValidator(validator ICacheDataValidator) {
	(*provider).cacheValidator = validator
}

func (provider StubCacheDataProvider) GetCachedCurrency(code string) models.CurrencyRate {
	rates := make(map[string]models.CurrencyRate)

	rates["btc"] := models.CurrencyRate{
		Code = "BTC",
		RateTo = "USD",
		Rate = 20000,
		timestamp = 90
	}

	rates["eth"] := models.CurrencyRate{
		Code = "ETH",
		RateTo = "USD",
		Rate = 500,
		timestamp = 50
	}

	if val, ok := rates[code]; !ok {
		return nil, fmt.Errorf("No cached data for code: %s", code)
	}

	if provider.cacheValidator.IsValid(rates[code]) {
		return rates[code], nil
	}

	return nil, fmt.Errorf("Data for code: %s is expired", code)
}

func (provider StubCacheDataProvider) StoreCurrency(code string, currencyRate models.CurrencyRate) error {
	return nil
}

func (provider StubCacheDataProvider) GetStoredCurrencies() []models.CurrencyRate {
	rates := make([]models.CurrencyRate, 2)
	rates[0] =  models.CurrencyRate{
		Code = "BTC",
		RateTo = "USD",
		Rate = 20000,
		timestamp = 90
	}
	rates[1] =  models.CurrencyRate{
		Code = "ETH",
		RateTo = "USD",
		Rate = 500,
		timestamp = 50
	}

	return rates
}