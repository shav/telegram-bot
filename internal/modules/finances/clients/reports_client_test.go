package finance_clients_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/shav/telegram-bot/internal/common/date"
	"github.com/shav/telegram-bot/internal/modules/finances/clients"
	"github.com/shav/telegram-bot/internal/modules/finances/clients/mocks"
)

const userId = 1

const reportingPeriodName = "За этот месяц"

var reportingDateInterval = date.NewInterval(date.New(2022, 11, 1), date.New(2022, 11, 30))

var serializationError = errors.New("serialization error")

var ctx = context.Background()

func Test_RequestSpendingReport_ShouldNotReturnError_WhenRequestReportMessageSuccessfullySentToMessageQueue(t *testing.T) {
	mqSenderMock := mocks.NewMessageQueueSenderMock(t)
	mqSenderMock.SendMessageAsyncMock.Return()

	serializerMock := mocks.NewSerializerMock(t)
	serializerMock.MarshalMock.Return([]byte("{userId: 1}"), nil)

	reportsClient, err := finance_clients.NewReportsClient(mqSenderMock, serializerMock)
	assert.NoError(t, err)

	err = reportsClient.RequestSpendingReport(ctx, userId, reportingPeriodName, reportingDateInterval)

	assert.NoError(t, err)
	assert.Equal(t, uint64(1), mqSenderMock.SendMessageAsyncAfterCounter())
}

func Test_RequestSpendingReport_ShouldReturnError_AndNotSentMessageToQueue_WhenSerializeRequestReportMessageFailed(t *testing.T) {
	mqSenderMock := mocks.NewMessageQueueSenderMock(t)
	mqSenderMock.SendMessageAsyncMock.Return()

	serializerMock := mocks.NewSerializerMock(t)
	serializerMock.MarshalMock.Return(nil, serializationError)

	reportsClient, err := finance_clients.NewReportsClient(mqSenderMock, serializerMock)
	assert.NoError(t, err)

	err = reportsClient.RequestSpendingReport(ctx, userId, reportingPeriodName, reportingDateInterval)

	assert.ErrorIs(t, err, serializationError)
	assert.Equal(t, uint64(0), mqSenderMock.SendMessageAsyncAfterCounter())
}
