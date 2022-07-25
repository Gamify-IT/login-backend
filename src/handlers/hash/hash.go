// Package hash implements password hashing APIs.
package hash

// The Hasher interface specifies the methods required for password hashing and validation.
type Hasher interface {
	// CompareHashAndPassword verifies the password agains the hashPassword
	CompareHashAndPassword(hashedPassword, password []byte) error
	// GenerateFromPassword generates a new hash from the provided password with the specified computational cost.
	GenerateFromPassword(password []byte, cost int) ([]byte, error)
}
