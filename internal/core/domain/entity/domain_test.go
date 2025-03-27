package entity

import (
	"testing"

	"github.com/jonecoboy/netCheck/pkg/entity"
	"github.com/stretchr/testify/assert"
)

func TestDomain(t *testing.T) {
	user := &User{
		ID: entity.NewId(),
	}

	domain := NewDomain("example.com", user, nil)

	assert.NotNil(t, domain.ID)
	assert.Equal(t, "example.com", domain.Domain)
	assert.Equal(t, user, domain.User)
	assert.NotZero(t, domain.CreatedAt)
	assert.NotZero(t, domain.UpdatedAt)
	assert.Nil(t, domain.Records)
}
