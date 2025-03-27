package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDomainStatus_ValidateStatus(t *testing.T) {
	domain := NewDomain("example.com", nil, nil)

	tests := []struct {
		statusEnum int
		expectErr  bool
	}{
		{DomainStatusEnumUnknown, false},
		{DomainStatusEnumOnline, false},
		{DomainStatusEnumOffline, false},
		{999, true}, // Invalid status
	}

	for _, test := range tests {
		domainStatus := NewDomainStatus(domain, nil, nil, DomainStatusEnumOnline)
		err := domainStatus.ValidateStatus()
		if test.expectErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestDomainStatus_Timestamps(t *testing.T) {
	domain := NewDomain("example.com", nil, nil)
	domainStatus := NewDomainStatus(domain, nil, nil, DomainStatusEnumOnline)

	assert.WithinDuration(t, time.Now(), domainStatus.CreatedAt, time.Second)
	assert.WithinDuration(t, time.Now(), domainStatus.UpdatedAt, time.Second)
}
