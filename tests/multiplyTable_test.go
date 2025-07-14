package tests

import (
	"fmt"
	primeNumbers "github.com/gouef/prime-numbers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMultiplyTable(t *testing.T) {
	tests := []struct {
		size     int
		expected map[int]map[int]int
	}{
		{1, map[int]map[int]int{
			2: {2: 4},
		}},
		{2, map[int]map[int]int{
			2: {2: 4, 3: 6},
			3: {2: 6, 3: 9},
		}},
		{3, map[int]map[int]int{
			2: {2: 4, 3: 6, 5: 10},
			3: {2: 6, 3: 9, 5: 15},
			5: {2: 10, 3: 15, 5: 25},
		}},
		{4, map[int]map[int]int{
			2: {2: 4, 3: 6, 5: 10, 7: 14},
			3: {2: 6, 3: 9, 5: 15, 7: 21},
			5: {2: 10, 3: 15, 5: 25, 7: 35},
			7: {2: 14, 3: 21, 5: 35, 7: 49},
		}},
		{5, map[int]map[int]int{
			2:  {2: 4, 3: 6, 5: 10, 7: 14, 11: 22},
			3:  {2: 6, 3: 9, 5: 15, 7: 21, 11: 33},
			5:  {2: 10, 3: 15, 5: 25, 7: 35, 11: 55},
			7:  {2: 14, 3: 21, 5: 35, 7: 49, 11: 77},
			11: {2: 22, 3: 33, 5: 55, 7: 77, 11: 121},
		}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Size: %v", tt.size), func(t *testing.T) {
			assert.Equal(t, tt.expected, primeNumbers.MultiplyTable(tt.size))
		})
	}
}
