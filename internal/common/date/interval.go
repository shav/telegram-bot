package date

import (
	"fmt"
	"time"
)

// Interval описывает интервал дат (с границами).
type Interval struct {
	// Начало интервала.
	start Date
	// Конец интервала (включительно).
	end Date
}

// NewInterval создаёт новый экземпляр интервала дат.
func NewInterval(start Date, end Date) Interval {
	// TODO: Прикрутить нормализацию интервала, либо валидацию некорректных значений.
	return Interval{
		start: start,
		end:   end,
	}
}

// NewIntervalFromPeriod создаёт интервал дат, соответствующий описанию периода.
func NewIntervalFromPeriod(period Period) Interval {
	now := time.Now()
	today := GetDate(now)
	switch period {
	case Periods.Today:
		return NewInterval(today, today)
	case Periods.ThisWeek:
		return NewInterval(today.StartOfWeek(), today.EndOfWeek())
	case Periods.ThisMonth:
		return NewInterval(today.StartOfMonth(), today.EndOfMonth())
	case Periods.ThisYear:
		return NewInterval(today.StartOfYear(), today.EndOfYear())
	}
	return Interval{}
}

// NewIntervalFromMonth создаёт интервал дат, соответствующий указанному месяцу.
func NewIntervalFromMonth(period Month) Interval {
	startOfMonth := New(period.Year, period.Month, 1)
	return NewInterval(startOfMonth, startOfMonth.EndOfMonth())
}

// Start возвращает начало интервала.
func (i Interval) Start() Date {
	return i.start
}

// End возвращает конец интервала.
func (i Interval) End() Date {
	return i.end
}

// String форматирует интервал дат в строковое представление.
func (i Interval) String() string {
	if i.Start() == i.End() {
		return i.Start().String()
	}
	return fmt.Sprintf("с %s по %s", i.Start().String(), i.End().String())
}

// Contains проверяет, входит ли дата date в указанный интервал дат.
func (i Interval) Contains(date Date) bool {
	dateOrder := date.GetOrderHash()
	return i.Start().GetOrderHash() <= dateOrder && dateOrder <= i.End().GetOrderHash()
}
