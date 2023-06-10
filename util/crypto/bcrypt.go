package crypto

import "golang.org/x/crypto/bcrypt"

const (
	cost = 8
)

type BcryptUtil struct{}

func (*BcryptUtil) EncryptPassword(pass string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	return string(bytes), err
}

func (*BcryptUtil) ValidatePassword(pass string, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err == nil, err
}
