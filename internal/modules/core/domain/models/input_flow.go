package core_models

// Пустое значение этапа ввода данных, означающее его отсутствие.
const EmptyInputStage = InputStage("")

// InputStage представляет из себя этап ввода данных.
type InputStage string

// InputFlowMetadata описывает порядок ввода данных в диалоге с чат-ботом.
type InputFlowMetadata struct {
	// Список этапов ввода (в порядке следования).
	stages []InputStage
	// Порядковая таблица этапов
	// (для каждого этапа ввода указан его порядковый номер в общей последовательности ввода).
	stagesOrderTable map[InputStage]int
}

// NewInputFlowMetadata создаёт метаданные схемы ввода данных в диалоге чат-бота.
func NewInputFlowMetadata(stages []InputStage) InputFlowMetadata {
	return InputFlowMetadata{
		stages:           stages,
		stagesOrderTable: makeStagesOrderTable(stages),
	}
}

// GetStages возвращает последовательность этапов.
func (f InputFlowMetadata) GetStages() []InputStage {
	return f.stages
}

// makeStagesOrderTable формирует таблицу, в которой для каждого этапа ввода
// указан его порядковый номер в общей последовательности ввода данных.
func makeStagesOrderTable(stages []InputStage) map[InputStage]int {
	stagesOrderTable := make(map[InputStage]int)
	for i := 0; i < len(stages); i++ {
		stage := stages[i]
		stagesOrderTable[stage] = i
	}
	return stagesOrderTable
}

// GetNextStage возвращает этап ввода данных, следующий после этапа currentStage.
func (f InputFlowMetadata) GetNextStage(currentStage InputStage) InputStage {
	currentStageIndex := f.stagesOrderTable[currentStage]
	nextStageIndex := currentStageIndex + 1
	if nextStageIndex < len(f.stages) {
		return f.stages[nextStageIndex]
	}
	return EmptyInputStage
}
