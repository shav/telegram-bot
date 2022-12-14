package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i github.com/shav/telegram-bot/internal/modules/finances/services/currency/rates.CurrencyRateUpdater -o ./mocks\currency_rate_updater.go -n CurrencyRateUpdaterMock

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	finance_models "github.com/shav/telegram-bot/internal/modules/finances/domain/models"
)

// CurrencyRateUpdaterMock implements finance_services_currency.CurrencyRateUpdater
type CurrencyRateUpdaterMock struct {
	t minimock.Tester

	funcGetCurrency          func() (c1 finance_models.Currency)
	inspectFuncGetCurrency   func()
	afterGetCurrencyCounter  uint64
	beforeGetCurrencyCounter uint64
	GetCurrencyMock          mCurrencyRateUpdaterMockGetCurrency

	funcStartMonitoringRate          func(ctx context.Context)
	inspectFuncStartMonitoringRate   func(ctx context.Context)
	afterStartMonitoringRateCounter  uint64
	beforeStartMonitoringRateCounter uint64
	StartMonitoringRateMock          mCurrencyRateUpdaterMockStartMonitoringRate

	funcUpdateRateAsync          func(ctx context.Context) (ch1 <-chan finance_models.CurrencyRate)
	inspectFuncUpdateRateAsync   func(ctx context.Context)
	afterUpdateRateAsyncCounter  uint64
	beforeUpdateRateAsyncCounter uint64
	UpdateRateAsyncMock          mCurrencyRateUpdaterMockUpdateRateAsync
}

// NewCurrencyRateUpdaterMock returns a mock for finance_services_currency.CurrencyRateUpdater
func NewCurrencyRateUpdaterMock(t minimock.Tester) *CurrencyRateUpdaterMock {
	m := &CurrencyRateUpdaterMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.GetCurrencyMock = mCurrencyRateUpdaterMockGetCurrency{mock: m}

	m.StartMonitoringRateMock = mCurrencyRateUpdaterMockStartMonitoringRate{mock: m}
	m.StartMonitoringRateMock.callArgs = []*CurrencyRateUpdaterMockStartMonitoringRateParams{}

	m.UpdateRateAsyncMock = mCurrencyRateUpdaterMockUpdateRateAsync{mock: m}
	m.UpdateRateAsyncMock.callArgs = []*CurrencyRateUpdaterMockUpdateRateAsyncParams{}

	return m
}

type mCurrencyRateUpdaterMockGetCurrency struct {
	mock               *CurrencyRateUpdaterMock
	defaultExpectation *CurrencyRateUpdaterMockGetCurrencyExpectation
	expectations       []*CurrencyRateUpdaterMockGetCurrencyExpectation
}

// CurrencyRateUpdaterMockGetCurrencyExpectation specifies expectation struct of the CurrencyRateUpdater.GetCurrency
type CurrencyRateUpdaterMockGetCurrencyExpectation struct {
	mock *CurrencyRateUpdaterMock

	results *CurrencyRateUpdaterMockGetCurrencyResults
	Counter uint64
}

// CurrencyRateUpdaterMockGetCurrencyResults contains results of the CurrencyRateUpdater.GetCurrency
type CurrencyRateUpdaterMockGetCurrencyResults struct {
	c1 finance_models.Currency
}

// Expect sets up expected params for CurrencyRateUpdater.GetCurrency
func (mmGetCurrency *mCurrencyRateUpdaterMockGetCurrency) Expect() *mCurrencyRateUpdaterMockGetCurrency {
	if mmGetCurrency.mock.funcGetCurrency != nil {
		mmGetCurrency.mock.t.Fatalf("CurrencyRateUpdaterMock.GetCurrency mock is already set by Set")
	}

	if mmGetCurrency.defaultExpectation == nil {
		mmGetCurrency.defaultExpectation = &CurrencyRateUpdaterMockGetCurrencyExpectation{}
	}

	return mmGetCurrency
}

// Inspect accepts an inspector function that has same arguments as the CurrencyRateUpdater.GetCurrency
func (mmGetCurrency *mCurrencyRateUpdaterMockGetCurrency) Inspect(f func()) *mCurrencyRateUpdaterMockGetCurrency {
	if mmGetCurrency.mock.inspectFuncGetCurrency != nil {
		mmGetCurrency.mock.t.Fatalf("Inspect function is already set for CurrencyRateUpdaterMock.GetCurrency")
	}

	mmGetCurrency.mock.inspectFuncGetCurrency = f

	return mmGetCurrency
}

// Return sets up results that will be returned by CurrencyRateUpdater.GetCurrency
func (mmGetCurrency *mCurrencyRateUpdaterMockGetCurrency) Return(c1 finance_models.Currency) *CurrencyRateUpdaterMock {
	if mmGetCurrency.mock.funcGetCurrency != nil {
		mmGetCurrency.mock.t.Fatalf("CurrencyRateUpdaterMock.GetCurrency mock is already set by Set")
	}

	if mmGetCurrency.defaultExpectation == nil {
		mmGetCurrency.defaultExpectation = &CurrencyRateUpdaterMockGetCurrencyExpectation{mock: mmGetCurrency.mock}
	}
	mmGetCurrency.defaultExpectation.results = &CurrencyRateUpdaterMockGetCurrencyResults{c1}
	return mmGetCurrency.mock
}

//Set uses given function f to mock the CurrencyRateUpdater.GetCurrency method
func (mmGetCurrency *mCurrencyRateUpdaterMockGetCurrency) Set(f func() (c1 finance_models.Currency)) *CurrencyRateUpdaterMock {
	if mmGetCurrency.defaultExpectation != nil {
		mmGetCurrency.mock.t.Fatalf("Default expectation is already set for the CurrencyRateUpdater.GetCurrency method")
	}

	if len(mmGetCurrency.expectations) > 0 {
		mmGetCurrency.mock.t.Fatalf("Some expectations are already set for the CurrencyRateUpdater.GetCurrency method")
	}

	mmGetCurrency.mock.funcGetCurrency = f
	return mmGetCurrency.mock
}

// GetCurrency implements finance_services_currency.CurrencyRateUpdater
func (mmGetCurrency *CurrencyRateUpdaterMock) GetCurrency() (c1 finance_models.Currency) {
	mm_atomic.AddUint64(&mmGetCurrency.beforeGetCurrencyCounter, 1)
	defer mm_atomic.AddUint64(&mmGetCurrency.afterGetCurrencyCounter, 1)

	if mmGetCurrency.inspectFuncGetCurrency != nil {
		mmGetCurrency.inspectFuncGetCurrency()
	}

	if mmGetCurrency.GetCurrencyMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetCurrency.GetCurrencyMock.defaultExpectation.Counter, 1)

		mm_results := mmGetCurrency.GetCurrencyMock.defaultExpectation.results
		if mm_results == nil {
			mmGetCurrency.t.Fatal("No results are set for the CurrencyRateUpdaterMock.GetCurrency")
		}
		return (*mm_results).c1
	}
	if mmGetCurrency.funcGetCurrency != nil {
		return mmGetCurrency.funcGetCurrency()
	}
	mmGetCurrency.t.Fatalf("Unexpected call to CurrencyRateUpdaterMock.GetCurrency.")
	return
}

// GetCurrencyAfterCounter returns a count of finished CurrencyRateUpdaterMock.GetCurrency invocations
func (mmGetCurrency *CurrencyRateUpdaterMock) GetCurrencyAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetCurrency.afterGetCurrencyCounter)
}

// GetCurrencyBeforeCounter returns a count of CurrencyRateUpdaterMock.GetCurrency invocations
func (mmGetCurrency *CurrencyRateUpdaterMock) GetCurrencyBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetCurrency.beforeGetCurrencyCounter)
}

// MinimockGetCurrencyDone returns true if the count of the GetCurrency invocations corresponds
// the number of defined expectations
func (m *CurrencyRateUpdaterMock) MinimockGetCurrencyDone() bool {
	for _, e := range m.GetCurrencyMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetCurrencyMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCurrencyCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetCurrency != nil && mm_atomic.LoadUint64(&m.afterGetCurrencyCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetCurrencyInspect logs each unmet expectation
func (m *CurrencyRateUpdaterMock) MinimockGetCurrencyInspect() {
	for _, e := range m.GetCurrencyMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to CurrencyRateUpdaterMock.GetCurrency")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetCurrencyMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCurrencyCounter) < 1 {
		m.t.Error("Expected call to CurrencyRateUpdaterMock.GetCurrency")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetCurrency != nil && mm_atomic.LoadUint64(&m.afterGetCurrencyCounter) < 1 {
		m.t.Error("Expected call to CurrencyRateUpdaterMock.GetCurrency")
	}
}

type mCurrencyRateUpdaterMockStartMonitoringRate struct {
	mock               *CurrencyRateUpdaterMock
	defaultExpectation *CurrencyRateUpdaterMockStartMonitoringRateExpectation
	expectations       []*CurrencyRateUpdaterMockStartMonitoringRateExpectation

	callArgs []*CurrencyRateUpdaterMockStartMonitoringRateParams
	mutex    sync.RWMutex
}

// CurrencyRateUpdaterMockStartMonitoringRateExpectation specifies expectation struct of the CurrencyRateUpdater.StartMonitoringRate
type CurrencyRateUpdaterMockStartMonitoringRateExpectation struct {
	mock   *CurrencyRateUpdaterMock
	params *CurrencyRateUpdaterMockStartMonitoringRateParams

	Counter uint64
}

// CurrencyRateUpdaterMockStartMonitoringRateParams contains parameters of the CurrencyRateUpdater.StartMonitoringRate
type CurrencyRateUpdaterMockStartMonitoringRateParams struct {
	ctx context.Context
}

// Expect sets up expected params for CurrencyRateUpdater.StartMonitoringRate
func (mmStartMonitoringRate *mCurrencyRateUpdaterMockStartMonitoringRate) Expect(ctx context.Context) *mCurrencyRateUpdaterMockStartMonitoringRate {
	if mmStartMonitoringRate.mock.funcStartMonitoringRate != nil {
		mmStartMonitoringRate.mock.t.Fatalf("CurrencyRateUpdaterMock.StartMonitoringRate mock is already set by Set")
	}

	if mmStartMonitoringRate.defaultExpectation == nil {
		mmStartMonitoringRate.defaultExpectation = &CurrencyRateUpdaterMockStartMonitoringRateExpectation{}
	}

	mmStartMonitoringRate.defaultExpectation.params = &CurrencyRateUpdaterMockStartMonitoringRateParams{ctx}
	for _, e := range mmStartMonitoringRate.expectations {
		if minimock.Equal(e.params, mmStartMonitoringRate.defaultExpectation.params) {
			mmStartMonitoringRate.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmStartMonitoringRate.defaultExpectation.params)
		}
	}

	return mmStartMonitoringRate
}

// Inspect accepts an inspector function that has same arguments as the CurrencyRateUpdater.StartMonitoringRate
func (mmStartMonitoringRate *mCurrencyRateUpdaterMockStartMonitoringRate) Inspect(f func(ctx context.Context)) *mCurrencyRateUpdaterMockStartMonitoringRate {
	if mmStartMonitoringRate.mock.inspectFuncStartMonitoringRate != nil {
		mmStartMonitoringRate.mock.t.Fatalf("Inspect function is already set for CurrencyRateUpdaterMock.StartMonitoringRate")
	}

	mmStartMonitoringRate.mock.inspectFuncStartMonitoringRate = f

	return mmStartMonitoringRate
}

// Return sets up results that will be returned by CurrencyRateUpdater.StartMonitoringRate
func (mmStartMonitoringRate *mCurrencyRateUpdaterMockStartMonitoringRate) Return() *CurrencyRateUpdaterMock {
	if mmStartMonitoringRate.mock.funcStartMonitoringRate != nil {
		mmStartMonitoringRate.mock.t.Fatalf("CurrencyRateUpdaterMock.StartMonitoringRate mock is already set by Set")
	}

	if mmStartMonitoringRate.defaultExpectation == nil {
		mmStartMonitoringRate.defaultExpectation = &CurrencyRateUpdaterMockStartMonitoringRateExpectation{mock: mmStartMonitoringRate.mock}
	}

	return mmStartMonitoringRate.mock
}

//Set uses given function f to mock the CurrencyRateUpdater.StartMonitoringRate method
func (mmStartMonitoringRate *mCurrencyRateUpdaterMockStartMonitoringRate) Set(f func(ctx context.Context)) *CurrencyRateUpdaterMock {
	if mmStartMonitoringRate.defaultExpectation != nil {
		mmStartMonitoringRate.mock.t.Fatalf("Default expectation is already set for the CurrencyRateUpdater.StartMonitoringRate method")
	}

	if len(mmStartMonitoringRate.expectations) > 0 {
		mmStartMonitoringRate.mock.t.Fatalf("Some expectations are already set for the CurrencyRateUpdater.StartMonitoringRate method")
	}

	mmStartMonitoringRate.mock.funcStartMonitoringRate = f
	return mmStartMonitoringRate.mock
}

// StartMonitoringRate implements finance_services_currency.CurrencyRateUpdater
func (mmStartMonitoringRate *CurrencyRateUpdaterMock) StartMonitoringRate(ctx context.Context) {
	mm_atomic.AddUint64(&mmStartMonitoringRate.beforeStartMonitoringRateCounter, 1)
	defer mm_atomic.AddUint64(&mmStartMonitoringRate.afterStartMonitoringRateCounter, 1)

	if mmStartMonitoringRate.inspectFuncStartMonitoringRate != nil {
		mmStartMonitoringRate.inspectFuncStartMonitoringRate(ctx)
	}

	mm_params := &CurrencyRateUpdaterMockStartMonitoringRateParams{ctx}

	// Record call args
	mmStartMonitoringRate.StartMonitoringRateMock.mutex.Lock()
	mmStartMonitoringRate.StartMonitoringRateMock.callArgs = append(mmStartMonitoringRate.StartMonitoringRateMock.callArgs, mm_params)
	mmStartMonitoringRate.StartMonitoringRateMock.mutex.Unlock()

	for _, e := range mmStartMonitoringRate.StartMonitoringRateMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return
		}
	}

	if mmStartMonitoringRate.StartMonitoringRateMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmStartMonitoringRate.StartMonitoringRateMock.defaultExpectation.Counter, 1)
		mm_want := mmStartMonitoringRate.StartMonitoringRateMock.defaultExpectation.params
		mm_got := CurrencyRateUpdaterMockStartMonitoringRateParams{ctx}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmStartMonitoringRate.t.Errorf("CurrencyRateUpdaterMock.StartMonitoringRate got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		return

	}
	if mmStartMonitoringRate.funcStartMonitoringRate != nil {
		mmStartMonitoringRate.funcStartMonitoringRate(ctx)
		return
	}
	mmStartMonitoringRate.t.Fatalf("Unexpected call to CurrencyRateUpdaterMock.StartMonitoringRate. %v", ctx)

}

// StartMonitoringRateAfterCounter returns a count of finished CurrencyRateUpdaterMock.StartMonitoringRate invocations
func (mmStartMonitoringRate *CurrencyRateUpdaterMock) StartMonitoringRateAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmStartMonitoringRate.afterStartMonitoringRateCounter)
}

// StartMonitoringRateBeforeCounter returns a count of CurrencyRateUpdaterMock.StartMonitoringRate invocations
func (mmStartMonitoringRate *CurrencyRateUpdaterMock) StartMonitoringRateBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmStartMonitoringRate.beforeStartMonitoringRateCounter)
}

// Calls returns a list of arguments used in each call to CurrencyRateUpdaterMock.StartMonitoringRate.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmStartMonitoringRate *mCurrencyRateUpdaterMockStartMonitoringRate) Calls() []*CurrencyRateUpdaterMockStartMonitoringRateParams {
	mmStartMonitoringRate.mutex.RLock()

	argCopy := make([]*CurrencyRateUpdaterMockStartMonitoringRateParams, len(mmStartMonitoringRate.callArgs))
	copy(argCopy, mmStartMonitoringRate.callArgs)

	mmStartMonitoringRate.mutex.RUnlock()

	return argCopy
}

// MinimockStartMonitoringRateDone returns true if the count of the StartMonitoringRate invocations corresponds
// the number of defined expectations
func (m *CurrencyRateUpdaterMock) MinimockStartMonitoringRateDone() bool {
	for _, e := range m.StartMonitoringRateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.StartMonitoringRateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterStartMonitoringRateCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcStartMonitoringRate != nil && mm_atomic.LoadUint64(&m.afterStartMonitoringRateCounter) < 1 {
		return false
	}
	return true
}

// MinimockStartMonitoringRateInspect logs each unmet expectation
func (m *CurrencyRateUpdaterMock) MinimockStartMonitoringRateInspect() {
	for _, e := range m.StartMonitoringRateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to CurrencyRateUpdaterMock.StartMonitoringRate with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.StartMonitoringRateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterStartMonitoringRateCounter) < 1 {
		if m.StartMonitoringRateMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to CurrencyRateUpdaterMock.StartMonitoringRate")
		} else {
			m.t.Errorf("Expected call to CurrencyRateUpdaterMock.StartMonitoringRate with params: %#v", *m.StartMonitoringRateMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcStartMonitoringRate != nil && mm_atomic.LoadUint64(&m.afterStartMonitoringRateCounter) < 1 {
		m.t.Error("Expected call to CurrencyRateUpdaterMock.StartMonitoringRate")
	}
}

type mCurrencyRateUpdaterMockUpdateRateAsync struct {
	mock               *CurrencyRateUpdaterMock
	defaultExpectation *CurrencyRateUpdaterMockUpdateRateAsyncExpectation
	expectations       []*CurrencyRateUpdaterMockUpdateRateAsyncExpectation

	callArgs []*CurrencyRateUpdaterMockUpdateRateAsyncParams
	mutex    sync.RWMutex
}

// CurrencyRateUpdaterMockUpdateRateAsyncExpectation specifies expectation struct of the CurrencyRateUpdater.UpdateRateAsync
type CurrencyRateUpdaterMockUpdateRateAsyncExpectation struct {
	mock    *CurrencyRateUpdaterMock
	params  *CurrencyRateUpdaterMockUpdateRateAsyncParams
	results *CurrencyRateUpdaterMockUpdateRateAsyncResults
	Counter uint64
}

// CurrencyRateUpdaterMockUpdateRateAsyncParams contains parameters of the CurrencyRateUpdater.UpdateRateAsync
type CurrencyRateUpdaterMockUpdateRateAsyncParams struct {
	ctx context.Context
}

// CurrencyRateUpdaterMockUpdateRateAsyncResults contains results of the CurrencyRateUpdater.UpdateRateAsync
type CurrencyRateUpdaterMockUpdateRateAsyncResults struct {
	ch1 <-chan finance_models.CurrencyRate
}

// Expect sets up expected params for CurrencyRateUpdater.UpdateRateAsync
func (mmUpdateRateAsync *mCurrencyRateUpdaterMockUpdateRateAsync) Expect(ctx context.Context) *mCurrencyRateUpdaterMockUpdateRateAsync {
	if mmUpdateRateAsync.mock.funcUpdateRateAsync != nil {
		mmUpdateRateAsync.mock.t.Fatalf("CurrencyRateUpdaterMock.UpdateRateAsync mock is already set by Set")
	}

	if mmUpdateRateAsync.defaultExpectation == nil {
		mmUpdateRateAsync.defaultExpectation = &CurrencyRateUpdaterMockUpdateRateAsyncExpectation{}
	}

	mmUpdateRateAsync.defaultExpectation.params = &CurrencyRateUpdaterMockUpdateRateAsyncParams{ctx}
	for _, e := range mmUpdateRateAsync.expectations {
		if minimock.Equal(e.params, mmUpdateRateAsync.defaultExpectation.params) {
			mmUpdateRateAsync.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmUpdateRateAsync.defaultExpectation.params)
		}
	}

	return mmUpdateRateAsync
}

// Inspect accepts an inspector function that has same arguments as the CurrencyRateUpdater.UpdateRateAsync
func (mmUpdateRateAsync *mCurrencyRateUpdaterMockUpdateRateAsync) Inspect(f func(ctx context.Context)) *mCurrencyRateUpdaterMockUpdateRateAsync {
	if mmUpdateRateAsync.mock.inspectFuncUpdateRateAsync != nil {
		mmUpdateRateAsync.mock.t.Fatalf("Inspect function is already set for CurrencyRateUpdaterMock.UpdateRateAsync")
	}

	mmUpdateRateAsync.mock.inspectFuncUpdateRateAsync = f

	return mmUpdateRateAsync
}

// Return sets up results that will be returned by CurrencyRateUpdater.UpdateRateAsync
func (mmUpdateRateAsync *mCurrencyRateUpdaterMockUpdateRateAsync) Return(ch1 <-chan finance_models.CurrencyRate) *CurrencyRateUpdaterMock {
	if mmUpdateRateAsync.mock.funcUpdateRateAsync != nil {
		mmUpdateRateAsync.mock.t.Fatalf("CurrencyRateUpdaterMock.UpdateRateAsync mock is already set by Set")
	}

	if mmUpdateRateAsync.defaultExpectation == nil {
		mmUpdateRateAsync.defaultExpectation = &CurrencyRateUpdaterMockUpdateRateAsyncExpectation{mock: mmUpdateRateAsync.mock}
	}
	mmUpdateRateAsync.defaultExpectation.results = &CurrencyRateUpdaterMockUpdateRateAsyncResults{ch1}
	return mmUpdateRateAsync.mock
}

//Set uses given function f to mock the CurrencyRateUpdater.UpdateRateAsync method
func (mmUpdateRateAsync *mCurrencyRateUpdaterMockUpdateRateAsync) Set(f func(ctx context.Context) (ch1 <-chan finance_models.CurrencyRate)) *CurrencyRateUpdaterMock {
	if mmUpdateRateAsync.defaultExpectation != nil {
		mmUpdateRateAsync.mock.t.Fatalf("Default expectation is already set for the CurrencyRateUpdater.UpdateRateAsync method")
	}

	if len(mmUpdateRateAsync.expectations) > 0 {
		mmUpdateRateAsync.mock.t.Fatalf("Some expectations are already set for the CurrencyRateUpdater.UpdateRateAsync method")
	}

	mmUpdateRateAsync.mock.funcUpdateRateAsync = f
	return mmUpdateRateAsync.mock
}

// When sets expectation for the CurrencyRateUpdater.UpdateRateAsync which will trigger the result defined by the following
// Then helper
func (mmUpdateRateAsync *mCurrencyRateUpdaterMockUpdateRateAsync) When(ctx context.Context) *CurrencyRateUpdaterMockUpdateRateAsyncExpectation {
	if mmUpdateRateAsync.mock.funcUpdateRateAsync != nil {
		mmUpdateRateAsync.mock.t.Fatalf("CurrencyRateUpdaterMock.UpdateRateAsync mock is already set by Set")
	}

	expectation := &CurrencyRateUpdaterMockUpdateRateAsyncExpectation{
		mock:   mmUpdateRateAsync.mock,
		params: &CurrencyRateUpdaterMockUpdateRateAsyncParams{ctx},
	}
	mmUpdateRateAsync.expectations = append(mmUpdateRateAsync.expectations, expectation)
	return expectation
}

// Then sets up CurrencyRateUpdater.UpdateRateAsync return parameters for the expectation previously defined by the When method
func (e *CurrencyRateUpdaterMockUpdateRateAsyncExpectation) Then(ch1 <-chan finance_models.CurrencyRate) *CurrencyRateUpdaterMock {
	e.results = &CurrencyRateUpdaterMockUpdateRateAsyncResults{ch1}
	return e.mock
}

// UpdateRateAsync implements finance_services_currency.CurrencyRateUpdater
func (mmUpdateRateAsync *CurrencyRateUpdaterMock) UpdateRateAsync(ctx context.Context) (ch1 <-chan finance_models.CurrencyRate) {
	mm_atomic.AddUint64(&mmUpdateRateAsync.beforeUpdateRateAsyncCounter, 1)
	defer mm_atomic.AddUint64(&mmUpdateRateAsync.afterUpdateRateAsyncCounter, 1)

	if mmUpdateRateAsync.inspectFuncUpdateRateAsync != nil {
		mmUpdateRateAsync.inspectFuncUpdateRateAsync(ctx)
	}

	mm_params := &CurrencyRateUpdaterMockUpdateRateAsyncParams{ctx}

	// Record call args
	mmUpdateRateAsync.UpdateRateAsyncMock.mutex.Lock()
	mmUpdateRateAsync.UpdateRateAsyncMock.callArgs = append(mmUpdateRateAsync.UpdateRateAsyncMock.callArgs, mm_params)
	mmUpdateRateAsync.UpdateRateAsyncMock.mutex.Unlock()

	for _, e := range mmUpdateRateAsync.UpdateRateAsyncMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.ch1
		}
	}

	if mmUpdateRateAsync.UpdateRateAsyncMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmUpdateRateAsync.UpdateRateAsyncMock.defaultExpectation.Counter, 1)
		mm_want := mmUpdateRateAsync.UpdateRateAsyncMock.defaultExpectation.params
		mm_got := CurrencyRateUpdaterMockUpdateRateAsyncParams{ctx}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmUpdateRateAsync.t.Errorf("CurrencyRateUpdaterMock.UpdateRateAsync got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmUpdateRateAsync.UpdateRateAsyncMock.defaultExpectation.results
		if mm_results == nil {
			mmUpdateRateAsync.t.Fatal("No results are set for the CurrencyRateUpdaterMock.UpdateRateAsync")
		}
		return (*mm_results).ch1
	}
	if mmUpdateRateAsync.funcUpdateRateAsync != nil {
		return mmUpdateRateAsync.funcUpdateRateAsync(ctx)
	}
	mmUpdateRateAsync.t.Fatalf("Unexpected call to CurrencyRateUpdaterMock.UpdateRateAsync. %v", ctx)
	return
}

// UpdateRateAsyncAfterCounter returns a count of finished CurrencyRateUpdaterMock.UpdateRateAsync invocations
func (mmUpdateRateAsync *CurrencyRateUpdaterMock) UpdateRateAsyncAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmUpdateRateAsync.afterUpdateRateAsyncCounter)
}

// UpdateRateAsyncBeforeCounter returns a count of CurrencyRateUpdaterMock.UpdateRateAsync invocations
func (mmUpdateRateAsync *CurrencyRateUpdaterMock) UpdateRateAsyncBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmUpdateRateAsync.beforeUpdateRateAsyncCounter)
}

// Calls returns a list of arguments used in each call to CurrencyRateUpdaterMock.UpdateRateAsync.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmUpdateRateAsync *mCurrencyRateUpdaterMockUpdateRateAsync) Calls() []*CurrencyRateUpdaterMockUpdateRateAsyncParams {
	mmUpdateRateAsync.mutex.RLock()

	argCopy := make([]*CurrencyRateUpdaterMockUpdateRateAsyncParams, len(mmUpdateRateAsync.callArgs))
	copy(argCopy, mmUpdateRateAsync.callArgs)

	mmUpdateRateAsync.mutex.RUnlock()

	return argCopy
}

// MinimockUpdateRateAsyncDone returns true if the count of the UpdateRateAsync invocations corresponds
// the number of defined expectations
func (m *CurrencyRateUpdaterMock) MinimockUpdateRateAsyncDone() bool {
	for _, e := range m.UpdateRateAsyncMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.UpdateRateAsyncMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterUpdateRateAsyncCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcUpdateRateAsync != nil && mm_atomic.LoadUint64(&m.afterUpdateRateAsyncCounter) < 1 {
		return false
	}
	return true
}

// MinimockUpdateRateAsyncInspect logs each unmet expectation
func (m *CurrencyRateUpdaterMock) MinimockUpdateRateAsyncInspect() {
	for _, e := range m.UpdateRateAsyncMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to CurrencyRateUpdaterMock.UpdateRateAsync with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.UpdateRateAsyncMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterUpdateRateAsyncCounter) < 1 {
		if m.UpdateRateAsyncMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to CurrencyRateUpdaterMock.UpdateRateAsync")
		} else {
			m.t.Errorf("Expected call to CurrencyRateUpdaterMock.UpdateRateAsync with params: %#v", *m.UpdateRateAsyncMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcUpdateRateAsync != nil && mm_atomic.LoadUint64(&m.afterUpdateRateAsyncCounter) < 1 {
		m.t.Error("Expected call to CurrencyRateUpdaterMock.UpdateRateAsync")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *CurrencyRateUpdaterMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockGetCurrencyInspect()

		m.MinimockStartMonitoringRateInspect()

		m.MinimockUpdateRateAsyncInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *CurrencyRateUpdaterMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *CurrencyRateUpdaterMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockGetCurrencyDone() &&
		m.MinimockStartMonitoringRateDone() &&
		m.MinimockUpdateRateAsyncDone()
}
