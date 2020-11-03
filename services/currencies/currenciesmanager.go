package currencies

import (
	"../../models"
)

type CurrenciesManager struct {
	remoteDataProvider IRemoteDataProvider
	cacheDataProvider ICacheDataProvider
	cacheDataValidator ICacheDataValidator
	errorsReporter IErrorsReporter
	asyncRunner IAsyncTaskRunner
}

func (manager CurrenciesManager) GetCurrencyRate(code string) models.CurrencyRate, err {
	currencyRate, err := manager.cacheDataProvider.GetCachedCurrency(code)
	if err != nil || !manager.cacheDataValidator.isValid(currencyRate) {
		currencyRate, err = manager.prepareNewValue(code)

		if err != nil {
			manager.errorsReporter.ReportError(err)
			return currencyRate, errors.New("Can`t fetch any data about currency rate!")
		}

		return currencyRate, nil 
	}
	
	manager.updateCachedCurrencyAsync(code)
	return currencyRate, nil
}

func (manager CurrenciesManager) prepareNewValue(code string) models.CurrencyRate, err {
	loadedCurrencyRate, err := manager.remoteDataProvider.GetCurrencyRate(code)

	if err != nil {
		return loadedCurrencyRate, err
	}

	err := manager.cacheDataProvider.StoreCurrency(code, loadedCurrencyRate)

	if err != nil {
		manager.errorsReporter.ReportError(err)
	}

	return loadedCurrencyRate, nil
}

func (manager CurrenciesManager) updateCachedCurrencyAsync(code string) {
	manager.asyncRunner.Run(func() {
		manager.prepareNewValue(code)
	})
}

func GetCachedCurrencyRates() []models.CurrencyRate {
	return nil
}