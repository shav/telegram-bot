package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i github.com/shav/telegram-bot/internal/modules/finances/domain/use_cases.transactionManager -o ./mocks\transaction_manager.go -n TransactionManagerMock

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	tr "github.com/shav/telegram-bot/internal/common/transactions"
)

// TransactionManagerMock implements fin_use_cases.transactionManager
type TransactionManagerMock struct {
	t minimock.Tester

	funcBeginTransaction          func(ctx context.Context) (t1 tr.Transaction, err error)
	inspectFuncBeginTransaction   func(ctx context.Context)
	afterBeginTransactionCounter  uint64
	beforeBeginTransactionCounter uint64
	BeginTransactionMock          mTransactionManagerMockBeginTransaction
}

// NewTransactionManagerMock returns a mock for fin_use_cases.transactionManager
func NewTransactionManagerMock(t minimock.Tester) *TransactionManagerMock {
	m := &TransactionManagerMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.BeginTransactionMock = mTransactionManagerMockBeginTransaction{mock: m}
	m.BeginTransactionMock.callArgs = []*TransactionManagerMockBeginTransactionParams{}

	return m
}

type mTransactionManagerMockBeginTransaction struct {
	mock               *TransactionManagerMock
	defaultExpectation *TransactionManagerMockBeginTransactionExpectation
	expectations       []*TransactionManagerMockBeginTransactionExpectation

	callArgs []*TransactionManagerMockBeginTransactionParams
	mutex    sync.RWMutex
}

// TransactionManagerMockBeginTransactionExpectation specifies expectation struct of the transactionManager.BeginTransaction
type TransactionManagerMockBeginTransactionExpectation struct {
	mock    *TransactionManagerMock
	params  *TransactionManagerMockBeginTransactionParams
	results *TransactionManagerMockBeginTransactionResults
	Counter uint64
}

// TransactionManagerMockBeginTransactionParams contains parameters of the transactionManager.BeginTransaction
type TransactionManagerMockBeginTransactionParams struct {
	ctx context.Context
}

// TransactionManagerMockBeginTransactionResults contains results of the transactionManager.BeginTransaction
type TransactionManagerMockBeginTransactionResults struct {
	t1  tr.Transaction
	err error
}

// Expect sets up expected params for transactionManager.BeginTransaction
func (mmBeginTransaction *mTransactionManagerMockBeginTransaction) Expect(ctx context.Context) *mTransactionManagerMockBeginTransaction {
	if mmBeginTransaction.mock.funcBeginTransaction != nil {
		mmBeginTransaction.mock.t.Fatalf("TransactionManagerMock.BeginTransaction mock is already set by Set")
	}

	if mmBeginTransaction.defaultExpectation == nil {
		mmBeginTransaction.defaultExpectation = &TransactionManagerMockBeginTransactionExpectation{}
	}

	mmBeginTransaction.defaultExpectation.params = &TransactionManagerMockBeginTransactionParams{ctx}
	for _, e := range mmBeginTransaction.expectations {
		if minimock.Equal(e.params, mmBeginTransaction.defaultExpectation.params) {
			mmBeginTransaction.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmBeginTransaction.defaultExpectation.params)
		}
	}

	return mmBeginTransaction
}

// Inspect accepts an inspector function that has same arguments as the transactionManager.BeginTransaction
func (mmBeginTransaction *mTransactionManagerMockBeginTransaction) Inspect(f func(ctx context.Context)) *mTransactionManagerMockBeginTransaction {
	if mmBeginTransaction.mock.inspectFuncBeginTransaction != nil {
		mmBeginTransaction.mock.t.Fatalf("Inspect function is already set for TransactionManagerMock.BeginTransaction")
	}

	mmBeginTransaction.mock.inspectFuncBeginTransaction = f

	return mmBeginTransaction
}

// Return sets up results that will be returned by transactionManager.BeginTransaction
func (mmBeginTransaction *mTransactionManagerMockBeginTransaction) Return(t1 tr.Transaction, err error) *TransactionManagerMock {
	if mmBeginTransaction.mock.funcBeginTransaction != nil {
		mmBeginTransaction.mock.t.Fatalf("TransactionManagerMock.BeginTransaction mock is already set by Set")
	}

	if mmBeginTransaction.defaultExpectation == nil {
		mmBeginTransaction.defaultExpectation = &TransactionManagerMockBeginTransactionExpectation{mock: mmBeginTransaction.mock}
	}
	mmBeginTransaction.defaultExpectation.results = &TransactionManagerMockBeginTransactionResults{t1, err}
	return mmBeginTransaction.mock
}

//Set uses given function f to mock the transactionManager.BeginTransaction method
func (mmBeginTransaction *mTransactionManagerMockBeginTransaction) Set(f func(ctx context.Context) (t1 tr.Transaction, err error)) *TransactionManagerMock {
	if mmBeginTransaction.defaultExpectation != nil {
		mmBeginTransaction.mock.t.Fatalf("Default expectation is already set for the transactionManager.BeginTransaction method")
	}

	if len(mmBeginTransaction.expectations) > 0 {
		mmBeginTransaction.mock.t.Fatalf("Some expectations are already set for the transactionManager.BeginTransaction method")
	}

	mmBeginTransaction.mock.funcBeginTransaction = f
	return mmBeginTransaction.mock
}

// When sets expectation for the transactionManager.BeginTransaction which will trigger the result defined by the following
// Then helper
func (mmBeginTransaction *mTransactionManagerMockBeginTransaction) When(ctx context.Context) *TransactionManagerMockBeginTransactionExpectation {
	if mmBeginTransaction.mock.funcBeginTransaction != nil {
		mmBeginTransaction.mock.t.Fatalf("TransactionManagerMock.BeginTransaction mock is already set by Set")
	}

	expectation := &TransactionManagerMockBeginTransactionExpectation{
		mock:   mmBeginTransaction.mock,
		params: &TransactionManagerMockBeginTransactionParams{ctx},
	}
	mmBeginTransaction.expectations = append(mmBeginTransaction.expectations, expectation)
	return expectation
}

// Then sets up transactionManager.BeginTransaction return parameters for the expectation previously defined by the When method
func (e *TransactionManagerMockBeginTransactionExpectation) Then(t1 tr.Transaction, err error) *TransactionManagerMock {
	e.results = &TransactionManagerMockBeginTransactionResults{t1, err}
	return e.mock
}

// BeginTransaction implements fin_use_cases.transactionManager
func (mmBeginTransaction *TransactionManagerMock) BeginTransaction(ctx context.Context) (t1 tr.Transaction, err error) {
	mm_atomic.AddUint64(&mmBeginTransaction.beforeBeginTransactionCounter, 1)
	defer mm_atomic.AddUint64(&mmBeginTransaction.afterBeginTransactionCounter, 1)

	if mmBeginTransaction.inspectFuncBeginTransaction != nil {
		mmBeginTransaction.inspectFuncBeginTransaction(ctx)
	}

	mm_params := &TransactionManagerMockBeginTransactionParams{ctx}

	// Record call args
	mmBeginTransaction.BeginTransactionMock.mutex.Lock()
	mmBeginTransaction.BeginTransactionMock.callArgs = append(mmBeginTransaction.BeginTransactionMock.callArgs, mm_params)
	mmBeginTransaction.BeginTransactionMock.mutex.Unlock()

	for _, e := range mmBeginTransaction.BeginTransactionMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.t1, e.results.err
		}
	}

	if mmBeginTransaction.BeginTransactionMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmBeginTransaction.BeginTransactionMock.defaultExpectation.Counter, 1)
		mm_want := mmBeginTransaction.BeginTransactionMock.defaultExpectation.params
		mm_got := TransactionManagerMockBeginTransactionParams{ctx}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmBeginTransaction.t.Errorf("TransactionManagerMock.BeginTransaction got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmBeginTransaction.BeginTransactionMock.defaultExpectation.results
		if mm_results == nil {
			mmBeginTransaction.t.Fatal("No results are set for the TransactionManagerMock.BeginTransaction")
		}
		return (*mm_results).t1, (*mm_results).err
	}
	if mmBeginTransaction.funcBeginTransaction != nil {
		return mmBeginTransaction.funcBeginTransaction(ctx)
	}
	mmBeginTransaction.t.Fatalf("Unexpected call to TransactionManagerMock.BeginTransaction. %v", ctx)
	return
}

// BeginTransactionAfterCounter returns a count of finished TransactionManagerMock.BeginTransaction invocations
func (mmBeginTransaction *TransactionManagerMock) BeginTransactionAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmBeginTransaction.afterBeginTransactionCounter)
}

// BeginTransactionBeforeCounter returns a count of TransactionManagerMock.BeginTransaction invocations
func (mmBeginTransaction *TransactionManagerMock) BeginTransactionBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmBeginTransaction.beforeBeginTransactionCounter)
}

// Calls returns a list of arguments used in each call to TransactionManagerMock.BeginTransaction.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmBeginTransaction *mTransactionManagerMockBeginTransaction) Calls() []*TransactionManagerMockBeginTransactionParams {
	mmBeginTransaction.mutex.RLock()

	argCopy := make([]*TransactionManagerMockBeginTransactionParams, len(mmBeginTransaction.callArgs))
	copy(argCopy, mmBeginTransaction.callArgs)

	mmBeginTransaction.mutex.RUnlock()

	return argCopy
}

// MinimockBeginTransactionDone returns true if the count of the BeginTransaction invocations corresponds
// the number of defined expectations
func (m *TransactionManagerMock) MinimockBeginTransactionDone() bool {
	for _, e := range m.BeginTransactionMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.BeginTransactionMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterBeginTransactionCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcBeginTransaction != nil && mm_atomic.LoadUint64(&m.afterBeginTransactionCounter) < 1 {
		return false
	}
	return true
}

// MinimockBeginTransactionInspect logs each unmet expectation
func (m *TransactionManagerMock) MinimockBeginTransactionInspect() {
	for _, e := range m.BeginTransactionMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to TransactionManagerMock.BeginTransaction with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.BeginTransactionMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterBeginTransactionCounter) < 1 {
		if m.BeginTransactionMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to TransactionManagerMock.BeginTransaction")
		} else {
			m.t.Errorf("Expected call to TransactionManagerMock.BeginTransaction with params: %#v", *m.BeginTransactionMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcBeginTransaction != nil && mm_atomic.LoadUint64(&m.afterBeginTransactionCounter) < 1 {
		m.t.Error("Expected call to TransactionManagerMock.BeginTransaction")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *TransactionManagerMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockBeginTransactionInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *TransactionManagerMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *TransactionManagerMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockBeginTransactionDone()
}