package entity

import (
	pkgEntity "github.com/jonecoboy/netCheck/pkg/entity"
	"time"
)

type Domain struct {
	ID        pkgEntity.ID
	Domain    string
	User      *User
	Records   []*Record
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
	DeletedAt time.Time `bson:"deleted_at" json:"deleted_at"`
}

type Record struct {
	Domain  *Domain
	Type    string
	Name    string
	Content string
}

type DomainRole struct {
	Domain *Domain
	User   *User
	Roles  int
}

const (
	Domain_type_enum_A = iota
	Domain_type_enum_AAA
	Domain_type_enum_CNAME
	Domain_type_enum_TXT
	Domain_type_enum_MX
)
