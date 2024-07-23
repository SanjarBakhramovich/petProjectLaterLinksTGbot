package telegram

import "net/http"

type Client struct {
	// host - это базовый URL API Telegram
	host string

	// basePath - это путь к API, обычно "/bot<token>"
	basePath string

	// client - это HTTP-клиент для выполнения запросов
	client http.Client
}

// Создает новый экземпляр клиента Telegram
// host: адрес API Telegram
// token: токен бота для аутентификации
func New(host string, token string) Client {
	return Client{
		// Устанавливает хост API
		host: host,
		// Создает базовый путь для API, используя токен
		basePath: newBasePath(token),
		// Инициализирует HTTP-клиент для выполнения запросов
		client: http.Client{},
	}
}

// newBasePath создает базовый путь для API Telegram
// token: токен бота для аутентификации
func newBasePath(token string) string {
	// Формирует и возвращает базовый путь, добавляя префикс "bot" к токену
	return "bot" + token
}

func (c *Client) Updated() {

}

func (c *Client) SendMessage() {

}
