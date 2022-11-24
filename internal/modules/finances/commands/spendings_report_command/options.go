package finance_commands_spendings_report

import (
	"github.com/shav/telegram-bot/internal/common/date"
	"github.com/shav/telegram-bot/internal/modules/core/domain/models"
)

// Формат отображения периодов времени.
var periodFormat = date.PeriodDisplayFormats.For

// Опции команды составления отчёта по тратам.
var options = map[core_models.InputStage][]core_models.Option{
	inputStages.period: {
		{Value: string(date.Periods.Today), Text: date.Periods.Today.String(periodFormat)},
		{Value: string(date.Periods.ThisWeek), Text: date.Periods.ThisWeek.String(periodFormat)},
		{Value: string(date.Periods.ThisMonth), Text: date.Periods.ThisMonth.String(periodFormat)},
		{Value: string(date.Periods.ThisYear), Text: date.Periods.ThisYear.String(periodFormat)},
	},
}

// GetSpendingReportPeriods возвращает список периодов времени, за которые можно составлять отчёты по тратам.
func GetSpendingReportPeriods() []date.Period {
	reportPeriods := make([]date.Period, 0)
	for _, period := range Metadata.Options[inputStages.period] {
		reportPeriods = append(reportPeriods, date.Period(period.Value))
	}
	return reportPeriods
}
