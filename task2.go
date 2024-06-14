package main

import "sort"

/*
Напишите функцию подсчета каждого голоса за кандидата. Входной аргумент - массив с именами кандидатов.
Результативный - массив структуры Candidate, отсортированный по убыванию количества голосов.
Пример.
Вход: ["Ann", "Kate", "Peter", "Kate", "Ann", "Ann", "Helen"]
Вывод: [{Ann, 3}, {Kate, 2}, {Peter, 1}, {Helen, 1}]
*/

type Candidate struct {
	Name  string
	Votes int
}

func countVotes(votes []string) []Candidate {
	if len(votes) == 0 {
		return []Candidate{}
	}

	// Считаем голоса
	counter := make(map[string]int)
	for _, name := range votes {
		if _, found := counter[name]; found {
			counter[name]++

			continue
		}

		counter[name] = 1
	}

	// делаем финальную мапу
	var result []Candidate
	for name, votesCounter := range counter {
		result = append(result, Candidate{Name: name, Votes: votesCounter})
	}

	// Сортируем результат по убыванию голосов
	sort.SliceStable(result, func(i, j int) bool {
		return result[i].Votes > result[j].Votes
	})

	return result
}
