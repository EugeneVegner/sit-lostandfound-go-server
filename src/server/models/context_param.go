package model

import (
	"gopkg.in/gin-gonic/gin.v1"
)

type ContextParam struct {
	Id            int64  `json:"id"`
	UserId        int64  `json:"fullName"`
	ClientDeviceToken string `json:"deviceToken"`
	ClientPlatform string `json:"platform"`
}

type handlerContextFunc func(*gin.Context)
//func (f handlerContextFunc) ServeSessionHTTP(route *Route) {
//	f(route)
//}

func (f handlerContextFunc) ConvertToParam() (*ContextParam, error) {

	err := func() {


		return nil
	}




	return err
}

//func ConvertToParam(ctx *gin.Context) (*ContextParam, error) {
//
//
//
//
//
//
//	return nil
//}