package finance_models

// Sort содержит список возможных направлений сортировки.
var Sort = sortOrderEnum{
	// По возрастанию.
	Asc: SortOrder("Asc"),
	// По убыванию.
	Desc: SortOrder("Desc"),
}

// SortOrder задаёт направление сортировки.
type SortOrder string

// sortOrderEnum перечисляет возможные направления сортировки.
type sortOrderEnum struct {
	Asc  SortOrder
	Desc SortOrder
}

// Column хранит название колонки таблицы.
type Column string

// TableSortMode описывает настройки сортировки таблицы.
type TableSortMode struct {
	// Колонка для сортироовки.
	Column Column
	// Направление сортировки.
	Order SortOrder
}

// NewTableSortMode создаёт настройки сортировки таблицы.
func NewTableSortMode(column Column, order SortOrder) TableSortMode {
	return TableSortMode{
		Column: column,
		Order:  order,
	}
}
