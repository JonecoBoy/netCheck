package utils

import (
	"math/rand"
)

func RandomDigits(n int) string {
	digits := ""
	for i := 0; i < n; i++ {
		digits += string('0' + rune(rand.Intn(10)))
	}
	return digits
}
