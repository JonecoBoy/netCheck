package entity

import (
	pkgEntity "github.com/jonecoboy/netCheck/pkg/entity"
	"strings"
)

type Address struct {
	ID           pkgEntity.ID
	Street       string
	Number       string
	Complement   string
	Neighborhood string
	City         string
	State        string
	Zip          string
	Country      string
}

func NewAddress(
	street, number, complement, neighborhood, city, state, zip, country string,
) Address {
	return Address{
		ID:           pkgEntity.NewId(),
		Street:       street,
		Number:       number,
		Complement:   complement,
		Neighborhood: neighborhood,
		City:         city,
		State:        state,
		Zip:          zip,
		Country:      country,
	}
}

func (a Address) FullAddress() string {
	var sb strings.Builder

	sb.WriteString(a.Street)
	sb.WriteString(" ")
	sb.WriteString(a.Number)

	if a.Complement != "" {
		sb.WriteString(", ")
		sb.WriteString(a.Complement)
	}

	if a.Neighborhood != "" {
		sb.WriteString(" - ")
		sb.WriteString(a.Neighborhood)
	}

	sb.WriteString(", ")
	sb.WriteString(a.City)

	if a.State != "" {
		sb.WriteString(" - ")
		sb.WriteString(a.State)
	}

	if a.Zip != "" {
		sb.WriteString(" - ")
		sb.WriteString(strings.ToUpper(a.Zip))
	}

	if a.Country != "" {
		sb.WriteString(" - ")
		sb.WriteString(strings.ToUpper(a.Country))
	}

	return sb.String()
}
