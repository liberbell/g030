package urlsigner

type Signer struct {
	Secret []byte
}

func (s *Signer) GenerateTokenFromString(data string) string {

}

func (s *Signer) VeryfyToken(token string) bool {

}
