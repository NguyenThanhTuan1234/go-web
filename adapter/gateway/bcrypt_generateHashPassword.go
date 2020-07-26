package gateway

import "golang.org/x/crypto/bcrypt"

func (b *bcyptClient) GenerateHashPassword(password string) ([]byte, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}
	return bs, nil
}
