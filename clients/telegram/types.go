package telegram

// UpdateResponse представляет собой структуру ответа на запрос обновлений Telegram
type UpdateResponse struct {
	// Ok указывает, был ли запрос успешным
	Ok bool `json:"ok"`
	// Result содержит массив обновлений
	Result []Update `json:"response"`
}

// Update представляет собой структуру для обновления Telegram
type Update struct {
	// ID - уникальный идентификатор обновления
	ID int `json:"update_id"`
	// Message - содержание сообщения
	Message string `json:"message"`
}
