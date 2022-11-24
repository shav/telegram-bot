package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i github.com/shav/telegram-bot/internal/modules/finances/services/spendings.SpendingReportsCache -o ./mocks\spending_reports_cache.go -n SpendingReportsCacheMock

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"github.com/shav/telegram-bot/internal/common/date"
	finance_models "github.com/shav/telegram-bot/internal/modules/finances/domain/models"
)

// SpendingReportsCacheMock implements finance_services_spendings.SpendingReportsCache
type SpendingReportsCacheMock struct {
	t minimock.Tester

	funcAdd          func(ctx context.Context, report finance_models.SpendingsByCategoryTable, userId int64, dateInterval date.Interval) (err error)
	inspectFuncAdd   func(ctx context.Context, report finance_models.SpendingsByCategoryTable, userId int64, dateInterval date.Interval)
	afterAddCounter  uint64
	beforeAddCounter uint64
	AddMock          mSpendingReportsCacheMockAdd

	funcGet          func(ctx context.Context, userId int64, dateInterval date.Interval) (report finance_models.SpendingsByCategoryTable, exists bool, err error)
	inspectFuncGet   func(ctx context.Context, userId int64, dateInterval date.Interval)
	afterGetCounter  uint64
	beforeGetCounter uint64
	GetMock          mSpendingReportsCacheMockGet

	funcInvalidateForDate          func(ctx context.Context, userId int64, invalidDate date.Date) (err error)
	inspectFuncInvalidateForDate   func(ctx context.Context, userId int64, invalidDate date.Date)
	afterInvalidateForDateCounter  uint64
	beforeInvalidateForDateCounter uint64
	InvalidateForDateMock          mSpendingReportsCacheMockInvalidateForDate
}

// NewSpendingReportsCacheMock returns a mock for finance_services_spendings.SpendingReportsCache
func NewSpendingReportsCacheMock(t minimock.Tester) *SpendingReportsCacheMock {
	m := &SpendingReportsCacheMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.AddMock = mSpendingReportsCacheMockAdd{mock: m}
	m.AddMock.callArgs = []*SpendingReportsCacheMockAddParams{}

	m.GetMock = mSpendingReportsCacheMockGet{mock: m}
	m.GetMock.callArgs = []*SpendingReportsCacheMockGetParams{}

	m.InvalidateForDateMock = mSpendingReportsCacheMockInvalidateForDate{mock: m}
	m.InvalidateForDateMock.callArgs = []*SpendingReportsCacheMockInvalidateForDateParams{}

	return m
}

type mSpendingReportsCacheMockAdd struct {
	mock               *SpendingReportsCacheMock
	defaultExpectation *SpendingReportsCacheMockAddExpectation
	expectations       []*SpendingReportsCacheMockAddExpectation

	callArgs []*SpendingReportsCacheMockAddParams
	mutex    sync.RWMutex
}

// SpendingReportsCacheMockAddExpectation specifies expectation struct of the SpendingReportsCache.Add
type SpendingReportsCacheMockAddExpectation struct {
	mock    *SpendingReportsCacheMock
	params  *SpendingReportsCacheMockAddParams
	results *SpendingReportsCacheMockAddResults
	Counter uint64
}

// SpendingReportsCacheMockAddParams contains parameters of the SpendingReportsCache.Add
type SpendingReportsCacheMockAddParams struct {
	ctx          context.Context
	report       finance_models.SpendingsByCategoryTable
	userId       int64
	dateInterval date.Interval
}

// SpendingReportsCacheMockAddResults contains results of the SpendingReportsCache.Add
type SpendingReportsCacheMockAddResults struct {
	err error
}

// Expect sets up expected params for SpendingReportsCache.Add
func (mmAdd *mSpendingReportsCacheMockAdd) Expect(ctx context.Context, report finance_models.SpendingsByCategoryTable, userId int64, dateInterval date.Interval) *mSpendingReportsCacheMockAdd {
	if mmAdd.mock.funcAdd != nil {
		mmAdd.mock.t.Fatalf("SpendingReportsCacheMock.Add mock is already set by Set")
	}

	if mmAdd.defaultExpectation == nil {
		mmAdd.defaultExpectation = &SpendingReportsCacheMockAddExpectation{}
	}

	mmAdd.defaultExpectation.params = &SpendingReportsCacheMockAddParams{ctx, report, userId, dateInterval}
	for _, e := range mmAdd.expectations {
		if minimock.Equal(e.params, mmAdd.defaultExpectation.params) {
			mmAdd.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmAdd.defaultExpectation.params)
		}
	}

	return mmAdd
}

// Inspect accepts an inspector function that has same arguments as the SpendingReportsCache.Add
func (mmAdd *mSpendingReportsCacheMockAdd) Inspect(f func(ctx context.Context, report finance_models.SpendingsByCategoryTable, userId int64, dateInterval date.Interval)) *mSpendingReportsCacheMockAdd {
	if mmAdd.mock.inspectFuncAdd != nil {
		mmAdd.mock.t.Fatalf("Inspect function is already set for SpendingReportsCacheMock.Add")
	}

	mmAdd.mock.inspectFuncAdd = f

	return mmAdd
}

// Return sets up results that will be returned by SpendingReportsCache.Add
func (mmAdd *mSpendingReportsCacheMockAdd) Return(err error) *SpendingReportsCacheMock {
	if mmAdd.mock.funcAdd != nil {
		mmAdd.mock.t.Fatalf("SpendingReportsCacheMock.Add mock is already set by Set")
	}

	if mmAdd.defaultExpectation == nil {
		mmAdd.defaultExpectation = &SpendingReportsCacheMockAddExpectation{mock: mmAdd.mock}
	}
	mmAdd.defaultExpectation.results = &SpendingReportsCacheMockAddResults{err}
	return mmAdd.mock
}

//Set uses given function f to mock the SpendingReportsCache.Add method
func (mmAdd *mSpendingReportsCacheMockAdd) Set(f func(ctx context.Context, report finance_models.SpendingsByCategoryTable, userId int64, dateInterval date.Interval) (err error)) *SpendingReportsCacheMock {
	if mmAdd.defaultExpectation != nil {
		mmAdd.mock.t.Fatalf("Default expectation is already set for the SpendingReportsCache.Add method")
	}

	if len(mmAdd.expectations) > 0 {
		mmAdd.mock.t.Fatalf("Some expectations are already set for the SpendingReportsCache.Add method")
	}

	mmAdd.mock.funcAdd = f
	return mmAdd.mock
}

// When sets expectation for the SpendingReportsCache.Add which will trigger the result defined by the following
// Then helper
func (mmAdd *mSpendingReportsCacheMockAdd) When(ctx context.Context, report finance_models.SpendingsByCategoryTable, userId int64, dateInterval date.Interval) *SpendingReportsCacheMockAddExpectation {
	if mmAdd.mock.funcAdd != nil {
		mmAdd.mock.t.Fatalf("SpendingReportsCacheMock.Add mock is already set by Set")
	}

	expectation := &SpendingReportsCacheMockAddExpectation{
		mock:   mmAdd.mock,
		params: &SpendingReportsCacheMockAddParams{ctx, report, userId, dateInterval},
	}
	mmAdd.expectations = append(mmAdd.expectations, expectation)
	return expectation
}

// Then sets up SpendingReportsCache.Add return parameters for the expectation previously defined by the When method
func (e *SpendingReportsCacheMockAddExpectation) Then(err error) *SpendingReportsCacheMock {
	e.results = &SpendingReportsCacheMockAddResults{err}
	return e.mock
}

// Add implements finance_services_spendings.SpendingReportsCache
func (mmAdd *SpendingReportsCacheMock) Add(ctx context.Context, report finance_models.SpendingsByCategoryTable, userId int64, dateInterval date.Interval) (err error) {
	mm_atomic.AddUint64(&mmAdd.beforeAddCounter, 1)
	defer mm_atomic.AddUint64(&mmAdd.afterAddCounter, 1)

	if mmAdd.inspectFuncAdd != nil {
		mmAdd.inspectFuncAdd(ctx, report, userId, dateInterval)
	}

	mm_params := &SpendingReportsCacheMockAddParams{ctx, report, userId, dateInterval}

	// Record call args
	mmAdd.AddMock.mutex.Lock()
	mmAdd.AddMock.callArgs = append(mmAdd.AddMock.callArgs, mm_params)
	mmAdd.AddMock.mutex.Unlock()

	for _, e := range mmAdd.AddMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmAdd.AddMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmAdd.AddMock.defaultExpectation.Counter, 1)
		mm_want := mmAdd.AddMock.defaultExpectation.params
		mm_got := SpendingReportsCacheMockAddParams{ctx, report, userId, dateInterval}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmAdd.t.Errorf("SpendingReportsCacheMock.Add got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmAdd.AddMock.defaultExpectation.results
		if mm_results == nil {
			mmAdd.t.Fatal("No results are set for the SpendingReportsCacheMock.Add")
		}
		return (*mm_results).err
	}
	if mmAdd.funcAdd != nil {
		return mmAdd.funcAdd(ctx, report, userId, dateInterval)
	}
	mmAdd.t.Fatalf("Unexpected call to SpendingReportsCacheMock.Add. %v %v %v %v", ctx, report, userId, dateInterval)
	return
}

// AddAfterCounter returns a count of finished SpendingReportsCacheMock.Add invocations
func (mmAdd *SpendingReportsCacheMock) AddAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAdd.afterAddCounter)
}

// AddBeforeCounter returns a count of SpendingReportsCacheMock.Add invocations
func (mmAdd *SpendingReportsCacheMock) AddBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAdd.beforeAddCounter)
}

// Calls returns a list of arguments used in each call to SpendingReportsCacheMock.Add.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmAdd *mSpendingReportsCacheMockAdd) Calls() []*SpendingReportsCacheMockAddParams {
	mmAdd.mutex.RLock()

	argCopy := make([]*SpendingReportsCacheMockAddParams, len(mmAdd.callArgs))
	copy(argCopy, mmAdd.callArgs)

	mmAdd.mutex.RUnlock()

	return argCopy
}

// MinimockAddDone returns true if the count of the Add invocations corresponds
// the number of defined expectations
func (m *SpendingReportsCacheMock) MinimockAddDone() bool {
	for _, e := range m.AddMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.AddMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterAddCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAdd != nil && mm_atomic.LoadUint64(&m.afterAddCounter) < 1 {
		return false
	}
	return true
}

// MinimockAddInspect logs each unmet expectation
func (m *SpendingReportsCacheMock) MinimockAddInspect() {
	for _, e := range m.AddMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to SpendingReportsCacheMock.Add with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.AddMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterAddCounter) < 1 {
		if m.AddMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to SpendingReportsCacheMock.Add")
		} else {
			m.t.Errorf("Expected call to SpendingReportsCacheMock.Add with params: %#v", *m.AddMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAdd != nil && mm_atomic.LoadUint64(&m.afterAddCounter) < 1 {
		m.t.Error("Expected call to SpendingReportsCacheMock.Add")
	}
}

type mSpendingReportsCacheMockGet struct {
	mock               *SpendingReportsCacheMock
	defaultExpectation *SpendingReportsCacheMockGetExpectation
	expectations       []*SpendingReportsCacheMockGetExpectation

	callArgs []*SpendingReportsCacheMockGetParams
	mutex    sync.RWMutex
}

// SpendingReportsCacheMockGetExpectation specifies expectation struct of the SpendingReportsCache.Get
type SpendingReportsCacheMockGetExpectation struct {
	mock    *SpendingReportsCacheMock
	params  *SpendingReportsCacheMockGetParams
	results *SpendingReportsCacheMockGetResults
	Counter uint64
}

// SpendingReportsCacheMockGetParams contains parameters of the SpendingReportsCache.Get
type SpendingReportsCacheMockGetParams struct {
	ctx          context.Context
	userId       int64
	dateInterval date.Interval
}

// SpendingReportsCacheMockGetResults contains results of the SpendingReportsCache.Get
type SpendingReportsCacheMockGetResults struct {
	report finance_models.SpendingsByCategoryTable
	exists bool
	err    error
}

// Expect sets up expected params for SpendingReportsCache.Get
func (mmGet *mSpendingReportsCacheMockGet) Expect(ctx context.Context, userId int64, dateInterval date.Interval) *mSpendingReportsCacheMockGet {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("SpendingReportsCacheMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &SpendingReportsCacheMockGetExpectation{}
	}

	mmGet.defaultExpectation.params = &SpendingReportsCacheMockGetParams{ctx, userId, dateInterval}
	for _, e := range mmGet.expectations {
		if minimock.Equal(e.params, mmGet.defaultExpectation.params) {
			mmGet.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGet.defaultExpectation.params)
		}
	}

	return mmGet
}

// Inspect accepts an inspector function that has same arguments as the SpendingReportsCache.Get
func (mmGet *mSpendingReportsCacheMockGet) Inspect(f func(ctx context.Context, userId int64, dateInterval date.Interval)) *mSpendingReportsCacheMockGet {
	if mmGet.mock.inspectFuncGet != nil {
		mmGet.mock.t.Fatalf("Inspect function is already set for SpendingReportsCacheMock.Get")
	}

	mmGet.mock.inspectFuncGet = f

	return mmGet
}

// Return sets up results that will be returned by SpendingReportsCache.Get
func (mmGet *mSpendingReportsCacheMockGet) Return(report finance_models.SpendingsByCategoryTable, exists bool, err error) *SpendingReportsCacheMock {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("SpendingReportsCacheMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &SpendingReportsCacheMockGetExpectation{mock: mmGet.mock}
	}
	mmGet.defaultExpectation.results = &SpendingReportsCacheMockGetResults{report, exists, err}
	return mmGet.mock
}

//Set uses given function f to mock the SpendingReportsCache.Get method
func (mmGet *mSpendingReportsCacheMockGet) Set(f func(ctx context.Context, userId int64, dateInterval date.Interval) (report finance_models.SpendingsByCategoryTable, exists bool, err error)) *SpendingReportsCacheMock {
	if mmGet.defaultExpectation != nil {
		mmGet.mock.t.Fatalf("Default expectation is already set for the SpendingReportsCache.Get method")
	}

	if len(mmGet.expectations) > 0 {
		mmGet.mock.t.Fatalf("Some expectations are already set for the SpendingReportsCache.Get method")
	}

	mmGet.mock.funcGet = f
	return mmGet.mock
}

// When sets expectation for the SpendingReportsCache.Get which will trigger the result defined by the following
// Then helper
func (mmGet *mSpendingReportsCacheMockGet) When(ctx context.Context, userId int64, dateInterval date.Interval) *SpendingReportsCacheMockGetExpectation {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("SpendingReportsCacheMock.Get mock is already set by Set")
	}

	expectation := &SpendingReportsCacheMockGetExpectation{
		mock:   mmGet.mock,
		params: &SpendingReportsCacheMockGetParams{ctx, userId, dateInterval},
	}
	mmGet.expectations = append(mmGet.expectations, expectation)
	return expectation
}

// Then sets up SpendingReportsCache.Get return parameters for the expectation previously defined by the When method
func (e *SpendingReportsCacheMockGetExpectation) Then(report finance_models.SpendingsByCategoryTable, exists bool, err error) *SpendingReportsCacheMock {
	e.results = &SpendingReportsCacheMockGetResults{report, exists, err}
	return e.mock
}

// Get implements finance_services_spendings.SpendingReportsCache
func (mmGet *SpendingReportsCacheMock) Get(ctx context.Context, userId int64, dateInterval date.Interval) (report finance_models.SpendingsByCategoryTable, exists bool, err error) {
	mm_atomic.AddUint64(&mmGet.beforeGetCounter, 1)
	defer mm_atomic.AddUint64(&mmGet.afterGetCounter, 1)

	if mmGet.inspectFuncGet != nil {
		mmGet.inspectFuncGet(ctx, userId, dateInterval)
	}

	mm_params := &SpendingReportsCacheMockGetParams{ctx, userId, dateInterval}

	// Record call args
	mmGet.GetMock.mutex.Lock()
	mmGet.GetMock.callArgs = append(mmGet.GetMock.callArgs, mm_params)
	mmGet.GetMock.mutex.Unlock()

	for _, e := range mmGet.GetMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.report, e.results.exists, e.results.err
		}
	}

	if mmGet.GetMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGet.GetMock.defaultExpectation.Counter, 1)
		mm_want := mmGet.GetMock.defaultExpectation.params
		mm_got := SpendingReportsCacheMockGetParams{ctx, userId, dateInterval}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGet.t.Errorf("SpendingReportsCacheMock.Get got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGet.GetMock.defaultExpectation.results
		if mm_results == nil {
			mmGet.t.Fatal("No results are set for the SpendingReportsCacheMock.Get")
		}
		return (*mm_results).report, (*mm_results).exists, (*mm_results).err
	}
	if mmGet.funcGet != nil {
		return mmGet.funcGet(ctx, userId, dateInterval)
	}
	mmGet.t.Fatalf("Unexpected call to SpendingReportsCacheMock.Get. %v %v %v", ctx, userId, dateInterval)
	return
}

// GetAfterCounter returns a count of finished SpendingReportsCacheMock.Get invocations
func (mmGet *SpendingReportsCacheMock) GetAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.afterGetCounter)
}

// GetBeforeCounter returns a count of SpendingReportsCacheMock.Get invocations
func (mmGet *SpendingReportsCacheMock) GetBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.beforeGetCounter)
}

// Calls returns a list of arguments used in each call to SpendingReportsCacheMock.Get.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGet *mSpendingReportsCacheMockGet) Calls() []*SpendingReportsCacheMockGetParams {
	mmGet.mutex.RLock()

	argCopy := make([]*SpendingReportsCacheMockGetParams, len(mmGet.callArgs))
	copy(argCopy, mmGet.callArgs)

	mmGet.mutex.RUnlock()

	return argCopy
}

// MinimockGetDone returns true if the count of the Get invocations corresponds
// the number of defined expectations
func (m *SpendingReportsCacheMock) MinimockGetDone() bool {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetInspect logs each unmet expectation
func (m *SpendingReportsCacheMock) MinimockGetInspect() {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to SpendingReportsCacheMock.Get with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		if m.GetMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to SpendingReportsCacheMock.Get")
		} else {
			m.t.Errorf("Expected call to SpendingReportsCacheMock.Get with params: %#v", *m.GetMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		m.t.Error("Expected call to SpendingReportsCacheMock.Get")
	}
}

type mSpendingReportsCacheMockInvalidateForDate struct {
	mock               *SpendingReportsCacheMock
	defaultExpectation *SpendingReportsCacheMockInvalidateForDateExpectation
	expectations       []*SpendingReportsCacheMockInvalidateForDateExpectation

	callArgs []*SpendingReportsCacheMockInvalidateForDateParams
	mutex    sync.RWMutex
}

// SpendingReportsCacheMockInvalidateForDateExpectation specifies expectation struct of the SpendingReportsCache.InvalidateForDate
type SpendingReportsCacheMockInvalidateForDateExpectation struct {
	mock    *SpendingReportsCacheMock
	params  *SpendingReportsCacheMockInvalidateForDateParams
	results *SpendingReportsCacheMockInvalidateForDateResults
	Counter uint64
}

// SpendingReportsCacheMockInvalidateForDateParams contains parameters of the SpendingReportsCache.InvalidateForDate
type SpendingReportsCacheMockInvalidateForDateParams struct {
	ctx         context.Context
	userId      int64
	invalidDate date.Date
}

// SpendingReportsCacheMockInvalidateForDateResults contains results of the SpendingReportsCache.InvalidateForDate
type SpendingReportsCacheMockInvalidateForDateResults struct {
	err error
}

// Expect sets up expected params for SpendingReportsCache.InvalidateForDate
func (mmInvalidateForDate *mSpendingReportsCacheMockInvalidateForDate) Expect(ctx context.Context, userId int64, invalidDate date.Date) *mSpendingReportsCacheMockInvalidateForDate {
	if mmInvalidateForDate.mock.funcInvalidateForDate != nil {
		mmInvalidateForDate.mock.t.Fatalf("SpendingReportsCacheMock.InvalidateForDate mock is already set by Set")
	}

	if mmInvalidateForDate.defaultExpectation == nil {
		mmInvalidateForDate.defaultExpectation = &SpendingReportsCacheMockInvalidateForDateExpectation{}
	}

	mmInvalidateForDate.defaultExpectation.params = &SpendingReportsCacheMockInvalidateForDateParams{ctx, userId, invalidDate}
	for _, e := range mmInvalidateForDate.expectations {
		if minimock.Equal(e.params, mmInvalidateForDate.defaultExpectation.params) {
			mmInvalidateForDate.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmInvalidateForDate.defaultExpectation.params)
		}
	}

	return mmInvalidateForDate
}

// Inspect accepts an inspector function that has same arguments as the SpendingReportsCache.InvalidateForDate
func (mmInvalidateForDate *mSpendingReportsCacheMockInvalidateForDate) Inspect(f func(ctx context.Context, userId int64, invalidDate date.Date)) *mSpendingReportsCacheMockInvalidateForDate {
	if mmInvalidateForDate.mock.inspectFuncInvalidateForDate != nil {
		mmInvalidateForDate.mock.t.Fatalf("Inspect function is already set for SpendingReportsCacheMock.InvalidateForDate")
	}

	mmInvalidateForDate.mock.inspectFuncInvalidateForDate = f

	return mmInvalidateForDate
}

// Return sets up results that will be returned by SpendingReportsCache.InvalidateForDate
func (mmInvalidateForDate *mSpendingReportsCacheMockInvalidateForDate) Return(err error) *SpendingReportsCacheMock {
	if mmInvalidateForDate.mock.funcInvalidateForDate != nil {
		mmInvalidateForDate.mock.t.Fatalf("SpendingReportsCacheMock.InvalidateForDate mock is already set by Set")
	}

	if mmInvalidateForDate.defaultExpectation == nil {
		mmInvalidateForDate.defaultExpectation = &SpendingReportsCacheMockInvalidateForDateExpectation{mock: mmInvalidateForDate.mock}
	}
	mmInvalidateForDate.defaultExpectation.results = &SpendingReportsCacheMockInvalidateForDateResults{err}
	return mmInvalidateForDate.mock
}

//Set uses given function f to mock the SpendingReportsCache.InvalidateForDate method
func (mmInvalidateForDate *mSpendingReportsCacheMockInvalidateForDate) Set(f func(ctx context.Context, userId int64, invalidDate date.Date) (err error)) *SpendingReportsCacheMock {
	if mmInvalidateForDate.defaultExpectation != nil {
		mmInvalidateForDate.mock.t.Fatalf("Default expectation is already set for the SpendingReportsCache.InvalidateForDate method")
	}

	if len(mmInvalidateForDate.expectations) > 0 {
		mmInvalidateForDate.mock.t.Fatalf("Some expectations are already set for the SpendingReportsCache.InvalidateForDate method")
	}

	mmInvalidateForDate.mock.funcInvalidateForDate = f
	return mmInvalidateForDate.mock
}

// When sets expectation for the SpendingReportsCache.InvalidateForDate which will trigger the result defined by the following
// Then helper
func (mmInvalidateForDate *mSpendingReportsCacheMockInvalidateForDate) When(ctx context.Context, userId int64, invalidDate date.Date) *SpendingReportsCacheMockInvalidateForDateExpectation {
	if mmInvalidateForDate.mock.funcInvalidateForDate != nil {
		mmInvalidateForDate.mock.t.Fatalf("SpendingReportsCacheMock.InvalidateForDate mock is already set by Set")
	}

	expectation := &SpendingReportsCacheMockInvalidateForDateExpectation{
		mock:   mmInvalidateForDate.mock,
		params: &SpendingReportsCacheMockInvalidateForDateParams{ctx, userId, invalidDate},
	}
	mmInvalidateForDate.expectations = append(mmInvalidateForDate.expectations, expectation)
	return expectation
}

// Then sets up SpendingReportsCache.InvalidateForDate return parameters for the expectation previously defined by the When method
func (e *SpendingReportsCacheMockInvalidateForDateExpectation) Then(err error) *SpendingReportsCacheMock {
	e.results = &SpendingReportsCacheMockInvalidateForDateResults{err}
	return e.mock
}

// InvalidateForDate implements finance_services_spendings.SpendingReportsCache
func (mmInvalidateForDate *SpendingReportsCacheMock) InvalidateForDate(ctx context.Context, userId int64, invalidDate date.Date) (err error) {
	mm_atomic.AddUint64(&mmInvalidateForDate.beforeInvalidateForDateCounter, 1)
	defer mm_atomic.AddUint64(&mmInvalidateForDate.afterInvalidateForDateCounter, 1)

	if mmInvalidateForDate.inspectFuncInvalidateForDate != nil {
		mmInvalidateForDate.inspectFuncInvalidateForDate(ctx, userId, invalidDate)
	}

	mm_params := &SpendingReportsCacheMockInvalidateForDateParams{ctx, userId, invalidDate}

	// Record call args
	mmInvalidateForDate.InvalidateForDateMock.mutex.Lock()
	mmInvalidateForDate.InvalidateForDateMock.callArgs = append(mmInvalidateForDate.InvalidateForDateMock.callArgs, mm_params)
	mmInvalidateForDate.InvalidateForDateMock.mutex.Unlock()

	for _, e := range mmInvalidateForDate.InvalidateForDateMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmInvalidateForDate.InvalidateForDateMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmInvalidateForDate.InvalidateForDateMock.defaultExpectation.Counter, 1)
		mm_want := mmInvalidateForDate.InvalidateForDateMock.defaultExpectation.params
		mm_got := SpendingReportsCacheMockInvalidateForDateParams{ctx, userId, invalidDate}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmInvalidateForDate.t.Errorf("SpendingReportsCacheMock.InvalidateForDate got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmInvalidateForDate.InvalidateForDateMock.defaultExpectation.results
		if mm_results == nil {
			mmInvalidateForDate.t.Fatal("No results are set for the SpendingReportsCacheMock.InvalidateForDate")
		}
		return (*mm_results).err
	}
	if mmInvalidateForDate.funcInvalidateForDate != nil {
		return mmInvalidateForDate.funcInvalidateForDate(ctx, userId, invalidDate)
	}
	mmInvalidateForDate.t.Fatalf("Unexpected call to SpendingReportsCacheMock.InvalidateForDate. %v %v %v", ctx, userId, invalidDate)
	return
}

// InvalidateForDateAfterCounter returns a count of finished SpendingReportsCacheMock.InvalidateForDate invocations
func (mmInvalidateForDate *SpendingReportsCacheMock) InvalidateForDateAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmInvalidateForDate.afterInvalidateForDateCounter)
}

// InvalidateForDateBeforeCounter returns a count of SpendingReportsCacheMock.InvalidateForDate invocations
func (mmInvalidateForDate *SpendingReportsCacheMock) InvalidateForDateBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmInvalidateForDate.beforeInvalidateForDateCounter)
}

// Calls returns a list of arguments used in each call to SpendingReportsCacheMock.InvalidateForDate.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmInvalidateForDate *mSpendingReportsCacheMockInvalidateForDate) Calls() []*SpendingReportsCacheMockInvalidateForDateParams {
	mmInvalidateForDate.mutex.RLock()

	argCopy := make([]*SpendingReportsCacheMockInvalidateForDateParams, len(mmInvalidateForDate.callArgs))
	copy(argCopy, mmInvalidateForDate.callArgs)

	mmInvalidateForDate.mutex.RUnlock()

	return argCopy
}

// MinimockInvalidateForDateDone returns true if the count of the InvalidateForDate invocations corresponds
// the number of defined expectations
func (m *SpendingReportsCacheMock) MinimockInvalidateForDateDone() bool {
	for _, e := range m.InvalidateForDateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.InvalidateForDateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterInvalidateForDateCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcInvalidateForDate != nil && mm_atomic.LoadUint64(&m.afterInvalidateForDateCounter) < 1 {
		return false
	}
	return true
}

// MinimockInvalidateForDateInspect logs each unmet expectation
func (m *SpendingReportsCacheMock) MinimockInvalidateForDateInspect() {
	for _, e := range m.InvalidateForDateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to SpendingReportsCacheMock.InvalidateForDate with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.InvalidateForDateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterInvalidateForDateCounter) < 1 {
		if m.InvalidateForDateMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to SpendingReportsCacheMock.InvalidateForDate")
		} else {
			m.t.Errorf("Expected call to SpendingReportsCacheMock.InvalidateForDate with params: %#v", *m.InvalidateForDateMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcInvalidateForDate != nil && mm_atomic.LoadUint64(&m.afterInvalidateForDateCounter) < 1 {
		m.t.Error("Expected call to SpendingReportsCacheMock.InvalidateForDate")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *SpendingReportsCacheMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockAddInspect()

		m.MinimockGetInspect()

		m.MinimockInvalidateForDateInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *SpendingReportsCacheMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *SpendingReportsCacheMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockAddDone() &&
		m.MinimockGetDone() &&
		m.MinimockInvalidateForDateDone()
}