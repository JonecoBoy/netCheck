package entity

import (
	domainError "github.com/jonecoboy/netCheck/internal/core/domain/error"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewSubscription(t *testing.T) {
	t.Run("should create a subscription for active user", func(t *testing.T) {
		user := &User{
			DeletedAt: time.Time{}, // not deleted
		}
		plan := &Plan{}

		subscription, err := NewSubscription(user, plan)

		assert.NoError(t, err)
		assert.NotNil(t, subscription)
		assert.Equal(t, *user, subscription.User)
		assert.Equal(t, *plan, subscription.Plan)
	})

	t.Run("should fail to create subscription for deleted user", func(t *testing.T) {
		user := &User{
			DeletedAt: time.Now(), // user marked as deleted
		}
		plan := &Plan{}

		subscription, err := NewSubscription(user, plan)

		assert.Nil(t, subscription)
		assert.ErrorIs(t, err, domainError.ErrUserDeleted)
	})
}
