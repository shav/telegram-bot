package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i github.com/shav/telegram-bot/internal/modules/finances/services/currency/convert.currencyRateGetter -o ./mocks\currency_rate_getter.go -n CurrencyRateGetterMock

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	finance_models "github.com/shav/telegram-bot/internal/modules/finances/domain/models"
)

// CurrencyRateGetterMock implements finance_services_currency.currencyRateGetter
type CurrencyRateGetterMock struct {
	t minimock.Tester

	funcGetRate          func(ctx context.Context, currency finance_models.Currency) (c2 finance_models.CurrencyRate, err error)
	inspectFuncGetRate   func(ctx context.Context, currency finance_models.Currency)
	afterGetRateCounter  uint64
	beforeGetRateCounter uint64
	GetRateMock          mCurrencyRateGetterMockGetRate
}

// NewCurrencyRateGetterMock returns a mock for finance_services_currency.currencyRateGetter
func NewCurrencyRateGetterMock(t minimock.Tester) *CurrencyRateGetterMock {
	m := &CurrencyRateGetterMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.GetRateMock = mCurrencyRateGetterMockGetRate{mock: m}
	m.GetRateMock.callArgs = []*CurrencyRateGetterMockGetRateParams{}

	return m
}

type mCurrencyRateGetterMockGetRate struct {
	mock               *CurrencyRateGetterMock
	defaultExpectation *CurrencyRateGetterMockGetRateExpectation
	expectations       []*CurrencyRateGetterMockGetRateExpectation

	callArgs []*CurrencyRateGetterMockGetRateParams
	mutex    sync.RWMutex
}

// CurrencyRateGetterMockGetRateExpectation specifies expectation struct of the currencyRateGetter.GetRate
type CurrencyRateGetterMockGetRateExpectation struct {
	mock    *CurrencyRateGetterMock
	params  *CurrencyRateGetterMockGetRateParams
	results *CurrencyRateGetterMockGetRateResults
	Counter uint64
}

// CurrencyRateGetterMockGetRateParams contains parameters of the currencyRateGetter.GetRate
type CurrencyRateGetterMockGetRateParams struct {
	ctx      context.Context
	currency finance_models.Currency
}

// CurrencyRateGetterMockGetRateResults contains results of the currencyRateGetter.GetRate
type CurrencyRateGetterMockGetRateResults struct {
	c2  finance_models.CurrencyRate
	err error
}

// Expect sets up expected params for currencyRateGetter.GetRate
func (mmGetRate *mCurrencyRateGetterMockGetRate) Expect(ctx context.Context, currency finance_models.Currency) *mCurrencyRateGetterMockGetRate {
	if mmGetRate.mock.funcGetRate != nil {
		mmGetRate.mock.t.Fatalf("CurrencyRateGetterMock.GetRate mock is already set by Set")
	}

	if mmGetRate.defaultExpectation == nil {
		mmGetRate.defaultExpectation = &CurrencyRateGetterMockGetRateExpectation{}
	}

	mmGetRate.defaultExpectation.params = &CurrencyRateGetterMockGetRateParams{ctx, currency}
	for _, e := range mmGetRate.expectations {
		if minimock.Equal(e.params, mmGetRate.defaultExpectation.params) {
			mmGetRate.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetRate.defaultExpectation.params)
		}
	}

	return mmGetRate
}

// Inspect accepts an inspector function that has same arguments as the currencyRateGetter.GetRate
func (mmGetRate *mCurrencyRateGetterMockGetRate) Inspect(f func(ctx context.Context, currency finance_models.Currency)) *mCurrencyRateGetterMockGetRate {
	if mmGetRate.mock.inspectFuncGetRate != nil {
		mmGetRate.mock.t.Fatalf("Inspect function is already set for CurrencyRateGetterMock.GetRate")
	}

	mmGetRate.mock.inspectFuncGetRate = f

	return mmGetRate
}

// Return sets up results that will be returned by currencyRateGetter.GetRate
func (mmGetRate *mCurrencyRateGetterMockGetRate) Return(c2 finance_models.CurrencyRate, err error) *CurrencyRateGetterMock {
	if mmGetRate.mock.funcGetRate != nil {
		mmGetRate.mock.t.Fatalf("CurrencyRateGetterMock.GetRate mock is already set by Set")
	}

	if mmGetRate.defaultExpectation == nil {
		mmGetRate.defaultExpectation = &CurrencyRateGetterMockGetRateExpectation{mock: mmGetRate.mock}
	}
	mmGetRate.defaultExpectation.results = &CurrencyRateGetterMockGetRateResults{c2, err}
	return mmGetRate.mock
}

//Set uses given function f to mock the currencyRateGetter.GetRate method
func (mmGetRate *mCurrencyRateGetterMockGetRate) Set(f func(ctx context.Context, currency finance_models.Currency) (c2 finance_models.CurrencyRate, err error)) *CurrencyRateGetterMock {
	if mmGetRate.defaultExpectation != nil {
		mmGetRate.mock.t.Fatalf("Default expectation is already set for the currencyRateGetter.GetRate method")
	}

	if len(mmGetRate.expectations) > 0 {
		mmGetRate.mock.t.Fatalf("Some expectations are already set for the currencyRateGetter.GetRate method")
	}

	mmGetRate.mock.funcGetRate = f
	return mmGetRate.mock
}

// When sets expectation for the currencyRateGetter.GetRate which will trigger the result defined by the following
// Then helper
func (mmGetRate *mCurrencyRateGetterMockGetRate) When(ctx context.Context, currency finance_models.Currency) *CurrencyRateGetterMockGetRateExpectation {
	if mmGetRate.mock.funcGetRate != nil {
		mmGetRate.mock.t.Fatalf("CurrencyRateGetterMock.GetRate mock is already set by Set")
	}

	expectation := &CurrencyRateGetterMockGetRateExpectation{
		mock:   mmGetRate.mock,
		params: &CurrencyRateGetterMockGetRateParams{ctx, currency},
	}
	mmGetRate.expectations = append(mmGetRate.expectations, expectation)
	return expectation
}

// Then sets up currencyRateGetter.GetRate return parameters for the expectation previously defined by the When method
func (e *CurrencyRateGetterMockGetRateExpectation) Then(c2 finance_models.CurrencyRate, err error) *CurrencyRateGetterMock {
	e.results = &CurrencyRateGetterMockGetRateResults{c2, err}
	return e.mock
}

// GetRate implements finance_services_currency.currencyRateGetter
func (mmGetRate *CurrencyRateGetterMock) GetRate(ctx context.Context, currency finance_models.Currency) (c2 finance_models.CurrencyRate, err error) {
	mm_atomic.AddUint64(&mmGetRate.beforeGetRateCounter, 1)
	defer mm_atomic.AddUint64(&mmGetRate.afterGetRateCounter, 1)

	if mmGetRate.inspectFuncGetRate != nil {
		mmGetRate.inspectFuncGetRate(ctx, currency)
	}

	mm_params := &CurrencyRateGetterMockGetRateParams{ctx, currency}

	// Record call args
	mmGetRate.GetRateMock.mutex.Lock()
	mmGetRate.GetRateMock.callArgs = append(mmGetRate.GetRateMock.callArgs, mm_params)
	mmGetRate.GetRateMock.mutex.Unlock()

	for _, e := range mmGetRate.GetRateMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.c2, e.results.err
		}
	}

	if mmGetRate.GetRateMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetRate.GetRateMock.defaultExpectation.Counter, 1)
		mm_want := mmGetRate.GetRateMock.defaultExpectation.params
		mm_got := CurrencyRateGetterMockGetRateParams{ctx, currency}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetRate.t.Errorf("CurrencyRateGetterMock.GetRate got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetRate.GetRateMock.defaultExpectation.results
		if mm_results == nil {
			mmGetRate.t.Fatal("No results are set for the CurrencyRateGetterMock.GetRate")
		}
		return (*mm_results).c2, (*mm_results).err
	}
	if mmGetRate.funcGetRate != nil {
		return mmGetRate.funcGetRate(ctx, currency)
	}
	mmGetRate.t.Fatalf("Unexpected call to CurrencyRateGetterMock.GetRate. %v %v", ctx, currency)
	return
}

// GetRateAfterCounter returns a count of finished CurrencyRateGetterMock.GetRate invocations
func (mmGetRate *CurrencyRateGetterMock) GetRateAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetRate.afterGetRateCounter)
}

// GetRateBeforeCounter returns a count of CurrencyRateGetterMock.GetRate invocations
func (mmGetRate *CurrencyRateGetterMock) GetRateBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetRate.beforeGetRateCounter)
}

// Calls returns a list of arguments used in each call to CurrencyRateGetterMock.GetRate.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetRate *mCurrencyRateGetterMockGetRate) Calls() []*CurrencyRateGetterMockGetRateParams {
	mmGetRate.mutex.RLock()

	argCopy := make([]*CurrencyRateGetterMockGetRateParams, len(mmGetRate.callArgs))
	copy(argCopy, mmGetRate.callArgs)

	mmGetRate.mutex.RUnlock()

	return argCopy
}

// MinimockGetRateDone returns true if the count of the GetRate invocations corresponds
// the number of defined expectations
func (m *CurrencyRateGetterMock) MinimockGetRateDone() bool {
	for _, e := range m.GetRateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetRateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetRateCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetRate != nil && mm_atomic.LoadUint64(&m.afterGetRateCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetRateInspect logs each unmet expectation
func (m *CurrencyRateGetterMock) MinimockGetRateInspect() {
	for _, e := range m.GetRateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to CurrencyRateGetterMock.GetRate with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetRateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetRateCounter) < 1 {
		if m.GetRateMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to CurrencyRateGetterMock.GetRate")
		} else {
			m.t.Errorf("Expected call to CurrencyRateGetterMock.GetRate with params: %#v", *m.GetRateMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetRate != nil && mm_atomic.LoadUint64(&m.afterGetRateCounter) < 1 {
		m.t.Error("Expected call to CurrencyRateGetterMock.GetRate")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *CurrencyRateGetterMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockGetRateInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *CurrencyRateGetterMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *CurrencyRateGetterMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockGetRateDone()
}