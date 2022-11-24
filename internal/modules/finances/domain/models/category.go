package finance_models

import (
	"strings"
)

// Categories содержит список категорий трат.
var Categories = categoryEnum{
	Food:           Category{"Food", "Еда"},
	Medicines:      Category{"Medicines", "Лекарства"},
	Clothes:        Category{"Clothes", "Одежда"},
	Transport:      Category{"Transport", "Транспорт"},
	HouseholdGoods: Category{"HouseholdGoods", "Хозтовары"},
	Electronics:    Category{"Electronics", "Электроника"},
	Furniture:      Category{"Furniture", "Мебель"},
	Entertainment:  Category{"Entertainment", "Развлечения"},
	Services:       Category{"Services", "Услуги"},
}

// Category хранит категорию трат.
type Category struct {
	// Системное имя категории.
	Value string
	// Отображаемое имя категории.
	DisplayText string
}

func (c Category) String() string {
	displayText := strings.TrimSpace(c.DisplayText)
	if displayText != "" {
		return displayText
	}
	return strings.TrimSpace(c.Value)
}

// categoryEnum является перчислением категорий трат.
type categoryEnum struct {
	Food           Category
	Medicines      Category
	Clothes        Category
	Transport      Category
	HouseholdGoods Category
	Electronics    Category
	Furniture      Category
	Entertainment  Category
	Services       Category
}

var categoriesByDisplayTexts = map[string]Category{
	strings.ToLower(Categories.Food.DisplayText):           Categories.Food,
	strings.ToLower(Categories.Medicines.DisplayText):      Categories.Medicines,
	strings.ToLower(Categories.Clothes.DisplayText):        Categories.Clothes,
	strings.ToLower(Categories.Transport.DisplayText):      Categories.Transport,
	strings.ToLower(Categories.HouseholdGoods.DisplayText): Categories.HouseholdGoods,
	strings.ToLower(Categories.Electronics.DisplayText):    Categories.Electronics,
	strings.ToLower(Categories.Furniture.DisplayText):      Categories.Furniture,
	strings.ToLower(Categories.Entertainment.DisplayText):  Categories.Entertainment,
	strings.ToLower(Categories.Services.DisplayText):       Categories.Services,
}

var categoriesByValue = map[string]Category{
	strings.ToLower(Categories.Food.Value):           Categories.Food,
	strings.ToLower(Categories.Medicines.Value):      Categories.Medicines,
	strings.ToLower(Categories.Clothes.Value):        Categories.Clothes,
	strings.ToLower(Categories.Transport.Value):      Categories.Transport,
	strings.ToLower(Categories.HouseholdGoods.Value): Categories.HouseholdGoods,
	strings.ToLower(Categories.Electronics.Value):    Categories.Electronics,
	strings.ToLower(Categories.Furniture.Value):      Categories.Furniture,
	strings.ToLower(Categories.Entertainment.Value):  Categories.Entertainment,
	strings.ToLower(Categories.Services.Value):       Categories.Services,
}

// ParseCategory парсит категорию товаров из строки.
func ParseCategory(text string) Category {
	text = strings.TrimSpace(text)
	if category, ok := categoriesByDisplayTexts[strings.ToLower(text)]; ok {
		return category
	}
	if category, ok := categoriesByValue[strings.ToLower(text)]; ok {
		return category
	}
	return Category{Value: text}
}
