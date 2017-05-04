package model

import (
	"time"
	"appengine"
	"appengine/datastore"
	log "src/server/logger"
	"crypto/rand"
	"encoding/base64"
	e "src/server/errors"
)

type Session struct {
	Id          int64  `json:"id" datastore:"-"`
	Token       string `json:"token" valid:"required"`
	DeviceId    string `json:"deviceId" valid:"deviceId, required"`
	DeviceToken string `json:"deviceToken"`
	Platform    string `json:"platform" valid:"required"`
	UserId      int64  `json:"userId" valid:"required"`
	Updated     int64  `json:"updated"`
	Created     int64  `json:"created"`
	Expired     int64  `json:"expired" valid:"required"`
}

func (session *Session) key(db appengine.Context, parentKey *datastore.Key) *datastore.Key {
	t := time.Now().UTC().Unix()
	session.Updated = t

	if session.Id == 0 {
		session.Created = t
		log.Debug("Session: NewIncompleteKey")
		return datastore.NewIncompleteKey(db, "Session", parentKey)
	}
	log.Debug("Session: NewKey")
	return datastore.NewKey(db, "Session", "", session.Id, parentKey)
}

func (session *Session) Save(db appengine.Context, parentKey *datastore.Key) (*datastore.Key, error) {
	log.Debug("Session: Save")
	k, err := datastore.Put(db, session.key(db, parentKey), session)
	if err != nil {
		return nil, err
	}
	session.Id = k.IntID()
	return k, nil
}

// Filtering

func GetSessionBy(ctx appengine.Context, filter string, value interface{}) (*datastore.Key, *Session, error) {
	log.Func(GetSessionBy)
	var sessions []Session
	q := datastore.NewQuery("Session").
		Filter(filter, value).
		Order("Created").
		Limit(1)
	ks, err := q.GetAll(ctx, &sessions)
	if err != nil {
		return nil, nil, err
	}
	if len(sessions) == 0 {
		return nil, nil, e.ErrorNotFound
	}
	for i := 0; i < len(sessions); i++ {
		sessions[i].Id = ks[i].IntID()
	}
	return ks[0], &sessions[0], nil
}

func GetSessionByUserKeyAndDeviceId(ctx appengine.Context, userKey *datastore.Key, deviceId interface{}) (*datastore.Key, *Session, error) {
	log.Func(GetSessionByUserKeyAndDeviceId)
	var sessions []Session
	q := datastore.NewQuery("Session").
		Ancestor(userKey).
		Filter("UserId=", userKey.IntID()).
		Filter("DeviceId=", deviceId).
		Order("Created").
		Limit(1)
	keys, err := q.GetAll(ctx, &sessions)
	if err != nil {
		return nil, nil, err
	}
	for i := 0; i < len(sessions); i++ {
		sessions[i].Id = keys[i].IntID()
	}
	if len(sessions) > 0 && len(keys) > 0 {
		return keys[0], &sessions[0], nil
	}
	log.Debug("No Sessions, No Keys ")
	return nil, nil, nil
}

// Methods

func (session *Session) GenerateToken() string {
	log.Debug("Generate token")
	b := make([]byte, TokenLength)
	rand.Read(b)

	currentTime := time.Now().UTC()
	session.Updated = currentTime.Unix()
	session.Expired = currentTime.Add(TokenDuration).Unix()
	session.Token = base64.StdEncoding.EncodeToString(b)
	log.Debug("Token is: ", session.Token)
	return session.Token
}

func (session *Session) IsExpired() bool {
	log.Debug("IsExpired")
	return session.Expired <= time.Now().UTC().Unix()
}
