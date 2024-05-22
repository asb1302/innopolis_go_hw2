package main

/*
Напишите функцию-дженерик IsEqualArrays для comparable типов, которая сравнивает два неотсортированных массива.
Функция выдает булевое значение как результат. true - если массивы равны, false - если нет.
Массивы считаются равными, если в элемент из первого массива существует в другом, и наоборот.
Вне зависимости от расположения.
*/

func IsEqualArrays[T comparable](arr1, arr2 []T) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	elemCount := make(map[T]int)

	for _, elem := range arr1 {
		elemCount[elem]++
	}

	for _, elem := range arr2 {
		if count, found := elemCount[elem]; !found || count == 0 {
			return false
		}
		
		elemCount[elem]--
	}

	return true
}
