package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddress(t *testing.T) {
	street := "123 Main St"
	number := "456"
	complement := "Apt 789"
	neighborhood := "Downtown"
	city := "Metropolis"
	state := "NY"
	zip := "12345"
	country := "USA"

	address := NewAddress(street, number, complement, neighborhood, city, state, zip, country)
	assert.NotNil(t, address.ID)
	assert.NotEmpty(t, address.ID)
	assert.Equal(t, street, address.Street)
	assert.Equal(t, number, address.Number)
	assert.Equal(t, complement, address.Complement)
	assert.Equal(t, neighborhood, address.Neighborhood)
	assert.Equal(t, city, address.City)
	assert.Equal(t, state, address.State)
	assert.Equal(t, zip, address.Zip)
	assert.Equal(t, country, address.Country)
	assert.Equal(t, "123 Main St 456, Apt 789 - Downtown, Metropolis - NY - 12345 - USA", address.FullAddress())
	assert.NotEqual(t, "123 Main St 456, Apt 789 - Metropolis- Downtown,  USA- NY - 12345", address.FullAddress())

}
