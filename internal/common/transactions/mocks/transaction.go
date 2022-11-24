package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i github.com/shav/telegram-bot/internal/common/transactions.Transaction -o ./mocks\transaction.go -n TransactionMock

import (
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// TransactionMock implements transactions.Transaction
type TransactionMock struct {
	t minimock.Tester

	funcCommit          func() (err error)
	inspectFuncCommit   func()
	afterCommitCounter  uint64
	beforeCommitCounter uint64
	CommitMock          mTransactionMockCommit

	funcRollback          func() (err error)
	inspectFuncRollback   func()
	afterRollbackCounter  uint64
	beforeRollbackCounter uint64
	RollbackMock          mTransactionMockRollback
}

// NewTransactionMock returns a mock for transactions.Transaction
func NewTransactionMock(t minimock.Tester) *TransactionMock {
	m := &TransactionMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CommitMock = mTransactionMockCommit{mock: m}

	m.RollbackMock = mTransactionMockRollback{mock: m}

	return m
}

type mTransactionMockCommit struct {
	mock               *TransactionMock
	defaultExpectation *TransactionMockCommitExpectation
	expectations       []*TransactionMockCommitExpectation
}

// TransactionMockCommitExpectation specifies expectation struct of the Transaction.Commit
type TransactionMockCommitExpectation struct {
	mock *TransactionMock

	results *TransactionMockCommitResults
	Counter uint64
}

// TransactionMockCommitResults contains results of the Transaction.Commit
type TransactionMockCommitResults struct {
	err error
}

// Expect sets up expected params for Transaction.Commit
func (mmCommit *mTransactionMockCommit) Expect() *mTransactionMockCommit {
	if mmCommit.mock.funcCommit != nil {
		mmCommit.mock.t.Fatalf("TransactionMock.Commit mock is already set by Set")
	}

	if mmCommit.defaultExpectation == nil {
		mmCommit.defaultExpectation = &TransactionMockCommitExpectation{}
	}

	return mmCommit
}

// Inspect accepts an inspector function that has same arguments as the Transaction.Commit
func (mmCommit *mTransactionMockCommit) Inspect(f func()) *mTransactionMockCommit {
	if mmCommit.mock.inspectFuncCommit != nil {
		mmCommit.mock.t.Fatalf("Inspect function is already set for TransactionMock.Commit")
	}

	mmCommit.mock.inspectFuncCommit = f

	return mmCommit
}

// Return sets up results that will be returned by Transaction.Commit
func (mmCommit *mTransactionMockCommit) Return(err error) *TransactionMock {
	if mmCommit.mock.funcCommit != nil {
		mmCommit.mock.t.Fatalf("TransactionMock.Commit mock is already set by Set")
	}

	if mmCommit.defaultExpectation == nil {
		mmCommit.defaultExpectation = &TransactionMockCommitExpectation{mock: mmCommit.mock}
	}
	mmCommit.defaultExpectation.results = &TransactionMockCommitResults{err}
	return mmCommit.mock
}

//Set uses given function f to mock the Transaction.Commit method
func (mmCommit *mTransactionMockCommit) Set(f func() (err error)) *TransactionMock {
	if mmCommit.defaultExpectation != nil {
		mmCommit.mock.t.Fatalf("Default expectation is already set for the Transaction.Commit method")
	}

	if len(mmCommit.expectations) > 0 {
		mmCommit.mock.t.Fatalf("Some expectations are already set for the Transaction.Commit method")
	}

	mmCommit.mock.funcCommit = f
	return mmCommit.mock
}

// Commit implements transactions.Transaction
func (mmCommit *TransactionMock) Commit() (err error) {
	mm_atomic.AddUint64(&mmCommit.beforeCommitCounter, 1)
	defer mm_atomic.AddUint64(&mmCommit.afterCommitCounter, 1)

	if mmCommit.inspectFuncCommit != nil {
		mmCommit.inspectFuncCommit()
	}

	if mmCommit.CommitMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCommit.CommitMock.defaultExpectation.Counter, 1)

		mm_results := mmCommit.CommitMock.defaultExpectation.results
		if mm_results == nil {
			mmCommit.t.Fatal("No results are set for the TransactionMock.Commit")
		}
		return (*mm_results).err
	}
	if mmCommit.funcCommit != nil {
		return mmCommit.funcCommit()
	}
	mmCommit.t.Fatalf("Unexpected call to TransactionMock.Commit.")
	return
}

// CommitAfterCounter returns a count of finished TransactionMock.Commit invocations
func (mmCommit *TransactionMock) CommitAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCommit.afterCommitCounter)
}

// CommitBeforeCounter returns a count of TransactionMock.Commit invocations
func (mmCommit *TransactionMock) CommitBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCommit.beforeCommitCounter)
}

// MinimockCommitDone returns true if the count of the Commit invocations corresponds
// the number of defined expectations
func (m *TransactionMock) MinimockCommitDone() bool {
	for _, e := range m.CommitMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CommitMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCommitCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCommit != nil && mm_atomic.LoadUint64(&m.afterCommitCounter) < 1 {
		return false
	}
	return true
}

// MinimockCommitInspect logs each unmet expectation
func (m *TransactionMock) MinimockCommitInspect() {
	for _, e := range m.CommitMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to TransactionMock.Commit")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CommitMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCommitCounter) < 1 {
		m.t.Error("Expected call to TransactionMock.Commit")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCommit != nil && mm_atomic.LoadUint64(&m.afterCommitCounter) < 1 {
		m.t.Error("Expected call to TransactionMock.Commit")
	}
}

type mTransactionMockRollback struct {
	mock               *TransactionMock
	defaultExpectation *TransactionMockRollbackExpectation
	expectations       []*TransactionMockRollbackExpectation
}

// TransactionMockRollbackExpectation specifies expectation struct of the Transaction.Rollback
type TransactionMockRollbackExpectation struct {
	mock *TransactionMock

	results *TransactionMockRollbackResults
	Counter uint64
}

// TransactionMockRollbackResults contains results of the Transaction.Rollback
type TransactionMockRollbackResults struct {
	err error
}

// Expect sets up expected params for Transaction.Rollback
func (mmRollback *mTransactionMockRollback) Expect() *mTransactionMockRollback {
	if mmRollback.mock.funcRollback != nil {
		mmRollback.mock.t.Fatalf("TransactionMock.Rollback mock is already set by Set")
	}

	if mmRollback.defaultExpectation == nil {
		mmRollback.defaultExpectation = &TransactionMockRollbackExpectation{}
	}

	return mmRollback
}

// Inspect accepts an inspector function that has same arguments as the Transaction.Rollback
func (mmRollback *mTransactionMockRollback) Inspect(f func()) *mTransactionMockRollback {
	if mmRollback.mock.inspectFuncRollback != nil {
		mmRollback.mock.t.Fatalf("Inspect function is already set for TransactionMock.Rollback")
	}

	mmRollback.mock.inspectFuncRollback = f

	return mmRollback
}

// Return sets up results that will be returned by Transaction.Rollback
func (mmRollback *mTransactionMockRollback) Return(err error) *TransactionMock {
	if mmRollback.mock.funcRollback != nil {
		mmRollback.mock.t.Fatalf("TransactionMock.Rollback mock is already set by Set")
	}

	if mmRollback.defaultExpectation == nil {
		mmRollback.defaultExpectation = &TransactionMockRollbackExpectation{mock: mmRollback.mock}
	}
	mmRollback.defaultExpectation.results = &TransactionMockRollbackResults{err}
	return mmRollback.mock
}

//Set uses given function f to mock the Transaction.Rollback method
func (mmRollback *mTransactionMockRollback) Set(f func() (err error)) *TransactionMock {
	if mmRollback.defaultExpectation != nil {
		mmRollback.mock.t.Fatalf("Default expectation is already set for the Transaction.Rollback method")
	}

	if len(mmRollback.expectations) > 0 {
		mmRollback.mock.t.Fatalf("Some expectations are already set for the Transaction.Rollback method")
	}

	mmRollback.mock.funcRollback = f
	return mmRollback.mock
}

// Rollback implements transactions.Transaction
func (mmRollback *TransactionMock) Rollback() (err error) {
	mm_atomic.AddUint64(&mmRollback.beforeRollbackCounter, 1)
	defer mm_atomic.AddUint64(&mmRollback.afterRollbackCounter, 1)

	if mmRollback.inspectFuncRollback != nil {
		mmRollback.inspectFuncRollback()
	}

	if mmRollback.RollbackMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmRollback.RollbackMock.defaultExpectation.Counter, 1)

		mm_results := mmRollback.RollbackMock.defaultExpectation.results
		if mm_results == nil {
			mmRollback.t.Fatal("No results are set for the TransactionMock.Rollback")
		}
		return (*mm_results).err
	}
	if mmRollback.funcRollback != nil {
		return mmRollback.funcRollback()
	}
	mmRollback.t.Fatalf("Unexpected call to TransactionMock.Rollback.")
	return
}

// RollbackAfterCounter returns a count of finished TransactionMock.Rollback invocations
func (mmRollback *TransactionMock) RollbackAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmRollback.afterRollbackCounter)
}

// RollbackBeforeCounter returns a count of TransactionMock.Rollback invocations
func (mmRollback *TransactionMock) RollbackBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmRollback.beforeRollbackCounter)
}

// MinimockRollbackDone returns true if the count of the Rollback invocations corresponds
// the number of defined expectations
func (m *TransactionMock) MinimockRollbackDone() bool {
	for _, e := range m.RollbackMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.RollbackMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterRollbackCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcRollback != nil && mm_atomic.LoadUint64(&m.afterRollbackCounter) < 1 {
		return false
	}
	return true
}

// MinimockRollbackInspect logs each unmet expectation
func (m *TransactionMock) MinimockRollbackInspect() {
	for _, e := range m.RollbackMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to TransactionMock.Rollback")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.RollbackMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterRollbackCounter) < 1 {
		m.t.Error("Expected call to TransactionMock.Rollback")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcRollback != nil && mm_atomic.LoadUint64(&m.afterRollbackCounter) < 1 {
		m.t.Error("Expected call to TransactionMock.Rollback")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *TransactionMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockCommitInspect()

		m.MinimockRollbackInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *TransactionMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *TransactionMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCommitDone() &&
		m.MinimockRollbackDone()
}
