package telegram

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
	"readAdviserBot/lib/e"
	"strconv"
)

type Client struct {
	// host - это базовый URL API Telegram
	host string

	// basePath - это путь к API, обычно "/bot<token>"
	basePath string

	// client - это HTTP-клиент для выполнения запросов
	client http.Client
}

const (
	// getUpdatesMethod - это константа, содержащая название метода API Telegram для получения обновлений.
	getUpdatesMethod = "getUpdates"

	sendMessageMethod = "sendMessage"
)

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

// Updates получает обновления от Telegram API
func (c *Client) Updates(offset int, limit int) ([]Update, error) {
	query := url.Values{}
	query.Add("offset", strconv.Itoa(offset))
	query.Add("limit", strconv.Itoa(limit))

	data, err := c.doRequest(getUpdatesMethod, query)
	if err != nil {
		return nil, err
	}

	var res UpdateResponse
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res.Result, nil
}

func (c *Client) SendMessage(chatID int, text string) error {
	query := url.Values{}
	query.Add("chat_id", strconv.Itoa(chatID))
	query.Add("text", text)

	_, err := c.doRequest(sendMessageMethod, query)
	if err != nil {
		return e.Wrap("can not send message", err)
	}

	return nil
}

// doRequest выполняет HTTP-запрос к Telegram API
func (c *Client) doRequest(method string, query url.Values) (data []byte, err error) {
	defer func() { err = e.WrapIfErr("can not do request", err) }()
	// Формируем URL для запроса
	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, method),
	}

	// Создаем новый HTTP-запрос
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	// Добавляем параметры запроса
	req.URL.RawQuery = query.Encode()

	// Выполняем запрос
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	// Читаем тело ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
