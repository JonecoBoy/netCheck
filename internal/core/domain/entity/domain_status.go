package entity

import (
	"errors"
	"time"
)

type DomainStatus struct {
	Domain           *Domain
	Task             *Task
	Records          []*Record
	DomainStatusEnum DomainStatusEnum `json:"domain_status_enum" bson:"domain_status_enum"`
	CreatedAt        time.Time        `bson:"created_at" json:"created_at"`
	UpdatedAt        time.Time        `bson:"updated_at" json:"updated_at"`
}

type DomainStatusEnum int

const (
	_ = iota
	DomainStatusEnumUnknown
	DomainStatusEnumOnline
	DomainStatusEnumOffline
)

func (e DomainStatusEnum) IsValid() bool {
	switch e {
	case DomainStatusEnumUnknown, DomainStatusEnumOnline, DomainStatusEnumOffline:
		return true
	default:
		return false
	}
}

func NewDomainStatus(
	domain *Domain,
	task *Task,
	records []*Record,
	statusEnum DomainStatusEnum,
) *DomainStatus {
	now := time.Now()

	return &DomainStatus{
		Domain:           domain,
		Task:             task,
		Records:          records,
		DomainStatusEnum: statusEnum,
		CreatedAt:        now,
		UpdatedAt:        now,
	}
}

func (d *DomainStatus) ValidateStatus() error {
	switch d.DomainStatusEnum {
	case DomainStatusEnumUnknown, DomainStatusEnumOnline, DomainStatusEnumOffline:
		return nil
	default:
		return errors.New("invalid domain status")
	}
}
