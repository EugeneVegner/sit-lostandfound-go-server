package model

import (
	"time"
	"errors"
	"appengine"
	"appengine/datastore"
	log "src/server/logger"
)

type Session struct {
	Id      int64  `json:"id" datastore:"-"`
	Token   Token  `json:"token"`
	Device  Device `json:"device"`
	UserId  int64  `json:"userId"`
	Updated int64  `json:"updated"`
	Created int64  `json:"created"`
	Expired int64  `json:"expired"`
}

func (session *Session) key(ctx appengine.Context, parentKey *datastore.Key) *datastore.Key {
	//log.Println("key: ")
	if session.Id == 0 {
		t := time.Now().UTC().Unix()
		session.Updated = t
		session.Created = t
		return datastore.NewIncompleteKey(ctx, "Session", parentKey)
	}
	return datastore.NewKey(ctx, "Session", "", session.Id, parentKey)
}

func SaveSession(ctx appengine.Context, session *Session, parentKey *datastore.Key) (*datastore.Key, error) {
	log.Func(SaveSession)
	session.Updated = time.Now().UTC().Unix()
	k, err := datastore.Put(ctx, session.key(ctx, parentKey), session)
	if err != nil {
		return nil, err
	}
	session.Id = k.IntID()
	return k, nil
}

// Filtering

func GetSessionByUserId(ctx appengine.Context, userId interface{}) (*datastore.Key, *Session, error) {
	log.Func(GetSessionByUserId)
	return GetSessionBy(ctx, "UserId=", userId)
}

func GetSessionBy(ctx appengine.Context, filter string, value interface{}) (*datastore.Key, *Session, error) {
	log.Func(GetSessionBy)
	var sessions []Session
	q := datastore.NewQuery("Session").Filter(filter, value).Order("Created").Limit(1)
	ks, err := q.GetAll(ctx, &sessions)
	if err != nil {
		return nil, nil, err
	}
	if len(sessions) == 0 {
		return nil, nil, errors.New("Session not found")
	}
	for i := 0; i < len(sessions); i++ {
		sessions[i].Id = ks[i].IntID()
	}
	return ks[0], &sessions[0], nil
}
