package primeNumbers

// Generate returns []int of prime numbers.
// Example:
//
//	size := 3
//	listPrimeNumbers := primeNumbers.Generate(size) // returns []int{2, 3, 5}
func Generate(size int) []int {
	var res []int

	i := 2
	for len(res) < size {
		n := Validate(i)

		if n {
			res = append(res, i)
		}

		i++
	}

	return res
}
