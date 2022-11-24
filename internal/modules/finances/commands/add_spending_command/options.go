package finance_commands_add_spending

import (
	"github.com/shav/telegram-bot/internal/common/date"
	"github.com/shav/telegram-bot/internal/modules/finances/domain/models"
	"github.com/shav/telegram-bot/internal/modules/core/domain/models"
)

// Формат отображения однодневных периодов времени.
var dayFormat = date.PeriodDisplayFormats.In

// Опции команды добавления трат.
var options = map[core_models.InputStage][]core_models.Option{
	inputStages.date: {
		{Value: string(date.Periods.Today), Text: date.Periods.Today.String(dayFormat)},
		{Value: string(date.Periods.Yesterday), Text: date.Periods.Yesterday.String(dayFormat)},
		{Value: string(date.Periods.DayBeforeYesterday), Text: date.Periods.DayBeforeYesterday.String(dayFormat)},
		// TODO: Добавить дни недели (для текущей недели)
	},

	inputStages.category: {
		{Value: finance_models.Categories.Food.Value, Text: finance_models.Categories.Food.DisplayText},
		{Value: finance_models.Categories.Clothes.Value, Text: finance_models.Categories.Clothes.DisplayText},
		{Value: finance_models.Categories.Medicines.Value, Text: finance_models.Categories.Medicines.DisplayText},
		{Value: finance_models.Categories.Electronics.Value, Text: finance_models.Categories.Electronics.DisplayText},
		{Value: finance_models.Categories.Furniture.Value, Text: finance_models.Categories.Furniture.DisplayText},
		{Value: finance_models.Categories.HouseholdGoods.Value, Text: finance_models.Categories.HouseholdGoods.DisplayText},
		{Value: finance_models.Categories.Transport.Value, Text: finance_models.Categories.Transport.DisplayText},
		{Value: finance_models.Categories.Entertainment.Value, Text: finance_models.Categories.Entertainment.DisplayText},
		{Value: finance_models.Categories.Services.Value, Text: finance_models.Categories.Services.DisplayText},
	},
}
