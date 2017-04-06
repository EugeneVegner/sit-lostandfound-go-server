package model

import (
	"gopkg.in/gin-gonic/gin.v1"
	c "src/server/constants"
	"time"
)

type Device struct {
	Id       string `json:"id"`
	Token    string `json:"token"`
	Platform string `json:"platform"`
	Updated  int64  `json:"updated"`
	Created  int64  `json:"created"`
}

func (device *Device) New(ctx *gin.Context) error {

	device.Id = ctx.Param(c.ParamKeyClientDeviceId)
	device.Token = ctx.Param(c.ParamKeyClientDeviceToken)
	device.Platform = ctx.Param(c.ParamKeyClientPlatform)
	currentTime := time.Now().UTC()
	device.Created = currentTime.Unix()
	device.Updated = currentTime.Unix()

	return nil
}