package router

import (
	"github.com/gin-gonic/gin"
)

type Route struct {
	Name    string
	Method  string
	Pattern string
	Error error
	Context *gin.Context
}

func RoutFunc(c *gin.Context) *Route  {
	var rout Route
	rout.Context = c
	rout.Error = nil
	return &rout
}


type HandlerSessionFunc func(*Route)
func (f HandlerSessionFunc) ServeSessionHTTP(route *Route) {
	f(route)
}

type HandlerGinFunc func(*gin.Context)
func Session(h HandlerSessionFunc) HandlerGinFunc {
	return func(c *gin.Context) {

		var rout Route
		rout.Context = c
		rout.Error = nil


		h.ServeSessionHTTP(&rout)
	}
}
