package entity

import (
	pkgEntity "github.com/jonecoboy/netCheck/pkg/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPlan(t *testing.T) {
	name := "Basic Plan"
	description := "This is a basic plan"
	price := pkgEntity.NewMoney(1000, "USD")
	plan, err := NewPlan(name, description, *price)
	assert.Nil(t, err)
	assert.NotNil(t, plan)
	assert.Equal(t, name, plan.Name)
	assert.Equal(t, description, plan.Description)
	assert.Equal(t, price, plan.Price)
	assert.False(t, plan.CreatedAt.IsZero())
	assert.False(t, plan.UpdatedAt.IsZero())
}
