package entity

import (
	pkgEntity "github.com/jonecoboy/netCheck/pkg/entity"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDomain(t *testing.T) {
	user := &User{
		ID: pkgEntity.NewId(),
	}
	domain := Domain{
		ID:        pkgEntity.NewId(),
		Domain:    "example.com",
		User:      user,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	assert.NotNil(t, domain.ID)
	assert.Equal(t, "example.com", domain.Domain)
	assert.Equal(t, user, domain.User)
}
