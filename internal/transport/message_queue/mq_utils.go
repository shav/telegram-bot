package message_queue

import "fmt"

// getFullQueueName возвращает полное имя очереди в контексте всего приложения.
func getFullQueueName(appName string, queue string) string {
	if appName != "" {
		return fmt.Sprintf("%s_%s", appName, queue)
	}
	return queue
}
