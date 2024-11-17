package errors

import "fmt"

// Error представляет собой структуру ошибки с кодом и сообщением.
type Error struct {
	Code    int    // Код ошибки (например, HTTP статус)
	Message string // Описание ошибки
}

func (e *Error) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

// Новый способ создания ошибки с кодом и сообщением
func New(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

func GetErroField(err error) (int, string) {
	if errWrap, ok := err.(*Error); ok {
		return errWrap.Code, errWrap.Message
	}
	return 500, "Internal server error"
}
