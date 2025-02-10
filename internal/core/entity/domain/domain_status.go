package domain

import (
	"errors"
	"github.com/JonecoBoy/netCheck/internal/core/entity/task"
	"time"
)

type DomainStatus struct {
	Domain           Domain
	Task             tasks.tasks
	Records          []Record
	DomainStatusEnum int       `json:"domain_status_enum" bson:"domain_status_enum"`
	CreatedAt        time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt        time.Time `bson:"updated_at" json:"updated_at"`
}

const (
	Domain_status_enum_unknown = iota
	Domain_status_enum_online
	Domain_status_enum_offline
)

func (d *DomainStatus) ValidateStatus() error {
	switch d.DomainStatusEnum {
	case Domain_status_enum_unknown, Domain_status_enum_online, Domain_status_enum_offline:
		return nil
	default:
		return errors.New("invalid domain status")
	}
}
