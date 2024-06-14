package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntersectSlices(t *testing.T) {
	tests := []struct {
		name     string
		slices   [][]int
		expected []int
	}{
		{
			name:     "Один слайс - получаем только его",
			slices:   [][]int{},
			expected: []int{},
		},
		{
			slices: [][]int{
				{1, 2, 3, 2},
			},
			expected: []int{1, 2, 3},
		},
		{
			name: "Два слайса с общими элементами - получаем отсортированное пересечение",
			slices: [][]int{
				{1, 2, 3, 2},
				{3, 2},
			},
			expected: []int{2, 3},
		},
		{
			name: "Три слайса, один пустой - получаем пустой слайс",
			slices: [][]int{
				{1, 2, 3, 2},
				{3, 2},
				{},
			},
			expected: []int{},
		},
		{
			name: "Три слайса с одним общим элементом",
			slices: [][]int{
				{1, 2, 3, 4},
				{3, 4, 5, 6},
				{4, 5, 6, 7},
			},
			expected: []int{4},
		},
	}

	for _, tt := range tests {
		ast := assert.New(t)

		ast.Equal(tt.expected, intersectSlices(tt.slices...), "значение"+
			":%v", tt)
	}
}
