package api

import (
	//"gopkg.in/gin-gonic/gin.v1"
	model "src/server/models"
	"github.com/asaskevich/govalidator"
	"appengine"
	"appengine/datastore"
	log "src/server/logger"
)

type AuthOutput struct {
	UserId int64 `json:"userId, required" binding:"required"`
	Token string `json:"token, required" binding:"required"`
	Expired int64 `json:"expired, required" binding:"required"`
}

func GetAndUpdateSessionIfNeeded(
	db appengine.Context,
	userKey *datastore.Key,
	deviceId string,
	deviceToken string,
	platform string) (*datastore.Key, *model.Session, error)  {

	log.Func(GetAndUpdateSessionIfNeeded)

	log.Debug("Try to get session ", userKey.Encode() )
	_, session, err1 := model.GetSessionByUserKeyAndDeviceId(db, userKey, deviceId)
	if err1 != nil {
		log.Debug("err1: ",err1.Error())
		return nil, nil, err1
	}
	log.Debug("Session not found")

	if session == nil {
		log.Debug("session nil: ","Session not found. Try to create new one")
		session = new(model.Session)
		session.UserId = userKey.IntID()
		session.DeviceId = deviceId
		session.Platform = platform
	}
	session.GenerateToken()
	session.DeviceToken = deviceToken

	log.Debug("Validate Session's fields...")
	if _, err2 := govalidator.ValidateStruct(session); err2 != nil {
		log.Debug("err2: ",err2.Error())
		return nil, nil, err2
	}

	log.Debug("Seve updated/new Session...")
	key, err3 := session.Save(db, userKey)
	if err3 != nil {
		log.Debug("err3: ",err3.Error())
		return nil, nil, err3
	}

	return key, session, nil
}