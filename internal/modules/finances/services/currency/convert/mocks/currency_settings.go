package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i github.com/shav/telegram-bot/internal/modules/finances/services/currency/convert.currencySettings -o ./mocks\currency_settings.go -n CurrencySettingsMock

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	tr "github.com/shav/telegram-bot/internal/common/transactions"
	finance_models "github.com/shav/telegram-bot/internal/modules/finances/domain/models"
)

// CurrencySettingsMock implements finance_services_currency.currencySettings
type CurrencySettingsMock struct {
	t minimock.Tester

	funcGetCurrency          func(ctx context.Context, ts tr.Transaction, userId int64) (c2 finance_models.Currency, err error)
	inspectFuncGetCurrency   func(ctx context.Context, ts tr.Transaction, userId int64)
	afterGetCurrencyCounter  uint64
	beforeGetCurrencyCounter uint64
	GetCurrencyMock          mCurrencySettingsMockGetCurrency

	funcGetDefaultCurrency          func() (c1 finance_models.Currency)
	inspectFuncGetDefaultCurrency   func()
	afterGetDefaultCurrencyCounter  uint64
	beforeGetDefaultCurrencyCounter uint64
	GetDefaultCurrencyMock          mCurrencySettingsMockGetDefaultCurrency
}

// NewCurrencySettingsMock returns a mock for finance_services_currency.currencySettings
func NewCurrencySettingsMock(t minimock.Tester) *CurrencySettingsMock {
	m := &CurrencySettingsMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.GetCurrencyMock = mCurrencySettingsMockGetCurrency{mock: m}
	m.GetCurrencyMock.callArgs = []*CurrencySettingsMockGetCurrencyParams{}

	m.GetDefaultCurrencyMock = mCurrencySettingsMockGetDefaultCurrency{mock: m}

	return m
}

type mCurrencySettingsMockGetCurrency struct {
	mock               *CurrencySettingsMock
	defaultExpectation *CurrencySettingsMockGetCurrencyExpectation
	expectations       []*CurrencySettingsMockGetCurrencyExpectation

	callArgs []*CurrencySettingsMockGetCurrencyParams
	mutex    sync.RWMutex
}

// CurrencySettingsMockGetCurrencyExpectation specifies expectation struct of the currencySettings.GetCurrency
type CurrencySettingsMockGetCurrencyExpectation struct {
	mock    *CurrencySettingsMock
	params  *CurrencySettingsMockGetCurrencyParams
	results *CurrencySettingsMockGetCurrencyResults
	Counter uint64
}

// CurrencySettingsMockGetCurrencyParams contains parameters of the currencySettings.GetCurrency
type CurrencySettingsMockGetCurrencyParams struct {
	ctx    context.Context
	ts     tr.Transaction
	userId int64
}

// CurrencySettingsMockGetCurrencyResults contains results of the currencySettings.GetCurrency
type CurrencySettingsMockGetCurrencyResults struct {
	c2  finance_models.Currency
	err error
}

// Expect sets up expected params for currencySettings.GetCurrency
func (mmGetCurrency *mCurrencySettingsMockGetCurrency) Expect(ctx context.Context, ts tr.Transaction, userId int64) *mCurrencySettingsMockGetCurrency {
	if mmGetCurrency.mock.funcGetCurrency != nil {
		mmGetCurrency.mock.t.Fatalf("CurrencySettingsMock.GetCurrency mock is already set by Set")
	}

	if mmGetCurrency.defaultExpectation == nil {
		mmGetCurrency.defaultExpectation = &CurrencySettingsMockGetCurrencyExpectation{}
	}

	mmGetCurrency.defaultExpectation.params = &CurrencySettingsMockGetCurrencyParams{ctx, ts, userId}
	for _, e := range mmGetCurrency.expectations {
		if minimock.Equal(e.params, mmGetCurrency.defaultExpectation.params) {
			mmGetCurrency.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetCurrency.defaultExpectation.params)
		}
	}

	return mmGetCurrency
}

// Inspect accepts an inspector function that has same arguments as the currencySettings.GetCurrency
func (mmGetCurrency *mCurrencySettingsMockGetCurrency) Inspect(f func(ctx context.Context, ts tr.Transaction, userId int64)) *mCurrencySettingsMockGetCurrency {
	if mmGetCurrency.mock.inspectFuncGetCurrency != nil {
		mmGetCurrency.mock.t.Fatalf("Inspect function is already set for CurrencySettingsMock.GetCurrency")
	}

	mmGetCurrency.mock.inspectFuncGetCurrency = f

	return mmGetCurrency
}

// Return sets up results that will be returned by currencySettings.GetCurrency
func (mmGetCurrency *mCurrencySettingsMockGetCurrency) Return(c2 finance_models.Currency, err error) *CurrencySettingsMock {
	if mmGetCurrency.mock.funcGetCurrency != nil {
		mmGetCurrency.mock.t.Fatalf("CurrencySettingsMock.GetCurrency mock is already set by Set")
	}

	if mmGetCurrency.defaultExpectation == nil {
		mmGetCurrency.defaultExpectation = &CurrencySettingsMockGetCurrencyExpectation{mock: mmGetCurrency.mock}
	}
	mmGetCurrency.defaultExpectation.results = &CurrencySettingsMockGetCurrencyResults{c2, err}
	return mmGetCurrency.mock
}

//Set uses given function f to mock the currencySettings.GetCurrency method
func (mmGetCurrency *mCurrencySettingsMockGetCurrency) Set(f func(ctx context.Context, ts tr.Transaction, userId int64) (c2 finance_models.Currency, err error)) *CurrencySettingsMock {
	if mmGetCurrency.defaultExpectation != nil {
		mmGetCurrency.mock.t.Fatalf("Default expectation is already set for the currencySettings.GetCurrency method")
	}

	if len(mmGetCurrency.expectations) > 0 {
		mmGetCurrency.mock.t.Fatalf("Some expectations are already set for the currencySettings.GetCurrency method")
	}

	mmGetCurrency.mock.funcGetCurrency = f
	return mmGetCurrency.mock
}

// When sets expectation for the currencySettings.GetCurrency which will trigger the result defined by the following
// Then helper
func (mmGetCurrency *mCurrencySettingsMockGetCurrency) When(ctx context.Context, ts tr.Transaction, userId int64) *CurrencySettingsMockGetCurrencyExpectation {
	if mmGetCurrency.mock.funcGetCurrency != nil {
		mmGetCurrency.mock.t.Fatalf("CurrencySettingsMock.GetCurrency mock is already set by Set")
	}

	expectation := &CurrencySettingsMockGetCurrencyExpectation{
		mock:   mmGetCurrency.mock,
		params: &CurrencySettingsMockGetCurrencyParams{ctx, ts, userId},
	}
	mmGetCurrency.expectations = append(mmGetCurrency.expectations, expectation)
	return expectation
}

// Then sets up currencySettings.GetCurrency return parameters for the expectation previously defined by the When method
func (e *CurrencySettingsMockGetCurrencyExpectation) Then(c2 finance_models.Currency, err error) *CurrencySettingsMock {
	e.results = &CurrencySettingsMockGetCurrencyResults{c2, err}
	return e.mock
}

// GetCurrency implements finance_services_currency.currencySettings
func (mmGetCurrency *CurrencySettingsMock) GetCurrency(ctx context.Context, ts tr.Transaction, userId int64) (c2 finance_models.Currency, err error) {
	mm_atomic.AddUint64(&mmGetCurrency.beforeGetCurrencyCounter, 1)
	defer mm_atomic.AddUint64(&mmGetCurrency.afterGetCurrencyCounter, 1)

	if mmGetCurrency.inspectFuncGetCurrency != nil {
		mmGetCurrency.inspectFuncGetCurrency(ctx, ts, userId)
	}

	mm_params := &CurrencySettingsMockGetCurrencyParams{ctx, ts, userId}

	// Record call args
	mmGetCurrency.GetCurrencyMock.mutex.Lock()
	mmGetCurrency.GetCurrencyMock.callArgs = append(mmGetCurrency.GetCurrencyMock.callArgs, mm_params)
	mmGetCurrency.GetCurrencyMock.mutex.Unlock()

	for _, e := range mmGetCurrency.GetCurrencyMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.c2, e.results.err
		}
	}

	if mmGetCurrency.GetCurrencyMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetCurrency.GetCurrencyMock.defaultExpectation.Counter, 1)
		mm_want := mmGetCurrency.GetCurrencyMock.defaultExpectation.params
		mm_got := CurrencySettingsMockGetCurrencyParams{ctx, ts, userId}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetCurrency.t.Errorf("CurrencySettingsMock.GetCurrency got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetCurrency.GetCurrencyMock.defaultExpectation.results
		if mm_results == nil {
			mmGetCurrency.t.Fatal("No results are set for the CurrencySettingsMock.GetCurrency")
		}
		return (*mm_results).c2, (*mm_results).err
	}
	if mmGetCurrency.funcGetCurrency != nil {
		return mmGetCurrency.funcGetCurrency(ctx, ts, userId)
	}
	mmGetCurrency.t.Fatalf("Unexpected call to CurrencySettingsMock.GetCurrency. %v %v %v", ctx, ts, userId)
	return
}

// GetCurrencyAfterCounter returns a count of finished CurrencySettingsMock.GetCurrency invocations
func (mmGetCurrency *CurrencySettingsMock) GetCurrencyAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetCurrency.afterGetCurrencyCounter)
}

// GetCurrencyBeforeCounter returns a count of CurrencySettingsMock.GetCurrency invocations
func (mmGetCurrency *CurrencySettingsMock) GetCurrencyBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetCurrency.beforeGetCurrencyCounter)
}

// Calls returns a list of arguments used in each call to CurrencySettingsMock.GetCurrency.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetCurrency *mCurrencySettingsMockGetCurrency) Calls() []*CurrencySettingsMockGetCurrencyParams {
	mmGetCurrency.mutex.RLock()

	argCopy := make([]*CurrencySettingsMockGetCurrencyParams, len(mmGetCurrency.callArgs))
	copy(argCopy, mmGetCurrency.callArgs)

	mmGetCurrency.mutex.RUnlock()

	return argCopy
}

// MinimockGetCurrencyDone returns true if the count of the GetCurrency invocations corresponds
// the number of defined expectations
func (m *CurrencySettingsMock) MinimockGetCurrencyDone() bool {
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
func (m *CurrencySettingsMock) MinimockGetCurrencyInspect() {
	for _, e := range m.GetCurrencyMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to CurrencySettingsMock.GetCurrency with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetCurrencyMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCurrencyCounter) < 1 {
		if m.GetCurrencyMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to CurrencySettingsMock.GetCurrency")
		} else {
			m.t.Errorf("Expected call to CurrencySettingsMock.GetCurrency with params: %#v", *m.GetCurrencyMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetCurrency != nil && mm_atomic.LoadUint64(&m.afterGetCurrencyCounter) < 1 {
		m.t.Error("Expected call to CurrencySettingsMock.GetCurrency")
	}
}

type mCurrencySettingsMockGetDefaultCurrency struct {
	mock               *CurrencySettingsMock
	defaultExpectation *CurrencySettingsMockGetDefaultCurrencyExpectation
	expectations       []*CurrencySettingsMockGetDefaultCurrencyExpectation
}

// CurrencySettingsMockGetDefaultCurrencyExpectation specifies expectation struct of the currencySettings.GetDefaultCurrency
type CurrencySettingsMockGetDefaultCurrencyExpectation struct {
	mock *CurrencySettingsMock

	results *CurrencySettingsMockGetDefaultCurrencyResults
	Counter uint64
}

// CurrencySettingsMockGetDefaultCurrencyResults contains results of the currencySettings.GetDefaultCurrency
type CurrencySettingsMockGetDefaultCurrencyResults struct {
	c1 finance_models.Currency
}

// Expect sets up expected params for currencySettings.GetDefaultCurrency
func (mmGetDefaultCurrency *mCurrencySettingsMockGetDefaultCurrency) Expect() *mCurrencySettingsMockGetDefaultCurrency {
	if mmGetDefaultCurrency.mock.funcGetDefaultCurrency != nil {
		mmGetDefaultCurrency.mock.t.Fatalf("CurrencySettingsMock.GetDefaultCurrency mock is already set by Set")
	}

	if mmGetDefaultCurrency.defaultExpectation == nil {
		mmGetDefaultCurrency.defaultExpectation = &CurrencySettingsMockGetDefaultCurrencyExpectation{}
	}

	return mmGetDefaultCurrency
}

// Inspect accepts an inspector function that has same arguments as the currencySettings.GetDefaultCurrency
func (mmGetDefaultCurrency *mCurrencySettingsMockGetDefaultCurrency) Inspect(f func()) *mCurrencySettingsMockGetDefaultCurrency {
	if mmGetDefaultCurrency.mock.inspectFuncGetDefaultCurrency != nil {
		mmGetDefaultCurrency.mock.t.Fatalf("Inspect function is already set for CurrencySettingsMock.GetDefaultCurrency")
	}

	mmGetDefaultCurrency.mock.inspectFuncGetDefaultCurrency = f

	return mmGetDefaultCurrency
}

// Return sets up results that will be returned by currencySettings.GetDefaultCurrency
func (mmGetDefaultCurrency *mCurrencySettingsMockGetDefaultCurrency) Return(c1 finance_models.Currency) *CurrencySettingsMock {
	if mmGetDefaultCurrency.mock.funcGetDefaultCurrency != nil {
		mmGetDefaultCurrency.mock.t.Fatalf("CurrencySettingsMock.GetDefaultCurrency mock is already set by Set")
	}

	if mmGetDefaultCurrency.defaultExpectation == nil {
		mmGetDefaultCurrency.defaultExpectation = &CurrencySettingsMockGetDefaultCurrencyExpectation{mock: mmGetDefaultCurrency.mock}
	}
	mmGetDefaultCurrency.defaultExpectation.results = &CurrencySettingsMockGetDefaultCurrencyResults{c1}
	return mmGetDefaultCurrency.mock
}

//Set uses given function f to mock the currencySettings.GetDefaultCurrency method
func (mmGetDefaultCurrency *mCurrencySettingsMockGetDefaultCurrency) Set(f func() (c1 finance_models.Currency)) *CurrencySettingsMock {
	if mmGetDefaultCurrency.defaultExpectation != nil {
		mmGetDefaultCurrency.mock.t.Fatalf("Default expectation is already set for the currencySettings.GetDefaultCurrency method")
	}

	if len(mmGetDefaultCurrency.expectations) > 0 {
		mmGetDefaultCurrency.mock.t.Fatalf("Some expectations are already set for the currencySettings.GetDefaultCurrency method")
	}

	mmGetDefaultCurrency.mock.funcGetDefaultCurrency = f
	return mmGetDefaultCurrency.mock
}

// GetDefaultCurrency implements finance_services_currency.currencySettings
func (mmGetDefaultCurrency *CurrencySettingsMock) GetDefaultCurrency() (c1 finance_models.Currency) {
	mm_atomic.AddUint64(&mmGetDefaultCurrency.beforeGetDefaultCurrencyCounter, 1)
	defer mm_atomic.AddUint64(&mmGetDefaultCurrency.afterGetDefaultCurrencyCounter, 1)

	if mmGetDefaultCurrency.inspectFuncGetDefaultCurrency != nil {
		mmGetDefaultCurrency.inspectFuncGetDefaultCurrency()
	}

	if mmGetDefaultCurrency.GetDefaultCurrencyMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetDefaultCurrency.GetDefaultCurrencyMock.defaultExpectation.Counter, 1)

		mm_results := mmGetDefaultCurrency.GetDefaultCurrencyMock.defaultExpectation.results
		if mm_results == nil {
			mmGetDefaultCurrency.t.Fatal("No results are set for the CurrencySettingsMock.GetDefaultCurrency")
		}
		return (*mm_results).c1
	}
	if mmGetDefaultCurrency.funcGetDefaultCurrency != nil {
		return mmGetDefaultCurrency.funcGetDefaultCurrency()
	}
	mmGetDefaultCurrency.t.Fatalf("Unexpected call to CurrencySettingsMock.GetDefaultCurrency.")
	return
}

// GetDefaultCurrencyAfterCounter returns a count of finished CurrencySettingsMock.GetDefaultCurrency invocations
func (mmGetDefaultCurrency *CurrencySettingsMock) GetDefaultCurrencyAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetDefaultCurrency.afterGetDefaultCurrencyCounter)
}

// GetDefaultCurrencyBeforeCounter returns a count of CurrencySettingsMock.GetDefaultCurrency invocations
func (mmGetDefaultCurrency *CurrencySettingsMock) GetDefaultCurrencyBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetDefaultCurrency.beforeGetDefaultCurrencyCounter)
}

// MinimockGetDefaultCurrencyDone returns true if the count of the GetDefaultCurrency invocations corresponds
// the number of defined expectations
func (m *CurrencySettingsMock) MinimockGetDefaultCurrencyDone() bool {
	for _, e := range m.GetDefaultCurrencyMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetDefaultCurrencyMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetDefaultCurrencyCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetDefaultCurrency != nil && mm_atomic.LoadUint64(&m.afterGetDefaultCurrencyCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetDefaultCurrencyInspect logs each unmet expectation
func (m *CurrencySettingsMock) MinimockGetDefaultCurrencyInspect() {
	for _, e := range m.GetDefaultCurrencyMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to CurrencySettingsMock.GetDefaultCurrency")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetDefaultCurrencyMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetDefaultCurrencyCounter) < 1 {
		m.t.Error("Expected call to CurrencySettingsMock.GetDefaultCurrency")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetDefaultCurrency != nil && mm_atomic.LoadUint64(&m.afterGetDefaultCurrencyCounter) < 1 {
		m.t.Error("Expected call to CurrencySettingsMock.GetDefaultCurrency")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *CurrencySettingsMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockGetCurrencyInspect()

		m.MinimockGetDefaultCurrencyInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *CurrencySettingsMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *CurrencySettingsMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockGetCurrencyDone() &&
		m.MinimockGetDefaultCurrencyDone()
}
