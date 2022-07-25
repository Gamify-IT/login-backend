package hash

import "golang.org/x/crypto/bcrypt"

// compile-time check that Bcrypt implements the Hasher interface
var _ Hasher = &Bcrypt{}

// Bcrypt implements the Hasher interface.
//
// It uses [golang.org/x/crypto/bcrypt] as cryptographic backend.
type Bcrypt struct{}

// CompareHashAndPassword implements [Hasher.CompareHashAndPassword]
func (b *Bcrypt) CompareHashAndPassword(hashedPassword, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}

// GenerateFromPassword implements [Hasher.GenerateFromPassword]
func (b *Bcrypt) GenerateFromPassword(password []byte, cost int) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword(password, cost)
	return hash, err
}
