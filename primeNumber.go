package primeNumbers

import (
	"math"
)

func Validate(number int) bool {
	if number == 2 {
		return true
	}

	if number < 2 || number%2 == 0 {
		return false
	}

	rootNumber := int(math.Sqrt(float64(number)))

	for i := 3; i <= rootNumber; i += 2 {
		v := float64(number) / float64(i)
		if v == float64(int64(v)) {
			return false
		}
	}

	return true
}
