package crypto

import "golang.org/x/crypto/bcrypt"

const (
	cost = 8
)

func EncryptPasswordWithBcrypt(pass string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	return string(bytes), err
}
