package model


type Session struct {
	Id        int64  `json:"id" datastore:"-"`
	Hash       string `json:"hash"`
	UserId    int64  `json:"userId"`
	Updated   int64  `json:"updated"`
	Generated int64	 `json:"generated"`
	Created   int64  `json:"created"`
	Expired   int64  `json:"expired"`
}

