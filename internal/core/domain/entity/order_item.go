package entity

import (
	"errors"
	"github.com/jonecoboy/netCheck/pkg/entity"
)

type OrderItem struct {
	ItemID   entity.ID    `json:"product_id" bson:"product_id"`
	ItemType ItemTypeEnum `json:"product_type" bson:"product_type"`
	Name     string       `json:"name" bson:"name"`
	Price    entity.Money `json:"price" bson:"price"`
	Quantity int          `json:"quantity" bson:"quantity"`
}

const (
	Product ItemTypeEnum = iota
	Subscription
)

func (e ItemTypeEnum) IsValid() bool {
	switch e {
	case Product, Subscription:
		return true
	default:
		return false
	}
}

func NewOrderItem(itemID entity.ID, itemType ItemTypeEnum, name string, price entity.Money, quantity int) (OrderItem, error) {
	if !itemType.IsValid() {
		return OrderItem{}, errors.New("invalid ItemTypeEnum")
	}
	return OrderItem{
		ItemID:   itemID,
		ItemType: itemType,
		Name:     name,
		Price:    price,
		Quantity: quantity,
	}, nil
}
