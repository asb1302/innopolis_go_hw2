# hw2

## task1

Напишите функцию, которая находит пересечение неопределенного количества слайсов типа int.
Каждый элемент в пересечении должен быть уникальным. Слайс-результат должен быть отсортирован в восходящем порядке.

Примеры:
1. Если на вход подается только 1 слайс [1, 2, 3, 2], результатом должен быть слайс [1, 2, 3].
2. Вход: 2 слайса [1, 2, 3, 2] и [3, 2], результат - [2, 3].
3. Вход: 3 слайса [1, 2, 3, 2], [3, 2] и [], результат - [].
   */

## task2

Напишите функцию подсчета каждого голоса за кандидата. Входной аргумент - массив с именами кандидатов.
Результативный - массив структуры Candidate, отсортированный по убыванию количества голосов.
Пример.
Вход: ["Ann", "Kate", "Peter", "Kate", "Ann", "Ann", "Helen"]
Вывод: [{Ann, 3}, {Kate, 2}, {Peter, 1}, {Helen, 1}]


```go
type Candidate struct {
	Name  string
	Votes int
}
```

## task3

У учеников старших классов прошел контрольный срез по нескольким предметам. Выведите данные в читаемом виде
в таблицу вида

```
_____________________________________
Student name  | Grade | Object    |   Result
____________________________________
Ann			  |     9 | Math	  |  4
Ann 		  |     9 | Biology   |  4
...
```

Вводные данные представлены в файле dz3.json

## task4

Для предыдущей задачи необходимо вывести сводную таблицу по всем предметам в виде:

```
________________
Math	 | Mean
________________
9 grade | 4.5
10 grade | 5
11 grade | 3.5
________________
mean     | 4		- среднее значение среди всех учеников
________________
________________
Biology	 | Mean
________________
...
```

Вводные данные представлены в файле dz3.json

## task5

Перепишите задачу #3 с использованием структуры-дженерик Cache, изученной на семинаре.
Храните в кеше таблицы студентов и предметов.


## task6

Перепишите задачу #4 с использованием функций высшего порядка, изученных на лекции.
Желательно реализуйте эти функции самостоятельно.

## task7
Выведите в консоль круглых отличников из числа студентов, используя функцию Filter.
Вывод реализуйте как в задаче #3.


## task8

Напишите функцию-дженерик IsEqualArrays для comparable типов, которая сравнивает два неотсортированных массива.
Функция выдает булевое значение как результат. true - если массивы равны, false - если нет.
Массивы считаются равными, если в элемент из первого массива существует в другом, и наоборот.
Вне зависимости от расположения.

## task9

Реализуйте тип-дженерик Numbers, который является слайсом численных типов.

Реализуйте следующие методы для этого типа:
* суммирование всех элементов,
* произведение всех элементов,
* сравнение с другим слайсом на равность,
* проверка аргумента, является ли он элементом массива, если да - вывести индекс первого найденного элемента,
* удаление элемента массива по значению,
* удаление элемента массива по индексу.
