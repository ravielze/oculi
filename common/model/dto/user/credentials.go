package user

type CredentialsDTO struct {
	ID       uint64      `json:"id"`
	Metadata interface{} `json:"metadata"`
}
