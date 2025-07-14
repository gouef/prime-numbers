package tests

import (
	"fmt"
	primeNumbers "github.com/gouef/prime-numbers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		size     int
		expected []int
	}{
		{1, []int{2}},
		{2, []int{2, 3}},
		{3, []int{2, 3, 5}},
		{4, []int{2, 3, 5, 7}},
		{5, []int{2, 3, 5, 7, 11}},
		{6, []int{2, 3, 5, 7, 11, 13}},
		{7, []int{2, 3, 5, 7, 11, 13, 17}},
		{8, []int{2, 3, 5, 7, 11, 13, 17, 19}},
		{9, []int{2, 3, 5, 7, 11, 13, 17, 19, 23}},
		{10, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}},
		{11, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Size: %v", tt.size), func(t *testing.T) {
			assert.Equal(t, tt.expected, primeNumbers.Generate(tt.size))
		})
	}
}
