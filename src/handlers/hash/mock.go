package hash

// compile-time check that Mock implements the Hasher interface
var _ Hasher = &Mock{}

type Mock struct {
	FixedReturnValue []byte
	FixedError       error
}

func (b *Mock) CompareHashAndPassword(_, _ []byte) error {
	return b.FixedError
}

func (b *Mock) GenerateFromPassword(_ []byte, _ int) ([]byte, error) {
	return b.FixedReturnValue, b.FixedError
}
