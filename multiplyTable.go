package primeNumbers

// MultiplyTable returns map[int]map[int]int of prime numbers
// Example:
//
//		size := 3
//		numberTable := primeNumbers.MultiplyTable(size)
//			// returns map[int]map[int]int{
//	 	// 		2: {2: 4, 3: 6, 5: 10},
//			//  	3: {2: 6, 3: 9, 5: 15},
//			//  	5: {2: 10, 3: 15, 5: 25},
//			// }
//
//		for i1, l := range numberTable {
//			for i2, m := range l {
//				log.Printf("Number %v x %v = %v", i1, i2, m)
//			}
//		}
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
