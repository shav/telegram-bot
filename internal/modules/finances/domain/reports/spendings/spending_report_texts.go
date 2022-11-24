package finance_reports

// Тексты для отчёта по тратам.
var spendingReportTexts = spendingReportTextsEnum{
	allSpendingsTemplate:              "Статистика по тратам %s:",
	noSpendings:                       "Трат пока нет",
	makeSpendingsReportFailedTemplate: "При формировании отчёта о тратах %s произошла ошибка:",
	cannotConvertCurrency:             "Не удалось выполнить конвертацию валюты",
	cannotGetSpendings:                "Не удалось получить траты",
}

// spendingReportTextsEnum перечисляет тексты для отчёта по тратам.
type spendingReportTextsEnum struct {
	noSpendings                       string
	allSpendingsTemplate              string
	makeSpendingsReportFailedTemplate string
	cannotConvertCurrency             string
	cannotGetSpendings                string
}
