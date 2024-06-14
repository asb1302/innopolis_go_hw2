package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

/*
Выведите в консоль круглых отличников из числа студентов, используя функцию Filter.
Вывод реализуйте как в задаче #3.
*/

func buildRatingTableForExcellentStudents() {
	file, err := os.ReadFile("dz3.json")
	if err != nil {
		log.Fatal(err)
	}

	var data Data
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	// Определение функции фильтрации для круглых отличников
	isExcellentStudent := func(student Student) bool {
		studentResults := Filter(data.Results, func(result Result) bool {
			return result.StudentID == student.ID
		})

		for _, result := range studentResults {
			if result.Result != 5 {
				return false
			}
		}

		return true
	}

	// Фильтрация студентов для получения только круглых отличников
	excellentStudents := Filter(data.Students, isExcellentStudent)

	fmt.Println("_____________________________________")
	fmt.Println("Student name  | Grade | Object    |   Result")
	fmt.Println("_____________________________________")

	// Создание мап для быстрого доступа к предметам по ID
	objectMap := make(map[int]Object)

	for _, object := range data.Objects {
		objectMap[object.ID] = object
	}

	for _, student := range excellentStudents {
		studentResults := Filter(data.Results, func(result Result) bool {
			return result.StudentID == student.ID
		})
		for _, result := range studentResults {
			object := objectMap[result.ObjectID]

			fmt.Printf("%-14s | %-5d | %-9s | %-7d\n", student.Name, student.Grade, object.Name, result.Result)
		}
	}
}
