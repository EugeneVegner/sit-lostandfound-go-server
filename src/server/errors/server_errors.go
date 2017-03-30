package errors

type Error struct {
	Tag string `json:"tag"`
	Code int    `json:"code"`
	Message string `json:"message"`
}

func Unknown(code int) Error {
	return Error{"unknown", code,"Unknown error"}
}

func New(tag string, code int, message string) Error {
	return Error{tag, code,message}
}

