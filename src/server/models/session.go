package model

import (
//"appengine/datastore"
)

type Session struct {
	Token   string `json:"validate_token"`
	Created int64  `json:"created"`
	Expired int64  `json:"expired"`
	UserId  string `json:"userId"`
}

//const (
//	//TokenLength   int           = 32
//	//TokenDuration time.Duration = 20 * time.Minute
//)
//
