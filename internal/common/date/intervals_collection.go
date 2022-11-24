package date

import "sync"

// IntervalsCollection представляет из себя коллекцию интервалов дат для указанных периодов.
type IntervalsCollection struct {
	// Периоды времени, для которых формируются интервалы дат.
	periods []Period
	// Интервалы дат, соответствующие указанным периодам, актуальные на конкретную дату.
	dateIntervals map[Date][]Interval
	// Объект синхронизации для доступа к коллекции.
	lock *sync.RWMutex
}

// NewIntervalsForPeriods создаёт коллекцию интервалов дат для указанных периодов.
func NewIntervalsForPeriods(periods []Period) *IntervalsCollection {
	return &IntervalsCollection{
		periods:       periods,
		dateIntervals: make(map[Date][]Interval),
		lock:          &sync.RWMutex{},
	}
}

// Get возвращает интервалы дат, актуальные на указанную дату date.
func (c *IntervalsCollection) Get(date Date) []Interval {
	var dateIntervals []Interval
	var exists bool
	func() {
		c.lock.RLock()
		defer c.lock.RUnlock()
		dateIntervals, exists = c.dateIntervals[date]
	}()
	if exists {
		return dateIntervals
	}
	return c.update(date)
}

// update обновляет интервалы дат оносительно указанной даты date.
func (c *IntervalsCollection) update(date Date) []Interval {
	c.lock.Lock()
	defer c.lock.Unlock()

	// Если за время ожидания блокировки другой поток уже успел обновить интервалы, то просто возвращаем готовый результат.
	dateIntervals, exists := c.dateIntervals[date]
	if exists {
		return dateIntervals
	}

	dateIntervals = make([]Interval, 0, len(c.periods))
	for _, period := range c.periods {
		dateIntervals = append(dateIntervals, NewIntervalFromPeriod(period))
	}
	c.dateIntervals[date] = dateIntervals
	return dateIntervals
}
