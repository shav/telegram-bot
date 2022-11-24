package core_models

// Отчёт.
type Report struct {
	// Заголовок отчёта.
	Title string
	// Содержимое отчёта.
	Content string
}

// NewReport создаёт новый отчёт с содержимым content с заголовком title
func NewReport(title string, content string) Report {
	return Report{
		Title:   title,
		Content: content,
	}
}
