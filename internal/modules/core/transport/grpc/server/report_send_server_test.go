package core_transport_grpc_test

import (
	"context"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/shav/telegram-bot/internal/modules/core/domain/models"
	"github.com/shav/telegram-bot/internal/modules/core/transport/grpc/api/generated"
	"github.com/shav/telegram-bot/internal/modules/core/transport/grpc/server"
	"github.com/shav/telegram-bot/internal/modules/core/transport/grpc/server/mocks"
)

const userId int64 = 1

var report = core_models.NewReport("Title", "Content")

const reportText = "Title\nContent"

var sendMessageError = errors.New("send message error")

var ctx = context.Background()

func Test_SendReport_ShouldSendMessageWithReportToUser(t *testing.T) {
	messageSenderMock := mocks.NewMessageSenderMock(t)
	messageSenderMock.SendMessageMock.Inspect(func(ctx context.Context, uid int64, text string) {
		assert.Equal(t, userId, uid)
		assert.Equal(t, reportText, text)
	}).Return(nil)

	sendReportServer, err := core_transport_grpc.NewSendReportServer(messageSenderMock)
	assert.NoError(t, err)

	request := &api.SendReportRequest{
		UserId: userId,
		Report: &api.Report{
			Title:   report.Title,
			Content: report.Content,
		},
	}
	_, err = sendReportServer.SendReport(ctx, request)

	assert.NoError(t, err)
	assert.Equal(t, uint64(1), messageSenderMock.SendMessageAfterCounter())
}

func Test_SendReport_ShouldReturnError_WhenSendMessageWithReportFailed(t *testing.T) {
	messageSenderMock := mocks.NewMessageSenderMock(t)
	messageSenderMock.SendMessageMock.Inspect(func(ctx context.Context, uid int64, text string) {
		assert.Equal(t, userId, uid)
		assert.Equal(t, reportText, text)
	}).Return(sendMessageError)

	sendReportServer, err := core_transport_grpc.NewSendReportServer(messageSenderMock)
	assert.NoError(t, err)

	request := &api.SendReportRequest{
		UserId: userId,
		Report: &api.Report{
			Title:   report.Title,
			Content: report.Content,
		},
	}
	_, err = sendReportServer.SendReport(ctx, request)

	assert.ErrorIs(t, err, sendMessageError)
	assert.Equal(t, uint64(1), messageSenderMock.SendMessageAfterCounter())
}
