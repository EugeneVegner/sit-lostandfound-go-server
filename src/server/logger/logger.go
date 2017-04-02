package logger

import (
	"log"
	"reflect"
	"runtime"
	"src/server/constants"
)

func Debug(format string, values ...interface{}) {
	debugPrint("*[DEBUG] "+format+"%#v\n", values)
}

func DebugError(format string, values ...interface{}) {
	debugPrint("*[ERROR] "+format+"%v\n", values)
}

func Error(err error) {
	debugPrint("*[ERROR] %v\n", err)
}

func Func(f interface{}) {
	handlerName := nameOfFunction(f)
	debugPrint("[FUNCTION] %v\n", handlerName)
}

func debugPrint(format string, values ...interface{}) {
	if constants.DevelopmentMode {
		log.Printf("LOG "+format, values...)
	}
}
func nameOfFunction(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}
