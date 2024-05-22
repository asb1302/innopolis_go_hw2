package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountVotes(t *testing.T) {
	tests := []struct {
		name     string
		votes    []string
		expected []Candidate
	}{
		{
			name:     "Нет голосов - пустой массив",
			votes:    []string{},
			expected: []Candidate{},
		},
		{
			name:     "Считаем голоса",
			votes:    []string{"Ann", "Kate", "Peter", "Kate", "Ann", "Ann", "Helen"},
			expected: []Candidate{{"Ann", 3}, {"Kate", 2}, {"Peter", 1}, {"Helen", 1}},
		},
	}

	for _, tt := range tests {
		ast := assert.New(t)

		ast.Equal(tt.expected, countVotes(tt.votes), "значение"+":%v", tt)
	}
}
