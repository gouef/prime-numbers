package primeNumbers

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
