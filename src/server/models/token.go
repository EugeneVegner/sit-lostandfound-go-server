package model

import (
	"appengine"
	"appengine/datastore"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"log"
	"time"
)

//type Token struct {
//	Id        int64  `json:"id" datastore:"-"`
//	Hash      string `json:"key" datastore:"hash"`
//	UserId    int64  `json:"userId" datastore:"user_id, requared"`
//	Updated   int64  `json:"updated" datastore:"updated"`
//	Generated int64  `json:"generated" datastore:"generated"`
//	Created   int64  `json:"created" datastore:"created"`
//	Expired   int64  `json:"expired" datastore:"expired"`
//}

//type Token struct {
//	Id      int64  `json:"id" datastore:"-"`
//	Key     string `json:"key" datastore:"key,"`
//	UserId  int64  `json:"userId" datastore:"userId, requared"`
//	Updated int64  `json:"updated" datastore:","`
//	Generated int64  `json:"generated" datastore:","`
//	Created int64  `json:"created" datastore:","`
//	Expired int64  `json:"expired" datastore:","`
//}

type Token struct {
	Id        int64  `json:"id" datastore:"-"`
	Hash       string `json:"hash"`
	UserId    int64  `json:"userId"`
	Updated   int64  `json:"updated"`
	Generated int64	 `json:"generated"`
	Created   int64  `json:"created"`
	Expired   int64  `json:"expired"`
}

const (
	TokenLength   int           = 32
	TokenDuration time.Duration = 20 * time.Minute
)

func (token *Token) key(ctx appengine.Context, parentKey *datastore.Key) *datastore.Key {
	log.Println("key: ")
	if token.Id == 0 {
		t := time.Now().UTC().Unix()
		token.Updated = t
		token.Created = t
		return datastore.NewIncompleteKey(ctx, "Token", parentKey)
	}
	return datastore.NewKey(ctx, "Token", "", token.Id, parentKey)
}

func SaveToken(ctx appengine.Context, token *Token, parentKey *datastore.Key) (*datastore.Key, error) {
	log.Println("SaveToken: ")
	token.Updated = time.Now().UTC().Unix()
	k, err := datastore.Put(ctx, token.key(ctx, parentKey), token)
	if err != nil {
		return nil, err
	}
	token.Id = k.IntID()
	return k, nil
}

func (token *Token) Generate() string {
	log.Println("Generate: ")
	b := make([]byte, TokenLength)
	rand.Read(b)

	currentTime := time.Now().UTC()
	token.Generated = currentTime.Unix()
	token.Expired = currentTime.Add(TokenDuration).Unix()
	token.Hash = base64.StdEncoding.EncodeToString(b)

	return token.Hash
}

func (token *Token) GenerateUUID() (string, error) {
	log.Println("GenerateUUID: ")
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40

	token.Hash = fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])

	currentTime := time.Now().UTC()
	token.Generated = currentTime.Unix()
	token.Created = currentTime.Unix()
	token.Expired = currentTime.Add(1).Unix()
	return token.Hash, nil
}

func (token *Token) IsExpired() bool {
	log.Println("IsExpired: ")
	return token.Expired <= time.Now().UTC().Unix()
}

// Filtering

func GetTokenByUserId(ctx appengine.Context, userId interface{}) (*datastore.Key, *Token, error) {
	log.Println("GetTokenByUserId: ")
	return GetTokenBy(ctx, "UserId=", userId)
}

func GetTokenByKey(ctx appengine.Context, key interface{}) (*datastore.Key, *Token, error) {
	log.Println("GetTokenByKey: ")
	return GetTokenBy(ctx, "Hash=", key)
}

func GetTokenBy(ctx appengine.Context, filter string, value interface{}) (*datastore.Key, *Token, error) {
	log.Println("GetTokenBy: ")
	var tokens []Token
	q := datastore.NewQuery("Token").Filter(filter, value).Order("Created").Limit(1)
	ks, err := q.GetAll(ctx, &tokens)
	if err != nil {
		return nil, nil, err
	}
	if len(tokens) == 0 {
		return nil, nil, errors.New("Token not found")
	}
	for i := 0; i < len(tokens); i++ {
		tokens[i].Id = ks[i].IntID()
	}
	return ks[0], &tokens[0], nil
}
