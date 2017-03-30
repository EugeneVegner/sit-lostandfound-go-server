package models

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"time"
)

type Token struct {
	Id      int64 `json:"id" datastore:"-"`
	Key     string `json:"key"`
	UserId  int64 `json:"userId"`
	Created int64 `json:"created"`
	Expired int64 `json:"expired"`
}

const (
	TokenLength   int           = 32
	TokenDuration time.Duration = 20 * time.Minute
)

func (token *Token) Generate() string {
	b := make([]byte, TokenLength)
	rand.Read(b)

	currentTime := time.Now().UTC()
	token.Created = currentTime.Unix()
	token.Expired = currentTime.UTC().Add(TokenDuration).Unix()
	token.Key = base64.StdEncoding.EncodeToString(b)

	return token.Key
}

func (token *Token) GenerateUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40

	token.Key = fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])

	currentTime := time.Now().UTC()
	token.Created = currentTime.Unix()
	token.Expired = currentTime.Add(1).Unix()
	return token.Key, nil
}

func (token *Token) IsExpired() bool {
	return token.Expired <= time.Now().UTC().Unix()
}
