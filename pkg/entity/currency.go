package entity

type Currency string

const (
	BRL Currency = "BRL"
	USD Currency = "USD"
	EUR Currency = "EUR"
	GBP Currency = "GBP"
	// ....
)

func getCurrencySymbol(currency Currency) string {
	switch currency {
	case BRL:
		return "R$"
	case USD:
		return "$"
	case EUR:
		return "€"
	case GBP:
		return "£"
		// ...
	}
	return ""
}
