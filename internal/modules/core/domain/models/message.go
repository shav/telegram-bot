package core_models

// Message - это сообщение из какого-либо мессенджера.
type Message struct {
	// Текст сообщения.
	Text string
	// ИД пользователя.
	UserID int64
}
