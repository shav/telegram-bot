package metrics

// Init инициализирует счётчики для сбора метрик.
func Init() {
	IncomingMessagesCount = newIncomingMessagesCount()
	IncomingMessageResponseTime = newIncomingMessageResponseTime()
}
