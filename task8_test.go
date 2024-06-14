package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsEqualArrays(t *testing.T) {
	tests := []struct {
		name     string
		arr1     []any
		arr2     []any
		expected bool
	}{
		{
			name:     "Массивы целых чисел одинаковые - равны",
			arr1:     []any{1, 2, 3, 4, 5},
			arr2:     []any{1, 2, 3, 4, 5},
			expected: true,
		},
		{
			name:     "Массивы целых чисел одинаковые, но в разном порядке - равны",
			arr1:     []any{2, 1, 3, 4, 5},
			arr2:     []any{1, 2, 3, 4, 5},
			expected: true,
		},
		{
			name:     "Массивы целых чисел разные",
			arr1:     []any{9, 6, 3, 4, 5},
			arr2:     []any{1, 2, 3, 4, 5},
			expected: false,
		},
		{
			name:     "Массивы чисел c плавающей запятой одинаковые - равны",
			arr1:     []any{1.1, 2.2, 3.3, 4.4, 5.5},
			arr2:     []any{1.1, 2.2, 3.3, 4.4, 5.5},
			expected: true,
		},
		{
			name:     "Массивы чисел c плавающей запятой одинаковые, но в разном порядке - равны",
			arr1:     []any{1.1, 2.2, 3.3, 4.4, 5.5},
			arr2:     []any{2.2, 1.1, 3.3, 4.4, 5.5},
			expected: true,
		},
		{
			name:     "Массивы чисел c плавающей запятой разные",
			arr1:     []any{1.1, 2.2, 3.3, 4.4, 5.5},
			arr2:     []any{6.6, 1.1, 3.3, 4.4, 5.5},
			expected: false,
		},
		{
			name:     "Массивы строк одинаковые - равны",
			arr1:     []any{"one", "two", "three"},
			arr2:     []any{"one", "two", "three"},
			expected: true,
		},
		{
			name:     "Массивы строк одинаковые, но в разном порядке - равны",
			arr1:     []any{"one", "two", "three"},
			arr2:     []any{"two", "one", "three"},
			expected: true,
		},
		{
			name:     "Массивы строк разные",
			arr1:     []any{"one", "two", "three"},
			arr2:     []any{"four", "one", "three"},
			expected: false,
		},
	}

	for _, tt := range tests {
		ast := assert.New(t)

		ast.Equal(tt.expected, IsEqualArrays(tt.arr1, tt.arr2), "значение"+":%v", tt)
	}
}
