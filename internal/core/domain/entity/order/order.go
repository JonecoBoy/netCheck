package entity

import (
	"github.com/jonecoboy/netCheck/pkg/entity"
	"time"
)

type ItemTypeEnum int
type DiscountType int

const (
	MoneyDiscount DiscountType = iota
	PercentDiscount
)

type Order struct {
	ID        entity.ID   `json:"id" bson:"id"`
	UserID    entity.ID   `json:"user_id" bson:"user_id"`
	Address   string      `json:"address" bson:"address"`
	Items     []OrderItem `json:"items" bson:"items"`
	Totals    Totals      `json:"totals" bson:"totals"`
	CreatedAt time.Time   `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time   `bson:"updated_at" json:"updated_at"`
	DeletedAt time.Time   `bson:"deleted_at" json:"deleted_at"`
}

type Discount struct {
	ID          entity.ID     `json:"id" bson:"id"`
	Description string        `json:"description" bson:"description"`
	Type        DiscountType  `json:"type" bson:"type"`
	Money       *entity.Money `json:"money,omitempty" bson:"money,omitempty"`
	Percent     *int          `json:"percent,omitempty" bson:"percent,omitempty"`
}

type Totals struct {
	Shipping  entity.Money `json:"shipping" bson:"shipping"`
	Subtotal  entity.Money `json:"subtotal" bson:"subtotal"`
	Discounts []Discount   `json:"discounts" bson:"discounts"`
	Total     entity.Money `json:"total" bson:"total"`
}

func (o *Order) CalculateSubtotal() {
	var subtotal int
	for _, item := range o.Items {
		subtotal += item.Price.Value * item.Quantity
	}
	o.Totals.Subtotal = entity.NewMoney(subtotal, o.Items[0].Price.Currency)
}

func (o *Order) CalculateTotal() {
	discountValue := 0
	for _, discount := range o.Totals.Discounts {
		switch discount.Type {
		case MoneyDiscount:
			if discount.Money != nil {
				discountValue += discount.Money.Value
			}
		case PercentDiscount:
			if discount.Percent != nil {
				discountValue += o.Totals.Subtotal.Value * *discount.Percent / 100
			}
		}
	}

	o.Totals.Total = entity.NewMoney(o.Totals.Subtotal.Value-discountValue+o.Totals.Shipping.Value, o.Totals.Subtotal.Currency)
}
