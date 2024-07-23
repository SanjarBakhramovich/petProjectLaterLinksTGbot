package main

import (
	"flag"
	"log"
)

func main() {
	t := mustToken()
	// Получение токена из флагов командной строки
	// token = flags.Get(token)

	// Создание нового экземпляра fetcher для получения данных
	// fetcher = fetcher.New()

	// Создание нового экземпляра processor для обработки полученных данных
	// processor = processor.New()

	// Запуск consumer, который использует fetcher для получения данных и processor для их обработки
	// consumer.Start(fetcher, processor)

}

// mustToken возвращает токен доступа к Telegram боту, который передается через флаг командной строки "token-is-bot-token".
func mustToken() string {
	token := flag.String("token-is-bot-token", "", "token for access to telegram bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}
	return *token
}
