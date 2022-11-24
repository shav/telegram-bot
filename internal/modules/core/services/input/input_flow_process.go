package core_services_input

import (
	"sync"

	"github.com/shav/telegram-bot/internal/modules/core/domain/models"
)

// FlowProcess хранит в себе состояние процесса ввода данных в диалоге с чат-ботом.
type FlowProcess struct {
	// Описание схемы ввода данных.
	metadata core_models.InputFlowMetadata
	// Текущий этап ввода данных.
	currentStage core_models.InputStage
	// Объект синхронизации доступа.
	lock *sync.RWMutex
}

// NewFlowProcess создаёт новый процесс ввода данных.
func NewFlowProcess(metadata core_models.InputFlowMetadata) *FlowProcess {
	return &FlowProcess{
		metadata:     metadata,
		currentStage: core_models.EmptyInputStage,
		lock:         &sync.RWMutex{},
	}
}

// Start начинает процесс ввода данных.
func (f *FlowProcess) Start() {
	f.lock.Lock()
	defer f.lock.Unlock()

	var stages = f.metadata.GetStages()
	if len(stages) > 0 {
		f.currentStage = stages[0]
	}
}

// Reset сбрасывает процесс ввода данных в дефолтное состояние.
func (f *FlowProcess) Reset() {
	f.lock.Lock()
	defer f.lock.Unlock()

	f.currentStage = core_models.EmptyInputStage
}

// GetCurrentStage возвращает текущий этап процесса ввода.
func (f *FlowProcess) GetCurrentStage() core_models.InputStage {
	f.lock.RLock()
	defer f.lock.RUnlock()

	return f.currentStage
}

// GoToNextStage переводит процесса ввода на следующий этап.
func (f *FlowProcess) GoToNextStage() {
	f.lock.Lock()
	defer f.lock.Unlock()

	f.currentStage = f.metadata.GetNextStage(f.currentStage)
}
