package entity

import (
	pkgEntity "github.com/jonecoboy/netCheck/pkg/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewOrderItem(t *testing.T) {
	itemID := pkgEntity.NewId()
	itemType := ItemTypeEnumProduct
	name := "Test Item"
	price := pkgEntity.NewMoney(100, "USD")
	quantity := 2

	item, err := NewOrderItem(itemID, itemType, name, price, quantity)
	assert.Nil(t, err)
	assert.Equal(t, itemID, item.ItemID)
	assert.Equal(t, itemType, item.ItemType)
	assert.Equal(t, name, item.Name)
	assert.Equal(t, price, item.Price)
	assert.Equal(t, quantity, item.Quantity)
}

func TestInvalidOrderItem(t *testing.T) {
	itemID := pkgEntity.NewId()
	itemType := ItemTypeEnum(999) // Invalid type
	name := "Test Item"
	price := pkgEntity.NewMoney(100, "USD")
	quantity := 2

	_, err := NewOrderItem(itemID, itemType, name, price, quantity)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid ItemTypeEnum", err.Error())
}
