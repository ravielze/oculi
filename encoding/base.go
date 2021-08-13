package encoding

type (
	Encoding interface {
		Marshal(val interface{}) ([]byte, error)
		Unmarshal(data []byte, val interface{}) error
	}
)
