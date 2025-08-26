package models

import "time"

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
	token := 
}