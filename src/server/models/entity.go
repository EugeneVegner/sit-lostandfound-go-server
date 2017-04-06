package model

import (
	"appengine"
	"appengine/datastore"
	log "src/server/logger"
	"errors"
)


// TODO: Need to make general function: GetEntityBy

func GetEntityBy(ctx appengine.Context, entityName string, filter string, value interface{}) (*datastore.Key, *Session, error) {
	log.Func(GetSessionBy)
	var sessions []Session
	q := datastore.NewQuery(entityName).Filter(filter, value).Order("Created").Limit(1)
	ks, err := q.GetAll(ctx, &sessions)
	if err != nil {
		return nil, nil, err
	}
	if len(sessions) == 0 {
		return nil, nil, errors.New("Token not found")
	}
	for i := 0; i < len(sessions); i++ {
		sessions[i].Id = ks[i].IntID()
	}
	return ks[0], &sessions[0], nil
}

//func SaveEntity(ctx appengine.Context, entity interface{}, parentKey *datastore.Key) (*datastore.Key, error) {
//	log.Func(SaveSession)
//	entity.Updated = time.Now().UTC().Unix()
//	k, err := datastore.Put(ctx, session.key(ctx, parentKey), session)
//	if err != nil {
//		return nil, err
//	}
//	entity.Id = k.IntID()
//	return k, nil
//}
