package date

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/bmuller/arrow"
	"github.com/goodsign/monday"
)

// daysInWeek задаёт количество дней в неделе.
const daysInWeek = 7

// Ошибка, возникающая в случае если, период не является однодневным.
var notSingleDayPeriodError = errors.New("period is not single day")

// Date представляет из себя структуру для хранения даты.
type Date struct {
	// Год.
	Year int
	// Месяц.
	Month int
	// День.
	Day int
}

// New создаёт новый экземпляр даты.
func New(year int, month int, day int) Date {
	// TODO: Прикрутить нормализацию даты, либо валидацию некорректных значений.
	return Date{
		Year:  year,
		Month: month,
		Day:   day,
	}
}

// NewFromPeriod создаёт новый экземпляр даты на основе однодневного периода.
// Если период не является однодневным, то возвращается ошибка.
func NewFromPeriod(period Period) (Date, error) {
	switch period {
	case Periods.Today:
		return Today(), nil
	case Periods.Yesterday:
		return Yesterday(), nil
	case Periods.DayBeforeYesterday:
		return DayBeforeYesterday(), nil
	default:
		return Date{}, notSingleDayPeriodError
	}
}

// Today возвращает сегодняшнюю дату.
func Today() Date {
	return GetDate(time.Now())
}

// Yesterday возвращает вчерашнюю дату.
func Yesterday() Date {
	return GetDate(time.Now().AddDate(0, 0, -1))
}

// DayBeforeYesterday возвращает позавчерашнюю дату.
func DayBeforeYesterday() Date {
	return GetDate(time.Now().AddDate(0, 0, -2))
}

// GetDate возвращает только дату для даты со временем dateTime.
func GetDate(dateTime time.Time) Date {
	year, month, day := dateTime.Date()
	return New(year, int(month), day)
}

// EndOfDay возвращает дату-время на конец дня.
func (d Date) EndOfDay() time.Time {
	return d.ToDateTime().AddDate(0, 0, 1).Add(-1 * time.Nanosecond)
}

// WeekDay возвращает номер сегодняшнего дня недели в российском формате (ПН=1, ВТ=2, ..., ВС=7)
func (d Date) WeekDay() int {
	dt := d.ToDateTime()
	return ((daysInWeek-1)+int(dt.Weekday()))%daysInWeek + 1
}

// StartOfWeek возвращает начало недели, в которой находится дата.
func (d Date) StartOfWeek() Date {
	dayOfWeek := d.WeekDay()
	return GetDate(d.ToDateTime().AddDate(0, 0, -(dayOfWeek - 1)))
}

// EndOfWeek возвращает конец недели, в которой находится дата.
func (d Date) EndOfWeek() Date {
	dayOfWeek := d.WeekDay()
	return GetDate(d.ToDateTime().AddDate(0, 0, daysInWeek-dayOfWeek))
}

// StartOfMonth возвращает начало месяца, в котором находится дата.
func (d Date) StartOfMonth() Date {
	return New(d.Year, d.Month, 1)
}

// EndOfMonth возвращает конец месяца, в котором находится дата.
func (d Date) EndOfMonth() Date {
	fistDayOfNextMonth := time.Date(d.Year, time.Month(int(d.Month)+1), 1, 0, 0, 0, 0, time.Local)
	return GetDate(fistDayOfNextMonth.Add(-time.Nanosecond))
}

// StartOfYear возвращает начало года, в котором находится дата.
func (d Date) StartOfYear() Date {
	return New(d.Year, 1, 1)
}

// EndOfYear возвращает конец года, в котором находится дата.
func (d Date) EndOfYear() Date {
	firstDayOfNextYear := time.Date(d.Year+1, 1, 1, 0, 0, 0, 0, time.Local)
	return GetDate(firstDayOfNextYear.Add(-time.Nanosecond))
}

// ToDateTime возвращает дату в формате даты со временем.
func (d Date) ToDateTime() time.Time {
	return time.Date(d.Year, time.Month(d.Month), d.Day, 0, 0, 0, 0, time.Local)
}

// String форматирует дату в строковое представление в российском формате.
func (d Date) String() string {
	return fmt.Sprintf("%02d.%02d.%d", d.Day, d.Month, d.Year)
}

// SystemString форматирует дату в строковое представление в системном формате.
func (d Date) SystemString() string {
	return fmt.Sprintf("%d-%02d-%02d", d.Year, d.Month, d.Day)
}

// ParseDate распознаёт дату из строки text.
func ParseDate(text string) (Date, error) {
	text = strings.TrimSpace(strings.ToLower(text))

	period, err := ParsePeriod(text)
	if err == nil {
		return NewFromPeriod(period)
	}

	t, err := arrow.CParse("%d.%m.%Y", text)
	if err != nil {
		// TODO: Данный парсер goodsign/monday не совсем идеальный, можно еще немного доработать его
		d, err := monday.Parse("2 January 2006", text, monday.LocaleRuRU)
		return GetDate(d), err
	}
	return GetDate(t.Time), nil
}

// GetOrderHash возвращает уникальный хэш-код, по которому упорядочены все даты
// т.е. если date1 < date2, то hash(date1) < hash(date2).
func (d Date) GetOrderHash() int {
	return d.Year*366 + d.Month*31 + d.Day
}
