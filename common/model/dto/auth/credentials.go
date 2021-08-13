package auth

type StandardCredentials struct {
	ID       uint64      `json:"id"`
	Metadata interface{} `json:"metadata"`
}
