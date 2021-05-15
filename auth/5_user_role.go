package auth

import "strings"

type Role int16

const (
	ROLE_DEFAULT Role = iota + 1
	ROLE_ADMIN
	//insert new role here
)

// lowercase role name => id (iota+1)
var roles = map[string]Role{
	"admin":   1,
	"default": 2,
	//insert new role here
}

func (r Role) IsExist() bool {
	//change 2 to the last number
	return int16(r) >= 1 && int16(r) <= 2
}

func (r Role) String() string {
	//add role name here
	return [...]string{"Default", "Admin"}[r]
}

func GetRole(role string) Role {
	return roles[strings.ToLower(role)]
}

func (r Role) Equal(other Role) bool {
	return (int16(r) == int16(other))
}

func (r Role) IsRestricted() bool {
	return r.Equal(ROLE_ADMIN)
}
