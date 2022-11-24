package finance_models

import (
	"fmt"
	"sort"
	"strings"

	"github.com/shopspring/decimal"
)

// SpendingsTableColumns содержит список колонок таблицы с тратами.
var SpendingsTableColumns = spendingsTableColumnsEnum{
	// Категория трат.
	Category: Column("Category"),
	// Сумма траты.
	Amount: Column("Amount"),
}

// spendingsTableColumnsEnum перечисляет колонки таблицы с тратами.
type spendingsTableColumnsEnum struct {
	Category Column
	Amount   Column
}

// SpendingsByCategoryTable представляет из себя таблицу трат по категориям.
type SpendingsByCategoryTable map[Category]decimal.Decimal

// String форматирует таблицу с тратами.
func (t SpendingsByCategoryTable) String(currency Currency, sortMode TableSortMode) string {
	sb := strings.Builder{}
	categories := t.Sort(sortMode)
	for i, category := range categories {
		amount := t[category]
		var lineBreak string
		if i < len(categories)-1 {
			lineBreak = "\n"
		}
		sb.WriteString(fmt.Sprintf("%s:  %s%s", category, FormatMoneyWithCurrency(amount, currency), lineBreak))
	}
	return sb.String()
}

// ********************************************************************************************************************
// Сортировка
// ********************************************************************************************************************

type categoriesByDisplayText []Category

func (s categoriesByDisplayText) Len() int           { return len(s) }
func (s categoriesByDisplayText) Less(i, j int) bool { return s[i].String() < s[j].String() }
func (s categoriesByDisplayText) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// Sort сортирует таблицу трат в соотвествии с указанными настройками sortMode,
// и возвращает список категорий в порядке сортировки.
func (t SpendingsByCategoryTable) Sort(sortMode TableSortMode) []Category {
	categories := make([]Category, 0, len(t))
	for category := range t {
		categories = append(categories, category)
	}

	switch strings.ToLower(string(sortMode.Column)) {
	case strings.ToLower(string(SpendingsTableColumns.Category)):
		switch sortMode.Order {
		case Sort.Asc:
			sort.Sort(categoriesByDisplayText(categories))
		case Sort.Desc:
			sort.Sort(sort.Reverse(categoriesByDisplayText(categories)))
		}
	case strings.ToLower(string(SpendingsTableColumns.Amount)):
		// TODO: Прикрутить сортировку по размеру трат.
		panic("Not implemented")
	}
	return categories
}
