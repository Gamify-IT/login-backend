package hash

// compile-time check that Mock implements the Hasher interface
var _ Hasher = &Mock{}

// Mock implements the [Hasher] interface. It MUST ONLY be used for testing.
//
// It returns the FixedReturnValue when [Mock.GenerateFromPassword] is called.
// It returns the specified FixedError when [Mock.CompareHashAndPassword] or [Mock.GenerateFromPassword] is called.
type Mock struct {
	FixedReturnValue []byte
	FixedError       error
}

// CompareHashAndPassword implements [Hasher.CompareHashAndPassword]
func (b *Mock) CompareHashAndPassword(_, _ []byte) error {
	return b.FixedError
}

// GenerateFromPassword implements [Hasher.GenerateFromPassword]
func (b *Mock) GenerateFromPassword(_ []byte, _ int) ([]byte, error) {
	return b.FixedReturnValue, b.FixedError
}
