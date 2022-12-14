package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i github.com/shav/telegram-bot/internal/modules/finances/clients.messageQueueSender -o ./mocks\message_queue_sender.go -n MessageQueueSenderMock

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// MessageQueueSenderMock implements finance_clients.messageQueueSender
type MessageQueueSenderMock struct {
	t minimock.Tester

	funcSendMessageAsync          func(ctx context.Context, queue string, key string, payload []byte)
	inspectFuncSendMessageAsync   func(ctx context.Context, queue string, key string, payload []byte)
	afterSendMessageAsyncCounter  uint64
	beforeSendMessageAsyncCounter uint64
	SendMessageAsyncMock          mMessageQueueSenderMockSendMessageAsync
}

// NewMessageQueueSenderMock returns a mock for finance_clients.messageQueueSender
func NewMessageQueueSenderMock(t minimock.Tester) *MessageQueueSenderMock {
	m := &MessageQueueSenderMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.SendMessageAsyncMock = mMessageQueueSenderMockSendMessageAsync{mock: m}
	m.SendMessageAsyncMock.callArgs = []*MessageQueueSenderMockSendMessageAsyncParams{}

	return m
}

type mMessageQueueSenderMockSendMessageAsync struct {
	mock               *MessageQueueSenderMock
	defaultExpectation *MessageQueueSenderMockSendMessageAsyncExpectation
	expectations       []*MessageQueueSenderMockSendMessageAsyncExpectation

	callArgs []*MessageQueueSenderMockSendMessageAsyncParams
	mutex    sync.RWMutex
}

// MessageQueueSenderMockSendMessageAsyncExpectation specifies expectation struct of the messageQueueSender.SendMessageAsync
type MessageQueueSenderMockSendMessageAsyncExpectation struct {
	mock   *MessageQueueSenderMock
	params *MessageQueueSenderMockSendMessageAsyncParams

	Counter uint64
}

// MessageQueueSenderMockSendMessageAsyncParams contains parameters of the messageQueueSender.SendMessageAsync
type MessageQueueSenderMockSendMessageAsyncParams struct {
	ctx     context.Context
	queue   string
	key     string
	payload []byte
}

// Expect sets up expected params for messageQueueSender.SendMessageAsync
func (mmSendMessageAsync *mMessageQueueSenderMockSendMessageAsync) Expect(ctx context.Context, queue string, key string, payload []byte) *mMessageQueueSenderMockSendMessageAsync {
	if mmSendMessageAsync.mock.funcSendMessageAsync != nil {
		mmSendMessageAsync.mock.t.Fatalf("MessageQueueSenderMock.SendMessageAsync mock is already set by Set")
	}

	if mmSendMessageAsync.defaultExpectation == nil {
		mmSendMessageAsync.defaultExpectation = &MessageQueueSenderMockSendMessageAsyncExpectation{}
	}

	mmSendMessageAsync.defaultExpectation.params = &MessageQueueSenderMockSendMessageAsyncParams{ctx, queue, key, payload}
	for _, e := range mmSendMessageAsync.expectations {
		if minimock.Equal(e.params, mmSendMessageAsync.defaultExpectation.params) {
			mmSendMessageAsync.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmSendMessageAsync.defaultExpectation.params)
		}
	}

	return mmSendMessageAsync
}

// Inspect accepts an inspector function that has same arguments as the messageQueueSender.SendMessageAsync
func (mmSendMessageAsync *mMessageQueueSenderMockSendMessageAsync) Inspect(f func(ctx context.Context, queue string, key string, payload []byte)) *mMessageQueueSenderMockSendMessageAsync {
	if mmSendMessageAsync.mock.inspectFuncSendMessageAsync != nil {
		mmSendMessageAsync.mock.t.Fatalf("Inspect function is already set for MessageQueueSenderMock.SendMessageAsync")
	}

	mmSendMessageAsync.mock.inspectFuncSendMessageAsync = f

	return mmSendMessageAsync
}

// Return sets up results that will be returned by messageQueueSender.SendMessageAsync
func (mmSendMessageAsync *mMessageQueueSenderMockSendMessageAsync) Return() *MessageQueueSenderMock {
	if mmSendMessageAsync.mock.funcSendMessageAsync != nil {
		mmSendMessageAsync.mock.t.Fatalf("MessageQueueSenderMock.SendMessageAsync mock is already set by Set")
	}

	if mmSendMessageAsync.defaultExpectation == nil {
		mmSendMessageAsync.defaultExpectation = &MessageQueueSenderMockSendMessageAsyncExpectation{mock: mmSendMessageAsync.mock}
	}

	return mmSendMessageAsync.mock
}

//Set uses given function f to mock the messageQueueSender.SendMessageAsync method
func (mmSendMessageAsync *mMessageQueueSenderMockSendMessageAsync) Set(f func(ctx context.Context, queue string, key string, payload []byte)) *MessageQueueSenderMock {
	if mmSendMessageAsync.defaultExpectation != nil {
		mmSendMessageAsync.mock.t.Fatalf("Default expectation is already set for the messageQueueSender.SendMessageAsync method")
	}

	if len(mmSendMessageAsync.expectations) > 0 {
		mmSendMessageAsync.mock.t.Fatalf("Some expectations are already set for the messageQueueSender.SendMessageAsync method")
	}

	mmSendMessageAsync.mock.funcSendMessageAsync = f
	return mmSendMessageAsync.mock
}

// SendMessageAsync implements finance_clients.messageQueueSender
func (mmSendMessageAsync *MessageQueueSenderMock) SendMessageAsync(ctx context.Context, queue string, key string, payload []byte) {
	mm_atomic.AddUint64(&mmSendMessageAsync.beforeSendMessageAsyncCounter, 1)
	defer mm_atomic.AddUint64(&mmSendMessageAsync.afterSendMessageAsyncCounter, 1)

	if mmSendMessageAsync.inspectFuncSendMessageAsync != nil {
		mmSendMessageAsync.inspectFuncSendMessageAsync(ctx, queue, key, payload)
	}

	mm_params := &MessageQueueSenderMockSendMessageAsyncParams{ctx, queue, key, payload}

	// Record call args
	mmSendMessageAsync.SendMessageAsyncMock.mutex.Lock()
	mmSendMessageAsync.SendMessageAsyncMock.callArgs = append(mmSendMessageAsync.SendMessageAsyncMock.callArgs, mm_params)
	mmSendMessageAsync.SendMessageAsyncMock.mutex.Unlock()

	for _, e := range mmSendMessageAsync.SendMessageAsyncMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return
		}
	}

	if mmSendMessageAsync.SendMessageAsyncMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSendMessageAsync.SendMessageAsyncMock.defaultExpectation.Counter, 1)
		mm_want := mmSendMessageAsync.SendMessageAsyncMock.defaultExpectation.params
		mm_got := MessageQueueSenderMockSendMessageAsyncParams{ctx, queue, key, payload}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmSendMessageAsync.t.Errorf("MessageQueueSenderMock.SendMessageAsync got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		return

	}
	if mmSendMessageAsync.funcSendMessageAsync != nil {
		mmSendMessageAsync.funcSendMessageAsync(ctx, queue, key, payload)
		return
	}
	mmSendMessageAsync.t.Fatalf("Unexpected call to MessageQueueSenderMock.SendMessageAsync. %v %v %v %v", ctx, queue, key, payload)

}

// SendMessageAsyncAfterCounter returns a count of finished MessageQueueSenderMock.SendMessageAsync invocations
func (mmSendMessageAsync *MessageQueueSenderMock) SendMessageAsyncAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSendMessageAsync.afterSendMessageAsyncCounter)
}

// SendMessageAsyncBeforeCounter returns a count of MessageQueueSenderMock.SendMessageAsync invocations
func (mmSendMessageAsync *MessageQueueSenderMock) SendMessageAsyncBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSendMessageAsync.beforeSendMessageAsyncCounter)
}

// Calls returns a list of arguments used in each call to MessageQueueSenderMock.SendMessageAsync.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmSendMessageAsync *mMessageQueueSenderMockSendMessageAsync) Calls() []*MessageQueueSenderMockSendMessageAsyncParams {
	mmSendMessageAsync.mutex.RLock()

	argCopy := make([]*MessageQueueSenderMockSendMessageAsyncParams, len(mmSendMessageAsync.callArgs))
	copy(argCopy, mmSendMessageAsync.callArgs)

	mmSendMessageAsync.mutex.RUnlock()

	return argCopy
}

// MinimockSendMessageAsyncDone returns true if the count of the SendMessageAsync invocations corresponds
// the number of defined expectations
func (m *MessageQueueSenderMock) MinimockSendMessageAsyncDone() bool {
	for _, e := range m.SendMessageAsyncMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SendMessageAsyncMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSendMessageAsyncCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSendMessageAsync != nil && mm_atomic.LoadUint64(&m.afterSendMessageAsyncCounter) < 1 {
		return false
	}
	return true
}

// MinimockSendMessageAsyncInspect logs each unmet expectation
func (m *MessageQueueSenderMock) MinimockSendMessageAsyncInspect() {
	for _, e := range m.SendMessageAsyncMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to MessageQueueSenderMock.SendMessageAsync with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SendMessageAsyncMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSendMessageAsyncCounter) < 1 {
		if m.SendMessageAsyncMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to MessageQueueSenderMock.SendMessageAsync")
		} else {
			m.t.Errorf("Expected call to MessageQueueSenderMock.SendMessageAsync with params: %#v", *m.SendMessageAsyncMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSendMessageAsync != nil && mm_atomic.LoadUint64(&m.afterSendMessageAsyncCounter) < 1 {
		m.t.Error("Expected call to MessageQueueSenderMock.SendMessageAsync")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *MessageQueueSenderMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockSendMessageAsyncInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *MessageQueueSenderMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *MessageQueueSenderMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockSendMessageAsyncDone()
}
