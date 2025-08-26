package models

import (
	"crypto/rand"
	"time"
)

const (
	ScopeAuthentication = "authentication"
)

type Token struct {
	PlaneText string    `json: "token"`
	UserID    int64     `json: "-"`
	Hash      []byte    `json: "-"`
	Expiry    time.Time `json: "expiry"`
	Scope     string    `json: "scope"`
}

func GenerateToken(userID int, ttl time.Duration, scope string) (*Token, error) {
	token := &Token{
		UserID: int64(userID),
		Expiry: time.Now().Add(ttl),
		Scope:  scope,
	}

	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}
}
