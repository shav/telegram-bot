package date

import (
	"errors"
	"strings"
)

// Ошибка форматирования периода.
var periodFormatError = errors.New("period string has wrong format")

// Periods хранит список возможных периодов времени.
var Periods = periodEnum{
	// Сегодня.
	Today: Period("Today"),
	// Вчера
	Yesterday: Period("Yesterday"),
	// Позавчера
	DayBeforeYesterday: Period("DayBeforeYesterday"),
	// На этой неделе.
	ThisWeek: Period("ThisWeek"),
	// В этом месяце.
	ThisMonth: Period("ThisMonth"),
	// В этом году.
	ThisYear: Period("ThisYear"),
}

// PeriodDisplayFormats хранит форматы отображения периода.
var PeriodDisplayFormats = periodDisplayFormatEnum{
	// В какой-то промежуток времени.
	In: PeriodDisplayFormat("In"),
	// За какой-то промежуток времени.
	For: PeriodDisplayFormat("For"),
}

// Отображаемые значения периодов в различных форматах.
var periodDisplayValues = map[Period]map[PeriodDisplayFormat]string{
	Periods.Today: {
		PeriodDisplayFormats.In:  "Сегодня",
		PeriodDisplayFormats.For: "За сегодня",
	},
	Periods.Yesterday: {
		PeriodDisplayFormats.In:  "Вчера",
		PeriodDisplayFormats.For: "За вчера",
	},
	Periods.DayBeforeYesterday: {
		PeriodDisplayFormats.In:  "Позавчера",
		PeriodDisplayFormats.For: "За позавчера",
	},
	Periods.ThisWeek: {
		PeriodDisplayFormats.In:  "На этой неделе",
		PeriodDisplayFormats.For: "За эту неделю",
	},
	Periods.ThisMonth: {
		PeriodDisplayFormats.In:  "В этом месяце",
		PeriodDisplayFormats.For: "За этот месяц",
	},
	Periods.ThisYear: {
		PeriodDisplayFormats.In:  "В этом году",
		PeriodDisplayFormats.For: "За этот год",
	},
}

// Period задаёт период времени.
type Period string

// PeriodDisplayFormat задаёт формат отображения периода.
type PeriodDisplayFormat string

// periodEnum является перечислением периодов времени.
type periodEnum struct {
	Today              Period
	Yesterday          Period
	DayBeforeYesterday Period
	ThisWeek           Period
	ThisMonth          Period
	ThisYear           Period
}

// periodDisplayFormatEnum является перечислением форматов отображения периода.
type periodDisplayFormatEnum struct {
	In  PeriodDisplayFormat
	For PeriodDisplayFormat
}

// String форматирует название периода времени.
func (p Period) String(format PeriodDisplayFormat) string {
	if displayValues, ok := periodDisplayValues[p]; ok {
		return displayValues[format]
	}
	return ""
}

// ParsePeriod распознаёт период дат из строки text.
func ParsePeriod(text string) (Period, error) {
	text = strings.TrimSpace(strings.ToLower(text))

	// TODO: Прикрутить парсинг других словесных обозначений дат.
	// Например, "в этот понедельник", "в прошлый вторник" и т.п.
	for period, displayValues := range periodDisplayValues {
		switch text {
		case strings.ToLower(string(period)):
			fallthrough
		case strings.ToLower(displayValues[PeriodDisplayFormats.In]):
			fallthrough
		case strings.ToLower(displayValues[PeriodDisplayFormats.For]):
			return period, nil
		}
	}
	return Period(""), periodFormatError
}
