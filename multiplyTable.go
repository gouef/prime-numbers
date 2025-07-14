package primeNumbers

func MultiplyTable(size int) map[int]map[int]int {
	numbers := Generate(size)
	table := make(map[int]map[int]int)

	for _, a := range numbers {
		table[a] = make(map[int]int)
		for _, b := range numbers {
			table[a][b] = a * b
		}
	}

	return table
}
