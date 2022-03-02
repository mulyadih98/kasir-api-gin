package helper

import "golang.org/x/crypto/bcrypt"

type passwordHash struct{}

type PasswordHash interface {
	Hash(string) (string, error)
	Compare(string, string) error
}

func NewPasswordHash() PasswordHash {
	return passwordHash{}
}

func (hash passwordHash) Hash(password string) (string, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(pass), err
}

func (hash passwordHash) Compare(hashedPassword, password string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err
}
