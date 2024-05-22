package main

import "sort"

/*
 1. Напишите функцию, которая находит пересечение неопределенного количества слайсов типа int.

Каждый элемент в пересечении должен быть уникальным. Слайс-результат должен быть отсортирован в восходящем порядке.
Примеры:
1. Если на вход подается только 1 слайс [1, 2, 3, 2], результатом должен быть слайс [1, 2, 3].
2. Вход: 2 слайса [1, 2, 3, 2] и [3, 2], результат - [2, 3].
3. Вход: 3 слайса [1, 2, 3, 2], [3, 2] и [], результат - [].
*/
func intersectSlices(slices ...[]int) []int {
	if len(slices) == 0 {
		return []int{}
	}

	// Мапа для подсчета частоты появления каждого элемента
	elementCount := make(map[int]int)

	// Мапа для отслеживания элементов первого слайса
	firstSliceSet := make(map[int]struct{})
	for _, num := range slices[0] {
		firstSliceSet[num] = struct{}{}
		elementCount[num] = 1
	}

	// Обрабатываем остальные слайсы
	for i := 1; i < len(slices); i++ {
		sliceSet := make(map[int]struct{})
		for _, num := range slices[i] {
			sliceSet[num] = struct{}{}
		}

		for num := range firstSliceSet {
			// Если элемент актуального слайса есть в наборе, увеличиваем счетчик для этого элемента в исходном наборе из первого слайса
			if _, found := sliceSet[num]; found {
				elementCount[num]++
			}
		}
	}

	// Составляем результат
	var result []int
	for num, count := range elementCount {
		if count == len(slices) {
			result = append(result, num)
		}
	}

	sort.Ints(result)

	if len(result) == 0 {
		return []int{}
	}

	return result
}
