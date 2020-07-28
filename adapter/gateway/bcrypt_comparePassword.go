package gateway

import "golang.org/x/crypto/bcrypt"

func (b *bcyptClient) ComparePassword(passwordDB []byte, password string) error {
	err := bcrypt.CompareHashAndPassword(passwordDB, []byte(password))
	if err != nil {
		return err
	}
	return nil
}
