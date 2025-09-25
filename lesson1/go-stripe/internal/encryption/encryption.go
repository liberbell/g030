package encryption

import "crypto/aes"

type Encryption struct {
	Key []byte
}

func (e *Encryption) Encrypt(text string) (string, error) {
	plaintext := []byte(text)
	block, err := aes.NewCipher(e.Key)
	if err != nil {
		return "", err
	}
}

func (e *Encryption) Decrypt(cryptoText string) (string, error) {

}
