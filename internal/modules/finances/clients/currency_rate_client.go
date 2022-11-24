package finance_clients

import (
	"context"
	"fmt"
	"sync"

	"github.com/me-io/go-swap/pkg/exchanger"
	"github.com/me-io/go-swap/pkg/swap"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"

	"github.com/shav/telegram-bot/internal/common/multi_error"
	"github.com/shav/telegram-bot/internal/modules/finances/domain/models"
	"github.com/shav/telegram-bot/internal/observability/tracing"
)

// CurrencyRateClient является клиентом для загрузки курсов обмена валют из открытых источников.
type CurrencyRateClient struct {
	// Клиент для загрузки курсов обмена валют.
	client *swap.Swap
	// Объект синхронизации для загрузки курса валют.
	lock *sync.Mutex
}

// NewCurrencyRateClient создаёт новый клиент для загрузки курсов обмена валют.
func NewCurrencyRateClient() *CurrencyRateClient {
	client := swap.NewSwap().
		AddExchanger(exchanger.NewYahooApi(nil)).
		Build()
	return &CurrencyRateClient{
		client: client,
		lock:   &sync.Mutex{},
	}
}

// LoadRate загружает курс обмена валюты baseCurrency на валюту targetCurrency.
func (c *CurrencyRateClient) LoadRate(ctx context.Context, baseCurrency finance_models.Currency, targetCurrency finance_models.Currency) (rate decimal.Decimal, err error) {
	span, _ := tracing.StartSpanFromContext(ctx, "CurrencyRateClient.LoadRate")
	defer span.Finish()

	c.lock.Lock()
	defer c.lock.Unlock()

	defer func() {
		if r := recover(); r != nil {
			switch e := r.(type) {
			case string:
				err = errors.New(e)
			case error:
				err = e
			case map[string]string:
				err = nil
				for srv, msg := range e {
					err = multi_error.Append(err, errors.New(fmt.Sprintf("%s: %s", srv, msg)))
				}
			default:
				err = errors.New(fmt.Sprintf("An error occurred:\n%v", e))
			}
			err = errors.Wrap(err, "load rate")
		}
	}()

	r := c.client.Latest(fmt.Sprintf("%s/%s", targetCurrency.Code, baseCurrency.Code))
	return decimal.NewFromFloat(r.GetRateValue()), nil
}
