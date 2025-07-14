package tests

import (
	"fmt"
	primeNumbers "github.com/gouef/prime-numbers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		needle   int
		expected bool
	}{
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{7, true},
		{8, false},
		{9, false},
		{10, false},
		{11, true},
		{12, false},
		{13, true},
		{14, false},
		{15, false},
		{16, false},
		{17, true},
		{18, false},
		{19, true},
		{20, false},
		{21, false},
		{22, false},
		{23, true},
		{24, false},
		{25, false},
		{26, false},
		{27, false},
		{28, false},
		{29, true},
		{30, false},
		{31, true},
		{32, false},
		{33, false},
		{34, false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Number: %v", tt.needle), func(t *testing.T) {
			assert.Equal(t, tt.expected, primeNumbers.Validate(tt.needle))
		})
	}
}
