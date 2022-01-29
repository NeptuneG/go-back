package db

import "golang.org/x/crypto/bcrypt"

func (u User) Authenticate(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password))
}
