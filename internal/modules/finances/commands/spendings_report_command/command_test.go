package finance_commands_spendings_report_test

import (
	"context"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/shav/telegram-bot/internal/common/date"
	"github.com/shav/telegram-bot/internal/modules/finances/commands/spendings_report_command"
	"github.com/shav/telegram-bot/internal/modules/finances/commands/spendings_report_command/mocks"
)

var userId int64 = 1

var reportingPeriod = date.Periods.Today

var requestReportError = errors.New("request report error")

var ctx = context.Background()

func Test_OnGetSpendingsReport_ShouldAnswerWaitForReport_WhenReportSuccessfullyRequested(t *testing.T) {
	useCaseMock := mocks.NewUseCaseMock(t)
	useCaseMock.RequestSpendingReportMock.Return(nil)

	command, err := finance_commands_spendings_report.NewHandler(finance_commands_spendings_report.Metadata, userId, useCaseMock)
	assert.NoError(t, err)

	_, _, err = command.StartHandleCommand(ctx)
	assert.NoError(t, err)
	answers, _, err := command.HandleNextMessage(ctx, string(reportingPeriod))
	assert.NoError(t, err)

	assert.Equal(t, 1, len(answers))
	assert.Equal(t, "Секундочку, готовим ваш отчёт...", answers[0].Text)
	assert.Empty(t, answers[0].Options)
	assert.Equal(t, uint64(1), useCaseMock.RequestSpendingReportAfterCounter())
}

func Test_OnGetSpendingsReport_ShouldAnswerError_WhenRequestReportFailed(t *testing.T) {
	useCaseMock := mocks.NewUseCaseMock(t)
	useCaseMock.RequestSpendingReportMock.Return(requestReportError)

	command, err := finance_commands_spendings_report.NewHandler(finance_commands_spendings_report.Metadata, userId, useCaseMock)
	assert.NoError(t, err)

	_, _, err = command.StartHandleCommand(ctx)
	assert.NoError(t, err)
	answers, _, err := command.HandleNextMessage(ctx, string(reportingPeriod))

	assert.ErrorIs(t, err, requestReportError)
	assert.Equal(t, 1, len(answers))
	assert.Equal(t, "Не удалось запросить отчёт о тратах: произошла ошибка", answers[0].Text)
	assert.Empty(t, answers[0].Options)
	assert.Equal(t, uint64(1), useCaseMock.RequestSpendingReportAfterCounter())
}
