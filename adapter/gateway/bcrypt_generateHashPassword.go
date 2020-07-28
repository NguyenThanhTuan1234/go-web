package gateway

import "golang.org/x/crypto/bcrypt"

func (b *bcyptClient) GenerateHashPassword(password string) (string, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}
