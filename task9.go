package main

import (
	"golang.org/x/exp/constraints"
)

/*
Реализуйте тип-дженерик Numbers, который является слайсом численных типов.

Реализуйте следующие методы для этого типа:
* суммирование всех элементов,
* произведение всех элементов,
* сравнение с другим слайсом на равность,
* проверка аргумента, является ли он элементом массива, если да - вывести индекс первого найденного элемента,
* удаление элемента массива по значению,
* удаление элемента массива по индексу.
*/

type Numbers[T constraints.Integer | constraints.Float] []T

func (n Numbers[T]) Sum() T {
	var sum T
	for _, v := range n {
		sum += v
	}
	return sum
}

func (n Numbers[T]) Multiple() T {
	if len(n) == 0 {
		return 0
	}
	product := T(1)
	for _, v := range n {
		product *= v
	}
	return product
}

func (n Numbers[T]) Equals(other Numbers[T]) bool {
	if len(n) != len(other) {
		return false
	}
	for i, v := range n {
		if v != other[i] {
			return false
		}
	}
	return true
}

func (n Numbers[T]) IndexOf(value T) (int, bool) {
	for i, v := range n {
		if v == value {
			return i, true
		}
	}
	return -1, false
}

func (n *Numbers[T]) RemoveValue(value T) bool {
	index, found := n.IndexOf(value)
	if found {
		*n = append((*n)[:index], (*n)[index+1:]...)
		return true
	}
	return false
}

func (n *Numbers[T]) RemoveAt(index int) bool {
	if index < 0 || index >= len(*n) {
		return false
	}
	*n = append((*n)[:index], (*n)[index+1:]...)
	return true
}
