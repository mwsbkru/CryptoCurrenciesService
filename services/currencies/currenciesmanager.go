package currencies

import (
	"models"
)

type CurrenciesManager struct {
	remoteDataProvider IRemoteDataProvider
	cacheDataProvider ICacheDataProvider
}

func GetCurrencyRate(code string) models.CurrencyRate {
	return nil
}
func GetCachedCurrencyRates() []models.CurrencyRate {
	return nil
}