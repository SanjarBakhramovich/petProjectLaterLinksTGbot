package e

import "fmt"

// Wrap оборачивает ошибку err с дополнительным сообщением msg
func Wrap(msg string, err error) error {
	return fmt.Errorf("%s:  %w", msg, err)
}

// WrapIfErr оборачивает ошибку err с дополнительным сообщением msg,
// если ошибка не равна nil. В противном случае возвращает nil.
func WrapIfErr(msg string, err error) error {
	// Проверяем, является ли ошибка nil
	if err == nil {
		// Если ошибка nil, возвращаем nil
		return nil
	}
	// Если ошибка не nil, оборачиваем её с помощью функции Wrap
	return Wrap(msg, err)
}
