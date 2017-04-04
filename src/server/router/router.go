package router

import (
	"gopkg.in/gin-gonic/gin.v1"
)

type Route struct {
	Name    string
	Method  string
	Pattern string
	Error error
	Context *gin.Context
}

type HandlerSessionFunc func(*Route)
func (f HandlerSessionFunc) ServeSessionHTTP(route *Route) {
	f(route)
}

//type HandlerGinFunc func(*gin.Context)
func Session(h HandlerSessionFunc) gin.HandlerFunc {
	return func(c *gin.Context) {

		var rout Route
		rout.Context = c
		rout.Error = nil


		h.ServeSessionHTTP(&rout)
	}
}
