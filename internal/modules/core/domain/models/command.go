package core_models

// Command задаёт название команды.
type Command string

// CommandMetadata представляет из себя метаданные команды для чат-бота.
type CommandMetadata struct {
	// Название команды.
	Name Command
	// Описание.
	Description string
	// Ответы на команду.
	Answers map[AnswerKey]string
	// Опции выбора в ответ на команду, сгруппированные по этапам ввода данных.
	Options map[InputStage][]Option
}

// GetDefaultAnswer возаращает ответ по-умолчанию на команду.
func (c CommandMetadata) GetDefaultAnswer() string {
	return c.Answers[DefaultAnswer]
}
