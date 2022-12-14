package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i github.com/shav/telegram-bot/internal/modules/finances/domain/reports/spendings.currencyConverter -o ./mocks\currency_converter.go -n CurrencyConverterMock

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	finance_models "github.com/shav/telegram-bot/internal/modules/finances/domain/models"
)

// CurrencyConverterMock implements finance_reports.currencyConverter
type CurrencyConverterMock struct {
	t minimock.Tester

	funcConvertSpendingsTableToUserCurrency          func(ctx context.Context, userId int64, spendingsTable finance_models.SpendingsByCategoryTable) (s1 finance_models.SpendingsByCategoryTable, err error)
	inspectFuncConvertSpendingsTableToUserCurrency   func(ctx context.Context, userId int64, spendingsTable finance_models.SpendingsByCategoryTable)
	afterConvertSpendingsTableToUserCurrencyCounter  uint64
	beforeConvertSpendingsTableToUserCurrencyCounter uint64
	ConvertSpendingsTableToUserCurrencyMock          mCurrencyConverterMockConvertSpendingsTableToUserCurrency
}

// NewCurrencyConverterMock returns a mock for finance_reports.currencyConverter
func NewCurrencyConverterMock(t minimock.Tester) *CurrencyConverterMock {
	m := &CurrencyConverterMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.ConvertSpendingsTableToUserCurrencyMock = mCurrencyConverterMockConvertSpendingsTableToUserCurrency{mock: m}
	m.ConvertSpendingsTableToUserCurrencyMock.callArgs = []*CurrencyConverterMockConvertSpendingsTableToUserCurrencyParams{}

	return m
}

type mCurrencyConverterMockConvertSpendingsTableToUserCurrency struct {
	mock               *CurrencyConverterMock
	defaultExpectation *CurrencyConverterMockConvertSpendingsTableToUserCurrencyExpectation
	expectations       []*CurrencyConverterMockConvertSpendingsTableToUserCurrencyExpectation

	callArgs []*CurrencyConverterMockConvertSpendingsTableToUserCurrencyParams
	mutex    sync.RWMutex
}

// CurrencyConverterMockConvertSpendingsTableToUserCurrencyExpectation specifies expectation struct of the currencyConverter.ConvertSpendingsTableToUserCurrency
type CurrencyConverterMockConvertSpendingsTableToUserCurrencyExpectation struct {
	mock    *CurrencyConverterMock
	params  *CurrencyConverterMockConvertSpendingsTableToUserCurrencyParams
	results *CurrencyConverterMockConvertSpendingsTableToUserCurrencyResults
	Counter uint64
}

// CurrencyConverterMockConvertSpendingsTableToUserCurrencyParams contains parameters of the currencyConverter.ConvertSpendingsTableToUserCurrency
type CurrencyConverterMockConvertSpendingsTableToUserCurrencyParams struct {
	ctx            context.Context
	userId         int64
	spendingsTable finance_models.SpendingsByCategoryTable
}

// CurrencyConverterMockConvertSpendingsTableToUserCurrencyResults contains results of the currencyConverter.ConvertSpendingsTableToUserCurrency
type CurrencyConverterMockConvertSpendingsTableToUserCurrencyResults struct {
	s1  finance_models.SpendingsByCategoryTable
	err error
}

// Expect sets up expected params for currencyConverter.ConvertSpendingsTableToUserCurrency
func (mmConvertSpendingsTableToUserCurrency *mCurrencyConverterMockConvertSpendingsTableToUserCurrency) Expect(ctx context.Context, userId int64, spendingsTable finance_models.SpendingsByCategoryTable) *mCurrencyConverterMockConvertSpendingsTableToUserCurrency {
	if mmConvertSpendingsTableToUserCurrency.mock.funcConvertSpendingsTableToUserCurrency != nil {
		mmConvertSpendingsTableToUserCurrency.mock.t.Fatalf("CurrencyConverterMock.ConvertSpendingsTableToUserCurrency mock is already set by Set")
	}

	if mmConvertSpendingsTableToUserCurrency.defaultExpectation == nil {
		mmConvertSpendingsTableToUserCurrency.defaultExpectation = &CurrencyConverterMockConvertSpendingsTableToUserCurrencyExpectation{}
	}

	mmConvertSpendingsTableToUserCurrency.defaultExpectation.params = &CurrencyConverterMockConvertSpendingsTableToUserCurrencyParams{ctx, userId, spendingsTable}
	for _, e := range mmConvertSpendingsTableToUserCurrency.expectations {
		if minimock.Equal(e.params, mmConvertSpendingsTableToUserCurrency.defaultExpectation.params) {
			mmConvertSpendingsTableToUserCurrency.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmConvertSpendingsTableToUserCurrency.defaultExpectation.params)
		}
	}

	return mmConvertSpendingsTableToUserCurrency
}

// Inspect accepts an inspector function that has same arguments as the currencyConverter.ConvertSpendingsTableToUserCurrency
func (mmConvertSpendingsTableToUserCurrency *mCurrencyConverterMockConvertSpendingsTableToUserCurrency) Inspect(f func(ctx context.Context, userId int64, spendingsTable finance_models.SpendingsByCategoryTable)) *mCurrencyConverterMockConvertSpendingsTableToUserCurrency {
	if mmConvertSpendingsTableToUserCurrency.mock.inspectFuncConvertSpendingsTableToUserCurrency != nil {
		mmConvertSpendingsTableToUserCurrency.mock.t.Fatalf("Inspect function is already set for CurrencyConverterMock.ConvertSpendingsTableToUserCurrency")
	}

	mmConvertSpendingsTableToUserCurrency.mock.inspectFuncConvertSpendingsTableToUserCurrency = f

	return mmConvertSpendingsTableToUserCurrency
}

// Return sets up results that will be returned by currencyConverter.ConvertSpendingsTableToUserCurrency
func (mmConvertSpendingsTableToUserCurrency *mCurrencyConverterMockConvertSpendingsTableToUserCurrency) Return(s1 finance_models.SpendingsByCategoryTable, err error) *CurrencyConverterMock {
	if mmConvertSpendingsTableToUserCurrency.mock.funcConvertSpendingsTableToUserCurrency != nil {
		mmConvertSpendingsTableToUserCurrency.mock.t.Fatalf("CurrencyConverterMock.ConvertSpendingsTableToUserCurrency mock is already set by Set")
	}

	if mmConvertSpendingsTableToUserCurrency.defaultExpectation == nil {
		mmConvertSpendingsTableToUserCurrency.defaultExpectation = &CurrencyConverterMockConvertSpendingsTableToUserCurrencyExpectation{mock: mmConvertSpendingsTableToUserCurrency.mock}
	}
	mmConvertSpendingsTableToUserCurrency.defaultExpectation.results = &CurrencyConverterMockConvertSpendingsTableToUserCurrencyResults{s1, err}
	return mmConvertSpendingsTableToUserCurrency.mock
}

//Set uses given function f to mock the currencyConverter.ConvertSpendingsTableToUserCurrency method
func (mmConvertSpendingsTableToUserCurrency *mCurrencyConverterMockConvertSpendingsTableToUserCurrency) Set(f func(ctx context.Context, userId int64, spendingsTable finance_models.SpendingsByCategoryTable) (s1 finance_models.SpendingsByCategoryTable, err error)) *CurrencyConverterMock {
	if mmConvertSpendingsTableToUserCurrency.defaultExpectation != nil {
		mmConvertSpendingsTableToUserCurrency.mock.t.Fatalf("Default expectation is already set for the currencyConverter.ConvertSpendingsTableToUserCurrency method")
	}

	if len(mmConvertSpendingsTableToUserCurrency.expectations) > 0 {
		mmConvertSpendingsTableToUserCurrency.mock.t.Fatalf("Some expectations are already set for the currencyConverter.ConvertSpendingsTableToUserCurrency method")
	}

	mmConvertSpendingsTableToUserCurrency.mock.funcConvertSpendingsTableToUserCurrency = f
	return mmConvertSpendingsTableToUserCurrency.mock
}

// When sets expectation for the currencyConverter.ConvertSpendingsTableToUserCurrency which will trigger the result defined by the following
// Then helper
func (mmConvertSpendingsTableToUserCurrency *mCurrencyConverterMockConvertSpendingsTableToUserCurrency) When(ctx context.Context, userId int64, spendingsTable finance_models.SpendingsByCategoryTable) *CurrencyConverterMockConvertSpendingsTableToUserCurrencyExpectation {
	if mmConvertSpendingsTableToUserCurrency.mock.funcConvertSpendingsTableToUserCurrency != nil {
		mmConvertSpendingsTableToUserCurrency.mock.t.Fatalf("CurrencyConverterMock.ConvertSpendingsTableToUserCurrency mock is already set by Set")
	}

	expectation := &CurrencyConverterMockConvertSpendingsTableToUserCurrencyExpectation{
		mock:   mmConvertSpendingsTableToUserCurrency.mock,
		params: &CurrencyConverterMockConvertSpendingsTableToUserCurrencyParams{ctx, userId, spendingsTable},
	}
	mmConvertSpendingsTableToUserCurrency.expectations = append(mmConvertSpendingsTableToUserCurrency.expectations, expectation)
	return expectation
}

// Then sets up currencyConverter.ConvertSpendingsTableToUserCurrency return parameters for the expectation previously defined by the When method
func (e *CurrencyConverterMockConvertSpendingsTableToUserCurrencyExpectation) Then(s1 finance_models.SpendingsByCategoryTable, err error) *CurrencyConverterMock {
	e.results = &CurrencyConverterMockConvertSpendingsTableToUserCurrencyResults{s1, err}
	return e.mock
}

// ConvertSpendingsTableToUserCurrency implements finance_reports.currencyConverter
func (mmConvertSpendingsTableToUserCurrency *CurrencyConverterMock) ConvertSpendingsTableToUserCurrency(ctx context.Context, userId int64, spendingsTable finance_models.SpendingsByCategoryTable) (s1 finance_models.SpendingsByCategoryTable, err error) {
	mm_atomic.AddUint64(&mmConvertSpendingsTableToUserCurrency.beforeConvertSpendingsTableToUserCurrencyCounter, 1)
	defer mm_atomic.AddUint64(&mmConvertSpendingsTableToUserCurrency.afterConvertSpendingsTableToUserCurrencyCounter, 1)

	if mmConvertSpendingsTableToUserCurrency.inspectFuncConvertSpendingsTableToUserCurrency != nil {
		mmConvertSpendingsTableToUserCurrency.inspectFuncConvertSpendingsTableToUserCurrency(ctx, userId, spendingsTable)
	}

	mm_params := &CurrencyConverterMockConvertSpendingsTableToUserCurrencyParams{ctx, userId, spendingsTable}

	// Record call args
	mmConvertSpendingsTableToUserCurrency.ConvertSpendingsTableToUserCurrencyMock.mutex.Lock()
	mmConvertSpendingsTableToUserCurrency.ConvertSpendingsTableToUserCurrencyMock.callArgs = append(mmConvertSpendingsTableToUserCurrency.ConvertSpendingsTableToUserCurrencyMock.callArgs, mm_params)
	mmConvertSpendingsTableToUserCurrency.ConvertSpendingsTableToUserCurrencyMock.mutex.Unlock()

	for _, e := range mmConvertSpendingsTableToUserCurrency.ConvertSpendingsTableToUserCurrencyMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.s1, e.results.err
		}
	}

	if mmConvertSpendingsTableToUserCurrency.ConvertSpendingsTableToUserCurrencyMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmConvertSpendingsTableToUserCurrency.ConvertSpendingsTableToUserCurrencyMock.defaultExpectation.Counter, 1)
		mm_want := mmConvertSpendingsTableToUserCurrency.ConvertSpendingsTableToUserCurrencyMock.defaultExpectation.params
		mm_got := CurrencyConverterMockConvertSpendingsTableToUserCurrencyParams{ctx, userId, spendingsTable}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmConvertSpendingsTableToUserCurrency.t.Errorf("CurrencyConverterMock.ConvertSpendingsTableToUserCurrency got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmConvertSpendingsTableToUserCurrency.ConvertSpendingsTableToUserCurrencyMock.defaultExpectation.results
		if mm_results == nil {
			mmConvertSpendingsTableToUserCurrency.t.Fatal("No results are set for the CurrencyConverterMock.ConvertSpendingsTableToUserCurrency")
		}
		return (*mm_results).s1, (*mm_results).err
	}
	if mmConvertSpendingsTableToUserCurrency.funcConvertSpendingsTableToUserCurrency != nil {
		return mmConvertSpendingsTableToUserCurrency.funcConvertSpendingsTableToUserCurrency(ctx, userId, spendingsTable)
	}
	mmConvertSpendingsTableToUserCurrency.t.Fatalf("Unexpected call to CurrencyConverterMock.ConvertSpendingsTableToUserCurrency. %v %v %v", ctx, userId, spendingsTable)
	return
}

// ConvertSpendingsTableToUserCurrencyAfterCounter returns a count of finished CurrencyConverterMock.ConvertSpendingsTableToUserCurrency invocations
func (mmConvertSpendingsTableToUserCurrency *CurrencyConverterMock) ConvertSpendingsTableToUserCurrencyAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmConvertSpendingsTableToUserCurrency.afterConvertSpendingsTableToUserCurrencyCounter)
}

// ConvertSpendingsTableToUserCurrencyBeforeCounter returns a count of CurrencyConverterMock.ConvertSpendingsTableToUserCurrency invocations
func (mmConvertSpendingsTableToUserCurrency *CurrencyConverterMock) ConvertSpendingsTableToUserCurrencyBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmConvertSpendingsTableToUserCurrency.beforeConvertSpendingsTableToUserCurrencyCounter)
}

// Calls returns a list of arguments used in each call to CurrencyConverterMock.ConvertSpendingsTableToUserCurrency.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmConvertSpendingsTableToUserCurrency *mCurrencyConverterMockConvertSpendingsTableToUserCurrency) Calls() []*CurrencyConverterMockConvertSpendingsTableToUserCurrencyParams {
	mmConvertSpendingsTableToUserCurrency.mutex.RLock()

	argCopy := make([]*CurrencyConverterMockConvertSpendingsTableToUserCurrencyParams, len(mmConvertSpendingsTableToUserCurrency.callArgs))
	copy(argCopy, mmConvertSpendingsTableToUserCurrency.callArgs)

	mmConvertSpendingsTableToUserCurrency.mutex.RUnlock()

	return argCopy
}

// MinimockConvertSpendingsTableToUserCurrencyDone returns true if the count of the ConvertSpendingsTableToUserCurrency invocations corresponds
// the number of defined expectations
func (m *CurrencyConverterMock) MinimockConvertSpendingsTableToUserCurrencyDone() bool {
	for _, e := range m.ConvertSpendingsTableToUserCurrencyMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ConvertSpendingsTableToUserCurrencyMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterConvertSpendingsTableToUserCurrencyCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcConvertSpendingsTableToUserCurrency != nil && mm_atomic.LoadUint64(&m.afterConvertSpendingsTableToUserCurrencyCounter) < 1 {
		return false
	}
	return true
}

// MinimockConvertSpendingsTableToUserCurrencyInspect logs each unmet expectation
func (m *CurrencyConverterMock) MinimockConvertSpendingsTableToUserCurrencyInspect() {
	for _, e := range m.ConvertSpendingsTableToUserCurrencyMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to CurrencyConverterMock.ConvertSpendingsTableToUserCurrency with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ConvertSpendingsTableToUserCurrencyMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterConvertSpendingsTableToUserCurrencyCounter) < 1 {
		if m.ConvertSpendingsTableToUserCurrencyMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to CurrencyConverterMock.ConvertSpendingsTableToUserCurrency")
		} else {
			m.t.Errorf("Expected call to CurrencyConverterMock.ConvertSpendingsTableToUserCurrency with params: %#v", *m.ConvertSpendingsTableToUserCurrencyMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcConvertSpendingsTableToUserCurrency != nil && mm_atomic.LoadUint64(&m.afterConvertSpendingsTableToUserCurrencyCounter) < 1 {
		m.t.Error("Expected call to CurrencyConverterMock.ConvertSpendingsTableToUserCurrency")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *CurrencyConverterMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockConvertSpendingsTableToUserCurrencyInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *CurrencyConverterMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *CurrencyConverterMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockConvertSpendingsTableToUserCurrencyDone()
}
