package numbers

import (
	"strings"
	
	"github.com/shopspring/decimal"
)

// ParseDecimal распознаёт из строки text точное вещественное число.
func ParseDecimal(text string) (decimal.Decimal, error) {
	text = strings.TrimSpace(text)
	text = strings.ReplaceAll(text, ",", ".")
	return decimal.NewFromString(text)
}
