package entity

import (
	pkgEntity "github.com/jonecoboy/netCheck/pkg/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrder_CalculateSubtotal(t *testing.T) {
	item := OrderItem{
		ItemID:   pkgEntity.NewId(),
		ItemType: ItemTypeEnumProduct,
		Name:     "Item",
		Price:    pkgEntity.NewMoney(100, "USD"),
		Quantity: 1,
	}

	order := Order{
		Items: []*OrderItem{&item},
		Totals: &Totals{
			//Shipping:  *pkgEntity.NewMoney(0, "USD"),
			Discounts: []Discount{},
		},
	}
	order.CalculateSubtotal()
	assert.Equal(t, 100, order.Totals.Subtotal.Value)
}

func TestOrder_CalculateTotal(t *testing.T) {
	item := OrderItem{
		ItemID:   pkgEntity.NewId(),
		ItemType: ItemTypeEnumProduct,
		Name:     "Item",
		Price:    pkgEntity.NewMoney(100, "USD"),
		Quantity: 1,
	}
	discount := Discount{
		ID:          pkgEntity.NewId(),
		Description: "10% off",
		Type:        DiscountTypeEnumPercent,
		Percent:     new(int),
	}
	*discount.Percent = 10
	order := Order{
		Items: []*OrderItem{&item},
		Totals: &Totals{
			Shipping:  *pkgEntity.NewMoney(10, "USD"),
			Discounts: []Discount{discount},
		},
	}
	order.CalculateSubtotal()
	order.CalculateTotal()
	assert.Equal(t, 100, order.Totals.Subtotal.Value)
	assert.Equal(t, 100-10+10, order.Totals.Total.Value)
}
