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
	Type    DomainTypeEnum
	Name    string
	Content string
}

type DomainRole struct {
	Domain *Domain
	User   *User
	Roles  int
}

type DomainTypeEnum int

const (
	_ DomainTypeEnum = iota
	DomainTypeEnumA
	DomainTypeEnumAAA
	DomainTypeEnumCNAME
	DomainTypeEnumTXT
	DomainTypeEnumMX
)

func (e DomainTypeEnum) IsValid() bool {
	switch e {
	case DomainTypeEnumA, DomainTypeEnumAAA, DomainTypeEnumCNAME, DomainTypeEnumTXT, DomainTypeEnumMX:
		return true
	default:
		return false
	}
}

func NewDomain(domain string, user *User, records []*Record) *Domain {
	now := time.Now()

	d := &Domain{
		ID:        pkgEntity.NewId(),
		Domain:    domain,
		User:      user,
		Records:   records,
		CreatedAt: now,
		UpdatedAt: now,
	}

	// Optionally, set domain reference in each record
	for _, r := range records {
		r.Domain = d
	}

	return d
}
