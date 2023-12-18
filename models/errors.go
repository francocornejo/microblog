package models

type ErrorMessage struct {
	Type   string  `json:"type"`
	Title  string  `json:"title"`
	Detail string  `json:"detail"`
	Status int     `json:"status"`
	Errors []Error `json:"errors"`
}

type Error struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ErrorResponse(Title string, Detail string, Status int, ListaDeErrores []Error) ErrorMessage {
	return ErrorMessage{Type: "about:blank", Title: Title, Detail: Detail, Status: Status, Errors: ListaDeErrores}
}
