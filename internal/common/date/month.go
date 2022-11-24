package date

import "time"

// Month представляет из себя структуру для хранения месяца.
type Month struct {
	// Год.
	Year int
	// Месяц.
	Month int
}

// NewMonth создаёт новый экземпляр месяца.
func NewMonth(year int, month int) Month {
	// TODO: Прикрутить нормализацию, либо валидацию некорректных значений.
	return Month{
		Year:  year,
		Month: month,
	}
}

// MonthOf возвращает месяц, в котором находится указанная дата date.
func MonthOf(date Date) Month {
	return Month{
		Year:  date.Year,
		Month: date.Month,
	}
}

// ThisMonth возвращает текущий месяц.
func ThisMonth() Month {
	now := time.Now()
	return NewMonth(now.Year(), int(now.Month()))
}
