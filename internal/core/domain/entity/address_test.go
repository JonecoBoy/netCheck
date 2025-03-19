package entity

import (
	pkgEntity "github.com/jonecoboy/netCheck/pkg/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddress(t *testing.T) {
	id := pkgEntity.NewId()
	street := "123 Main St"
	number := "456"
	complement := "Apt 789"
	neighborhood := "Downtown"
	city := "Metropolis"
	state := "NY"
	zip := "12345"
	country := "USA"

	address := Address{
		id:           id,
		street:       street,
		number:       number,
		complement:   complement,
		neighborhood: neighborhood,
		city:         city,
		state:        state,
		zip:          zip,
		country:      country,
	}

	assert.Equal(t, id, address.id)
	assert.Equal(t, street, address.street)
	assert.Equal(t, number, address.number)
	assert.Equal(t, complement, address.complement)
	assert.Equal(t, neighborhood, address.neighborhood)
	assert.Equal(t, city, address.city)
	assert.Equal(t, state, address.state)
	assert.Equal(t, zip, address.zip)
	assert.Equal(t, country, address.country)
}
