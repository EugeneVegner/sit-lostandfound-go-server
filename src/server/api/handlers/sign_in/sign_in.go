package sign_in_handler

import (
	"gopkg.in/gin-gonic/gin.v1"
	"src/server/response"
	"appengine"
	e "src/server/errors"
	"src/server/models"
	"src/server/utils"
	log "src/server/logger"
	"src/server/api"
)

type input struct {
	Email    string `valid:"email,required" binding:"required"`
	Password string `valid:"length(8|24),required" binding:"required"`
	DeviceId string `valid:"required" binding:"required"`
	DeviceToken string
}

func POST(ctx *gin.Context) {
	log.Func(POST)
	var input input
	errors := utils.EncodeBody(ctx, &input)
	if errors != nil {
		response.Failed(ctx, errors, "The body can't be encoded")
		return
	}

	db := appengine.NewContext(ctx.Request)
	userKey, user, err1 := model.GetUserBy(db, "Email=", input.Email)
	if err1 != nil {
		log.Debug("err1: ",err1.Error())
		errors = append(errors, e.New("user", e.ServerErrorEmailNotFound, err1.Error()))
		response.Failed(ctx, errors, "User with email not found")
		return
	}

	_, session, err2 := api.GetAndUpdateSessionIfNeeded(db, userKey, input.DeviceId, input.DeviceToken)
	if err2 != nil {
		log.Debug("err2: ",err2.Error())
		errors = append(errors, e.New("user", e.ServerErrorSessionNotFound, err2.Error()))
		response.Failed(ctx, errors, "Session error")
		return
	}

	var output api.AuthOutput
	output.UserId = user.Id
	output.Token = session.Token
	output.Expired = session.Expired
	response.Success(ctx, output)
}