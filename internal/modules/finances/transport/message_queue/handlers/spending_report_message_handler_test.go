package finance_transport_mq_handlers_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/shav/telegram-bot/internal/common/date"
	"github.com/shav/telegram-bot/internal/modules/core/domain/models"
	"github.com/shav/telegram-bot/internal/modules/finances/transport/message_queue/handlers"
	"github.com/shav/telegram-bot/internal/modules/finances/transport/message_queue/handlers/mocks"
	"github.com/shav/telegram-bot/internal/modules/finances/transport/message_queue/models"
)

const messageKey = "Key"

var messageData = []byte(`
{
  UserId: 1
  PeriodName: "Сегодня"
  StartDate: { "Year":2022, "Month":11, "Day":13 }
  EndDate: { "Year":2022, "Month":11, "Day":13 }
}`)

var reportingDateInterval = date.NewInterval(date.New(2022, 11, 13), date.New(2022, 11, 13))

const userId int64 = 1

const reportingPeriod = "Сегодня"

var message = finance_transport_mq_models.NewSpendingReportRequestMessage(userId, reportingPeriod,
	reportingDateInterval.Start(), reportingDateInterval.End())

var report = core_models.NewReport("Title", "Content")

var deserializationError = errors.New("deserialization error")
var reportBuildError = errors.New("report build error")
var reportSendError = errors.New("report send error")

var ctx = context.Background()

func Test_HandleMessage_ShouldBuildAndSendReportToUser(t *testing.T) {
	deserializerMock := mocks.NewDeserializerMock(t)
	deserializerMock.UnmarshalMock.Inspect(func(serializedObj []byte, object any) {
		assert.Equal(t, messageData, serializedObj)
		pMessage := object.(*finance_transport_mq_models.SpendingReportRequestMessage)
		*pMessage = message
	}).Return(nil)

	reportBuilderMock := mocks.NewSpendingReportBuilderMock(t)
	reportBuilderMock.GetSpendingReportMock.Inspect(func(ctx context.Context, uid int64, periodName string, dateInterval date.Interval) {
		assert.Equal(t, userId, uid)
		assert.Equal(t, reportingPeriod, periodName)
		assert.Equal(t, reportingDateInterval, dateInterval)
	}).Return(report, nil)

	reportSenderMock := mocks.NewReportSenderMock(t)
	reportSenderMock.SendReportMock.Inspect(func(ctx context.Context, uid int64, r core_models.Report) {
		assert.Equal(t, userId, uid)
		assert.Equal(t, report, r)
	}).Return(nil)

	messageHandler, err := finance_transport_mq_handlers.NewSpendingReportMessageHandler(reportBuilderMock, reportSenderMock, deserializerMock)
	assert.NoError(t, err)

	err = messageHandler.HandleMessage(ctx, messageKey, messageData)

	assert.NoError(t, err)
	assert.Equal(t, uint64(1), deserializerMock.UnmarshalAfterCounter())
	assert.Equal(t, uint64(1), reportBuilderMock.GetSpendingReportAfterCounter())
	assert.Equal(t, uint64(1), reportSenderMock.SendReportAfterCounter())
}

func Test_HandleMessage_ShouldReturnError_WhenDeserializeMessageFailed(t *testing.T) {
	deserializerMock := mocks.NewDeserializerMock(t)
	deserializerMock.UnmarshalMock.Return(deserializationError)

	reportBuilderMock := mocks.NewSpendingReportBuilderMock(t)
	reportSenderMock := mocks.NewReportSenderMock(t)

	messageHandler, err := finance_transport_mq_handlers.NewSpendingReportMessageHandler(reportBuilderMock, reportSenderMock, deserializerMock)
	assert.NoError(t, err)

	err = messageHandler.HandleMessage(ctx, messageKey, messageData)

	assert.ErrorIs(t, err, deserializationError)
	assert.Equal(t, uint64(1), deserializerMock.UnmarshalAfterCounter())
	assert.Equal(t, uint64(0), reportBuilderMock.GetSpendingReportAfterCounter())
	assert.Equal(t, uint64(0), reportSenderMock.SendReportAfterCounter())
}

func Test_HandleMessage_ShouldReturnError_WhenBuildSpendingReportFailed(t *testing.T) {
	deserializerMock := mocks.NewDeserializerMock(t)
	deserializerMock.UnmarshalMock.Inspect(func(serializedObj []byte, object any) {
		assert.Equal(t, messageData, serializedObj)
		pMessage := object.(*finance_transport_mq_models.SpendingReportRequestMessage)
		*pMessage = message
	}).Return(nil)

	reportBuilderMock := mocks.NewSpendingReportBuilderMock(t)
	reportBuilderMock.GetSpendingReportMock.Inspect(func(ctx context.Context, uid int64, periodName string, dateInterval date.Interval) {
		assert.Equal(t, userId, uid)
		assert.Equal(t, reportingPeriod, periodName)
		assert.Equal(t, reportingDateInterval, dateInterval)
	}).Return(report, reportBuildError)

	reportSenderMock := mocks.NewReportSenderMock(t)

	messageHandler, err := finance_transport_mq_handlers.NewSpendingReportMessageHandler(reportBuilderMock, reportSenderMock, deserializerMock)
	assert.NoError(t, err)

	err = messageHandler.HandleMessage(ctx, messageKey, messageData)

	assert.ErrorIs(t, err, reportBuildError)
	assert.Equal(t, uint64(1), deserializerMock.UnmarshalAfterCounter())
	assert.Equal(t, uint64(1), reportBuilderMock.GetSpendingReportAfterCounter())
	assert.Equal(t, uint64(0), reportSenderMock.SendReportAfterCounter())
}

func Test_HandleMessage_ShouldReturnError_WhenSendReportToUserFailed(t *testing.T) {
	deserializerMock := mocks.NewDeserializerMock(t)
	deserializerMock.UnmarshalMock.Inspect(func(serializedObj []byte, object any) {
		assert.Equal(t, messageData, serializedObj)
		pMessage := object.(*finance_transport_mq_models.SpendingReportRequestMessage)
		*pMessage = message
	}).Return(nil)

	reportBuilderMock := mocks.NewSpendingReportBuilderMock(t)
	reportBuilderMock.GetSpendingReportMock.Inspect(func(ctx context.Context, uid int64, periodName string, dateInterval date.Interval) {
		assert.Equal(t, userId, uid)
		assert.Equal(t, reportingPeriod, periodName)
		assert.Equal(t, reportingDateInterval, dateInterval)
	}).Return(report, nil)

	reportSenderMock := mocks.NewReportSenderMock(t)
	reportSenderMock.SendReportMock.Inspect(func(ctx context.Context, uid int64, r core_models.Report) {
		assert.Equal(t, userId, uid)
		assert.Equal(t, report, r)
	}).Return(reportSendError)

	messageHandler, err := finance_transport_mq_handlers.NewSpendingReportMessageHandler(reportBuilderMock, reportSenderMock, deserializerMock)
	assert.NoError(t, err)

	err = messageHandler.HandleMessage(ctx, messageKey, messageData)

	assert.ErrorIs(t, err, reportSendError)
	assert.Equal(t, uint64(1), deserializerMock.UnmarshalAfterCounter())
	assert.Equal(t, uint64(1), reportBuilderMock.GetSpendingReportAfterCounter())
	assert.Equal(t, uint64(1), reportSenderMock.SendReportAfterCounter())
}
