//go:generate minimock -i spendingStorage -o ./mocks/ -s ".go"
//go:generate minimock -i currencySettings -o ./mocks/ -s ".go"
//go:generate minimock -i currencyConverter -o ./mocks/ -s ".go"

package finance_reports

import (
	"context"
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/shav/telegram-bot/internal/common/date"
	tr "github.com/shav/telegram-bot/internal/common/transactions"
	"github.com/shav/telegram-bot/internal/modules/core/domain/models"
	"github.com/shav/telegram-bot/internal/modules/finances/domain/models"
	"github.com/shav/telegram-bot/internal/observability/logger"
	"github.com/shav/telegram-bot/internal/observability/tracing"
)

// Порядок сортировки таблицы трат.
var spendingsTableSort = finance_models.NewTableSortMode(finance_models.SpendingsTableColumns.Category, finance_models.Sort.Asc)

// currencySettings хранит валютные настройки.
type currencySettings interface {
	// GetCurrency возвращает текущую валюту для пользователя userId.
	GetCurrency(ctx context.Context, ts tr.Transaction, userId int64) (finance_models.Currency, error)
}

// spendingStorage является хранилищем трат.
type spendingStorage interface {
	// GetSpendingsByCategories возвращает отчёт по тратам пользователя userId
	// за указанный промежуток времени interval, сгруппированный по категориям.
	GetSpendingsByCategories(ctx context.Context, ts tr.Transaction, userId int64, interval date.Interval) (finance_models.SpendingsByCategoryTable, error)
}

// currencyConverter выполняет конвертацию валют.
type currencyConverter interface {
	// ConvertSpendingsTableToUserCurrency конвертирует таблицу расходов spendingsTable в валюту пользователя userId.
	ConvertSpendingsTableToUserCurrency(ctx context.Context, userId int64, spendingsTable finance_models.SpendingsByCategoryTable) (finance_models.SpendingsByCategoryTable, error)
}

// SpendingReportBuilder выполняет построение финансовых отчётов о тратах.
type SpendingReportBuilder struct {
	// Валютные настройки.
	currencySettings currencySettings
	// Хранилище трат.
	spendingStorage spendingStorage
	// Конвертер валют.
	converter currencyConverter
}

// NewSpendingReportBuilder создаёт построителя отчётов о тратах.
func NewSpendingReportBuilder(currencySettings currencySettings, spendingStorage spendingStorage, converter currencyConverter) (*SpendingReportBuilder, error) {
	if currencySettings == nil {
		return nil, errors.New("New SpendingReportBuilder: currency settings is not assigned")
	}
	if spendingStorage == nil {
		return nil, errors.New("New SpendingReportBuilder: spending storage is not assigned")
	}
	if converter == nil {
		return nil, errors.New("New SpendingReportBuilder: currency converter is not assigned")
	}

	return &SpendingReportBuilder{
		currencySettings: currencySettings,
		spendingStorage:  spendingStorage,
		converter:        converter,
	}, nil
}

// GetSpendingReport возвращает отчёт о тратах пользователя userId за указанный период времени dateInterval.
func (r *SpendingReportBuilder) GetSpendingReport(ctx context.Context, userId int64, periodName string, dateInterval date.Interval) (core_models.Report, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "SpendingReportBuilder.GetSpendingReport")
	defer span.Finish()

	spendingsTable, err := r.getSpendingsByCategoriesTable(ctx, userId, dateInterval)
	if err != nil {
		err = errors.Wrap(err, "Get user spendings by categories table failed")
		tracing.SetError(span)

		var cce *finance_models.CurrencyConvertError
		var errorReason string
		if errors.As(err, &cce) {
			errorReason = spendingReportTexts.cannotConvertCurrency
		} else {
			errorReason = spendingReportTexts.cannotGetSpendings
		}
		reportTitle := fmt.Sprintf(spendingReportTexts.makeSpendingsReportFailedTemplate, strings.ToLower(periodName))
		return core_models.NewReport(reportTitle, errorReason), err
	}

	var reportContent string
	if len(spendingsTable) == 0 {
		reportContent = fmt.Sprintf(spendingReportTexts.noSpendings)
	} else {
		reportContent = spendingsTable
	}
	reportTitle := fmt.Sprintf(spendingReportTexts.allSpendingsTemplate, strings.ToLower(periodName))
	return core_models.NewReport(reportTitle, reportContent), nil
}

// getSpendingsByCategoriesTable возвращает таблицу с тратами пользователя userId
// за указанный период времени dateInterval, сгруппированный по категориям.
func (r *SpendingReportBuilder) getSpendingsByCategoriesTable(ctx context.Context, userId int64, dateInterval date.Interval) (string, error) {
	spendingsTable, err := r.spendingStorage.GetSpendingsByCategories(ctx, nil, userId, dateInterval)
	if err != nil {
		return "", err
	}

	spendingsTable, err = r.converter.ConvertSpendingsTableToUserCurrency(ctx, userId, spendingsTable)
	if err != nil {
		return "", finance_models.NewCurrencyConvertError(err)
	}

	userCurrency, err := r.currencySettings.GetCurrency(ctx, nil, userId)
	if err != nil {
		logger.Error(ctx, "Get user currency failed", logger.Fields.Error(err))
	}

	report := spendingsTable.String(userCurrency, spendingsTableSort)
	return report, nil
}
