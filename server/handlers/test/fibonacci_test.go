package test

import (
	"testing"

	"github.com/osvaldosilitonga/phiraka/server/handlers"
	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	type testCase struct {
		name           string
		inputRow       int
		inputColumn    int
		expectedResult map[int][]int
		hasError       bool
	}

	testTable := []testCase{
		{
			name:        "normal input",
			inputRow:    3,
			inputColumn: 6,
			expectedResult: map[int][]int{
				0: {0, 1, 1, 2, 3, 5},
				1: {8, 13, 21, 34, 55, 89},
				2: {144, 233, 377, 610, 987, 1597},
			},
			hasError: false,
		},
		{
			name:        "input row 1, col 1",
			inputRow:    1,
			inputColumn: 1,
			expectedResult: map[int][]int{
				0: {0},
			},
			hasError: false,
		},
		{
			name:        "input row 1, col 2",
			inputRow:    1,
			inputColumn: 2,
			expectedResult: map[int][]int{
				0: {0, 1},
			},
			hasError: false,
		},
		{
			name:        "input row 2, col 1",
			inputRow:    2,
			inputColumn: 1,
			expectedResult: map[int][]int{
				0: {0},
				1: {1},
			},
			hasError: false,
		},
		{
			name:           "input with column 0",
			inputRow:       1,
			inputColumn:    0,
			expectedResult: nil,
			hasError:       true,
		},
		{
			name:           "input with row 0",
			inputRow:       0,
			inputColumn:    1,
			expectedResult: nil,
			hasError:       true,
		},
	}

	for _, test := range testTable {
		actual, err := handlers.Fibonacci(test.inputRow, test.inputColumn)
		assert.Equal(t, test.expectedResult, actual, test.name)

		if test.hasError {
			assert.NotNil(t, err, test.name)
		} else {
			assert.Nil(t, err, test.name)
		}
	}
}
