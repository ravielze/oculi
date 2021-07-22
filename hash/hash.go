package hash

type (
	Hash interface {
		Hash(raw string) (string, error)
		Verify(raw string, hashed string) error
	}
)
