package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDomainStatus_ValidateStatus(t *testing.T) {
	tests := []struct {
		statusEnum int
		expectErr  bool
	}{
		{Domain_status_enum_unknown, false},
		{Domain_status_enum_online, false},
		{Domain_status_enum_offline, false},
		{999, true}, // Invalid status
	}

	for _, test := range tests {
		domainStatus := &DomainStatus{
			DomainStatusEnum: test.statusEnum,
		}
		err := domainStatus.ValidateStatus()
		if test.expectErr {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}
	}
}

func TestDomainStatus_Timestamps(t *testing.T) {
	now := time.Now()
	domainStatus := &DomainStatus{
		CreatedAt: now,
		UpdatedAt: now,
	}

	assert.Equal(t, now, domainStatus.CreatedAt)
	assert.Equal(t, now, domainStatus.UpdatedAt)
}
