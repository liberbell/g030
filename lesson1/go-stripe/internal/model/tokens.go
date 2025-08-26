package models

const (
	ScopeAuthentication = "authentication"
)

type Token struct {
	PlaneText string `json: "token"`
	UserID    int64  `json: "-"`
}
