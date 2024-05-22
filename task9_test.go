package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := Numbers[int]{1, 2, 3, 4, 5}

	ast := assert.New(t)

	ast.Equal(15, numbers.Sum())
}

func TestMultiple(t *testing.T) {
	numbers := Numbers[int]{1, 2, 3, 4, 5}

	ast := assert.New(t)

	ast.Equal(true, numbers.Equals(Numbers[int]{1, 2, 3, 4, 5}))
}

func TestEquals(t *testing.T) {
	numbers := Numbers[int]{1, 2, 3, 4, 5}

	ast := assert.New(t)

	ast.Equal(true, numbers.Equals(Numbers[int]{1, 2, 3, 4, 5}))
}

func TestIndexOf(t *testing.T) {
	numbers := Numbers[int]{1, 2, 3, 4, 5}

	ast := assert.New(t)

	index, _ := numbers.IndexOf(3)
	ast.Equal(2, index)
}

func TestRemoveValue(t *testing.T) {
	numbers := Numbers[int]{1, 2, 3, 4, 5}

	ast := assert.New(t)

	result := numbers.RemoveValue(3)
	ast.Equal(true, result)
	ast.Equal(true, numbers.Equals(Numbers[int]{1, 2, 4, 5}))
}

func TestRemoveAt(t *testing.T) {
	numbers := Numbers[int]{1, 2, 3, 4, 5}

	ast := assert.New(t)

	result := numbers.RemoveAt(2)
	ast.Equal(true, result)
	ast.Equal(true, numbers.Equals(Numbers[int]{1, 2, 4, 5}))
}
