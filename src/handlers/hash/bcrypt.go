package hash

import "golang.org/x/crypto/bcrypt"

// compile-time check that Bcrypt implements the Hasher interface
var _ Hasher = &Bcrypt{}

type Bcrypt struct{}

func (b *Bcrypt) CompareHashAndPassword(hashedPassword, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}

func (b *Bcrypt) GenerateFromPassword(password []byte, cost int) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword(password, cost)
	return hash, err
}
