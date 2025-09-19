package urlsigner

import (
	goalone "github.com/bwmarrin/go-alone"
)

type Signer struct {
	Secret []byte
}

func (s *Signer) GenerateTokenFromString(data string) string {
	var urlToSign string

	crypt := goalone.New(s.Secret, goalone.Timestamp)
}

func (s *Signer) VeryfyToken(token string) bool {

}

func (s *Signer) Expired(token string, minutesUntilExpire int) bool {

}
