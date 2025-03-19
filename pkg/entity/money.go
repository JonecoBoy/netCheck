// pkg/entity/money.go
package entity

import (
	"fmt"
	"math"
)

type Money struct {
	Value    int
	Currency Currency
}

func NewMoney(value int, currency Currency) *Money {
	return &Money{Value: value, Currency: currency}
}

func (m Money) ToFloat64() float64 {
	return float64(m.Value) / 100
}

func Float64ToMoney(f float64, currency Currency) Money {
	roundedValue := int(math.Round(f * 100))
	return Money{Value: roundedValue, Currency: currency}
}

func (m Money) String(symbol bool) string {
	if symbol {
		return fmt.Sprintf("%s %.2f", getCurrencySymbol(m.Currency), m.ToFloat64())
	} else {
		return fmt.Sprintf("%.2f %s", m.ToFloat64(), m.Currency)
	}
}
