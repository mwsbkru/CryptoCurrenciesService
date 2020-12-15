package currencies

import "models"

type IRemoteDataProvider interface {
	GetCurrencyRate(code string) models.CurrencyRate, error
}

type ICacheDataProvider interface {
	GetCachedCurrency(code string) models.CurrencyRate, error
	StoreCurrency(code string, currencyRate models.CurrencyRate) error
	GetStoredCurrencies() []models.CurrencyRate
	SetCacheValidator(validator ICacheDataValidator)
}

type ICacheDataValidator interface {
	IsValid(currency models.CurrencyRate) bool
}

type IErrorsReporter interface {
	ReportError(err error)
}

type IAsyncTaskRunner interface {
	RunTask(task AsyncTastFunc)
}

type AsyncTastFunc func()
