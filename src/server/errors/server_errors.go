package errors

import (
	"errors"
)

const (
	ServerErrorUnknown          int = 0
	ServerError1                int = 1
	ServerError2                int = 2
	ServerError3                int = 3
	ServerError4                int = 4
	ServerErrorClientHeader     int = 5
	ServerError6                int = 6
	ServerError7                int = 7
	ServerErrorInvalidStructure int = 8
	ServerErrorInvalidBody      int = 9
	ServerError10               int = 10
	ServerErro11                int = 11
	ServerError12               int = 12
	ServerError13               int = 13
	ServerError14               int = 14
	ServerError15               int = 15
	ServerError16               int = 16
	ServerError17               int = 17
	ServerError18               int = 18
	ServerError19               int = 19
	ServerError20               int = 20
	ServerError21               int = 21
	ServerError22               int = 22
	ServerError23               int = 23
	ServerError24               int = 24
	ServerError25               int = 25
	ServerError26               int = 26
	ServerErrorNewEntity        int = 27
	ServerErrorSessionNotFound  int = 28
	ServerError29               int = 29
	ServerErrorUserNotFound     int = 30
	ServerError31               int = 31
	ServerErrorEmailNotFound    int = 32
	ServerErrorEmailExist       int = 33
)

type Error struct {
	Name    string `json:"name"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Unknown() Error {
	return Error{"unknown", ServerErrorUnknown, "Unknown error"}
}

func New(tag string, code int, message string) Error {
	return Error{tag, code, message}
}

// Error types
var (
	ErrorNotFound = errors.New("Not found")
	ErrorEOF      = errors.New("EOF")
)
