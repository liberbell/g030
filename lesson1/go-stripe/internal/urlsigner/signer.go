package urlsigner

type Signer struct {
	Secret []byte
}

func (s *Signer) GenerateTokenFromString(data string) string {
	var urlToSign string
}

func (s *Signer) VeryfyToken(token string) bool {

}

func (s *Signer) Expired(token string, minutesUntilExpire int) bool {

}
