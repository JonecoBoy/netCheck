package domain

import (
	"github.com/JonecoBoy/netCheck/internal/core/domain/entity/user"
	"time"
)

type Domain struct {
	Domain    string
	User      user.User
	Records   []Record
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

type Record struct {
	Domain  Domain
	Type    string
	Name    string
	Content string
}

type DomainRole struct {
	Domain Domain
	User   user.User
	Roles  int
}

const (
	Domain_type_enum_A = iota
	Domain_type_enum_AAA
	Domain_type_enum_CNAME
	Domain_type_enum_TXT
	Domain_type_enum_MX
)
