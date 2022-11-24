package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i github.com/shav/telegram-bot/internal/modules/finances/domain/use_cases.reportsBuilder -o ./mocks\reports_builder.go -n ReportsBuilderMock

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"github.com/shav/telegram-bot/internal/common/date"
)

// ReportsBuilderMock implements finances.reportsBuilder
type ReportsBuilderMock struct {
	t minimock.Tester

	funcRequestSpendingReport          func(ctx context.Context, userId int64, periodName string, dateInterval date.Interval) (err error)
	inspectFuncRequestSpendingReport   func(ctx context.Context, userId int64, periodName string, dateInterval date.Interval)
	afterRequestSpendingReportCounter  uint64
	beforeRequestSpendingReportCounter uint64
	RequestSpendingReportMock          mReportsBuilderMockRequestSpendingReport
}

// NewReportsBuilderMock returns a mock for finances.reportsBuilder
func NewReportsBuilderMock(t minimock.Tester) *ReportsBuilderMock {
	m := &ReportsBuilderMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.RequestSpendingReportMock = mReportsBuilderMockRequestSpendingReport{mock: m}
	m.RequestSpendingReportMock.callArgs = []*ReportsBuilderMockRequestSpendingReportParams{}

	return m
}

type mReportsBuilderMockRequestSpendingReport struct {
	mock               *ReportsBuilderMock
	defaultExpectation *ReportsBuilderMockRequestSpendingReportExpectation
	expectations       []*ReportsBuilderMockRequestSpendingReportExpectation

	callArgs []*ReportsBuilderMockRequestSpendingReportParams
	mutex    sync.RWMutex
}

// ReportsBuilderMockRequestSpendingReportExpectation specifies expectation struct of the reportsBuilder.RequestSpendingReport
type ReportsBuilderMockRequestSpendingReportExpectation struct {
	mock    *ReportsBuilderMock
	params  *ReportsBuilderMockRequestSpendingReportParams
	results *ReportsBuilderMockRequestSpendingReportResults
	Counter uint64
}

// ReportsBuilderMockRequestSpendingReportParams contains parameters of the reportsBuilder.RequestSpendingReport
type ReportsBuilderMockRequestSpendingReportParams struct {
	ctx          context.Context
	userId       int64
	periodName   string
	dateInterval date.Interval
}

// ReportsBuilderMockRequestSpendingReportResults contains results of the reportsBuilder.RequestSpendingReport
type ReportsBuilderMockRequestSpendingReportResults struct {
	err error
}

// Expect sets up expected params for reportsBuilder.RequestSpendingReport
func (mmRequestSpendingReport *mReportsBuilderMockRequestSpendingReport) Expect(ctx context.Context, userId int64, periodName string, dateInterval date.Interval) *mReportsBuilderMockRequestSpendingReport {
	if mmRequestSpendingReport.mock.funcRequestSpendingReport != nil {
		mmRequestSpendingReport.mock.t.Fatalf("ReportsBuilderMock.RequestSpendingReport mock is already set by Set")
	}

	if mmRequestSpendingReport.defaultExpectation == nil {
		mmRequestSpendingReport.defaultExpectation = &ReportsBuilderMockRequestSpendingReportExpectation{}
	}

	mmRequestSpendingReport.defaultExpectation.params = &ReportsBuilderMockRequestSpendingReportParams{ctx, userId, periodName, dateInterval}
	for _, e := range mmRequestSpendingReport.expectations {
		if minimock.Equal(e.params, mmRequestSpendingReport.defaultExpectation.params) {
			mmRequestSpendingReport.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmRequestSpendingReport.defaultExpectation.params)
		}
	}

	return mmRequestSpendingReport
}

// Inspect accepts an inspector function that has same arguments as the reportsBuilder.RequestSpendingReport
func (mmRequestSpendingReport *mReportsBuilderMockRequestSpendingReport) Inspect(f func(ctx context.Context, userId int64, periodName string, dateInterval date.Interval)) *mReportsBuilderMockRequestSpendingReport {
	if mmRequestSpendingReport.mock.inspectFuncRequestSpendingReport != nil {
		mmRequestSpendingReport.mock.t.Fatalf("Inspect function is already set for ReportsBuilderMock.RequestSpendingReport")
	}

	mmRequestSpendingReport.mock.inspectFuncRequestSpendingReport = f

	return mmRequestSpendingReport
}

// Return sets up results that will be returned by reportsBuilder.RequestSpendingReport
func (mmRequestSpendingReport *mReportsBuilderMockRequestSpendingReport) Return(err error) *ReportsBuilderMock {
	if mmRequestSpendingReport.mock.funcRequestSpendingReport != nil {
		mmRequestSpendingReport.mock.t.Fatalf("ReportsBuilderMock.RequestSpendingReport mock is already set by Set")
	}

	if mmRequestSpendingReport.defaultExpectation == nil {
		mmRequestSpendingReport.defaultExpectation = &ReportsBuilderMockRequestSpendingReportExpectation{mock: mmRequestSpendingReport.mock}
	}
	mmRequestSpendingReport.defaultExpectation.results = &ReportsBuilderMockRequestSpendingReportResults{err}
	return mmRequestSpendingReport.mock
}

//Set uses given function f to mock the reportsBuilder.RequestSpendingReport method
func (mmRequestSpendingReport *mReportsBuilderMockRequestSpendingReport) Set(f func(ctx context.Context, userId int64, periodName string, dateInterval date.Interval) (err error)) *ReportsBuilderMock {
	if mmRequestSpendingReport.defaultExpectation != nil {
		mmRequestSpendingReport.mock.t.Fatalf("Default expectation is already set for the reportsBuilder.RequestSpendingReport method")
	}

	if len(mmRequestSpendingReport.expectations) > 0 {
		mmRequestSpendingReport.mock.t.Fatalf("Some expectations are already set for the reportsBuilder.RequestSpendingReport method")
	}

	mmRequestSpendingReport.mock.funcRequestSpendingReport = f
	return mmRequestSpendingReport.mock
}

// When sets expectation for the reportsBuilder.RequestSpendingReport which will trigger the result defined by the following
// Then helper
func (mmRequestSpendingReport *mReportsBuilderMockRequestSpendingReport) When(ctx context.Context, userId int64, periodName string, dateInterval date.Interval) *ReportsBuilderMockRequestSpendingReportExpectation {
	if mmRequestSpendingReport.mock.funcRequestSpendingReport != nil {
		mmRequestSpendingReport.mock.t.Fatalf("ReportsBuilderMock.RequestSpendingReport mock is already set by Set")
	}

	expectation := &ReportsBuilderMockRequestSpendingReportExpectation{
		mock:   mmRequestSpendingReport.mock,
		params: &ReportsBuilderMockRequestSpendingReportParams{ctx, userId, periodName, dateInterval},
	}
	mmRequestSpendingReport.expectations = append(mmRequestSpendingReport.expectations, expectation)
	return expectation
}

// Then sets up reportsBuilder.RequestSpendingReport return parameters for the expectation previously defined by the When method
func (e *ReportsBuilderMockRequestSpendingReportExpectation) Then(err error) *ReportsBuilderMock {
	e.results = &ReportsBuilderMockRequestSpendingReportResults{err}
	return e.mock
}

// RequestSpendingReport implements finances.reportsBuilder
func (mmRequestSpendingReport *ReportsBuilderMock) RequestSpendingReport(ctx context.Context, userId int64, periodName string, dateInterval date.Interval) (err error) {
	mm_atomic.AddUint64(&mmRequestSpendingReport.beforeRequestSpendingReportCounter, 1)
	defer mm_atomic.AddUint64(&mmRequestSpendingReport.afterRequestSpendingReportCounter, 1)

	if mmRequestSpendingReport.inspectFuncRequestSpendingReport != nil {
		mmRequestSpendingReport.inspectFuncRequestSpendingReport(ctx, userId, periodName, dateInterval)
	}

	mm_params := &ReportsBuilderMockRequestSpendingReportParams{ctx, userId, periodName, dateInterval}

	// Record call args
	mmRequestSpendingReport.RequestSpendingReportMock.mutex.Lock()
	mmRequestSpendingReport.RequestSpendingReportMock.callArgs = append(mmRequestSpendingReport.RequestSpendingReportMock.callArgs, mm_params)
	mmRequestSpendingReport.RequestSpendingReportMock.mutex.Unlock()

	for _, e := range mmRequestSpendingReport.RequestSpendingReportMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmRequestSpendingReport.RequestSpendingReportMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmRequestSpendingReport.RequestSpendingReportMock.defaultExpectation.Counter, 1)
		mm_want := mmRequestSpendingReport.RequestSpendingReportMock.defaultExpectation.params
		mm_got := ReportsBuilderMockRequestSpendingReportParams{ctx, userId, periodName, dateInterval}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmRequestSpendingReport.t.Errorf("ReportsBuilderMock.RequestSpendingReport got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmRequestSpendingReport.RequestSpendingReportMock.defaultExpectation.results
		if mm_results == nil {
			mmRequestSpendingReport.t.Fatal("No results are set for the ReportsBuilderMock.RequestSpendingReport")
		}
		return (*mm_results).err
	}
	if mmRequestSpendingReport.funcRequestSpendingReport != nil {
		return mmRequestSpendingReport.funcRequestSpendingReport(ctx, userId, periodName, dateInterval)
	}
	mmRequestSpendingReport.t.Fatalf("Unexpected call to ReportsBuilderMock.RequestSpendingReport. %v %v %v %v", ctx, userId, periodName, dateInterval)
	return
}

// RequestSpendingReportAfterCounter returns a count of finished ReportsBuilderMock.RequestSpendingReport invocations
func (mmRequestSpendingReport *ReportsBuilderMock) RequestSpendingReportAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmRequestSpendingReport.afterRequestSpendingReportCounter)
}

// RequestSpendingReportBeforeCounter returns a count of ReportsBuilderMock.RequestSpendingReport invocations
func (mmRequestSpendingReport *ReportsBuilderMock) RequestSpendingReportBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmRequestSpendingReport.beforeRequestSpendingReportCounter)
}

// Calls returns a list of arguments used in each call to ReportsBuilderMock.RequestSpendingReport.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmRequestSpendingReport *mReportsBuilderMockRequestSpendingReport) Calls() []*ReportsBuilderMockRequestSpendingReportParams {
	mmRequestSpendingReport.mutex.RLock()

	argCopy := make([]*ReportsBuilderMockRequestSpendingReportParams, len(mmRequestSpendingReport.callArgs))
	copy(argCopy, mmRequestSpendingReport.callArgs)

	mmRequestSpendingReport.mutex.RUnlock()

	return argCopy
}

// MinimockRequestSpendingReportDone returns true if the count of the RequestSpendingReport invocations corresponds
// the number of defined expectations
func (m *ReportsBuilderMock) MinimockRequestSpendingReportDone() bool {
	for _, e := range m.RequestSpendingReportMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.RequestSpendingReportMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterRequestSpendingReportCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcRequestSpendingReport != nil && mm_atomic.LoadUint64(&m.afterRequestSpendingReportCounter) < 1 {
		return false
	}
	return true
}

// MinimockRequestSpendingReportInspect logs each unmet expectation
func (m *ReportsBuilderMock) MinimockRequestSpendingReportInspect() {
	for _, e := range m.RequestSpendingReportMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ReportsBuilderMock.RequestSpendingReport with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.RequestSpendingReportMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterRequestSpendingReportCounter) < 1 {
		if m.RequestSpendingReportMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ReportsBuilderMock.RequestSpendingReport")
		} else {
			m.t.Errorf("Expected call to ReportsBuilderMock.RequestSpendingReport with params: %#v", *m.RequestSpendingReportMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcRequestSpendingReport != nil && mm_atomic.LoadUint64(&m.afterRequestSpendingReportCounter) < 1 {
		m.t.Error("Expected call to ReportsBuilderMock.RequestSpendingReport")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *ReportsBuilderMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockRequestSpendingReportInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *ReportsBuilderMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *ReportsBuilderMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockRequestSpendingReportDone()
}