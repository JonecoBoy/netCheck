package entity

import (
	"github.com/jonecoboy/netCheck/internal/utils"
	"math/rand"
)

var sampleStreets = []string{"Main St", "2nd Ave", "Baker St", "Elm St"}
var sampleCities = []string{"Metropolis", "Gotham", "Springfield"}
var sampleStates = []string{"NY", "CA", "TX"}
var sampleCountries = []string{"USA", "Brazil", "Canada"}

func RandomAddress() Address {
	return NewAddress(
		sampleStreets[rand.Intn(len(sampleStreets))],
		utils.RandomDigits(3),
		"Apt "+utils.RandomDigits(2),
		"Downtown",
		sampleCities[rand.Intn(len(sampleCities))],
		sampleStates[rand.Intn(len(sampleStates))],
		randomZip(),
		sampleCountries[rand.Intn(len(sampleCountries))],
	)
}

func randomZip() string {
	return utils.RandomDigits(5) + "-" + utils.RandomDigits(3)
}
