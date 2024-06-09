package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
)

/**
Для предыдущей задачи необходимо вывести сводную таблицу по всем предметам в виде:
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
*/

func buildRatingBySubjectsTable() {
	file, err := os.ReadFile("dz3.json")
	if err != nil {
		log.Fatal(err)
	}

	var data Data
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	// Создание мапы для хранения результатов каждого предмета по каждому классу
	subjectResults := make(map[int]map[int][]int)
	for _, object := range data.Objects {
		subjectResults[object.ID] = make(map[int][]int)
	}

	// Создание мап для быстрого доступа к студентам и предметам по ID
	studentMap := make(map[int]Student)
	objectMap := make(map[int]Object)

	for _, student := range data.Students {
		studentMap[student.ID] = student
	}

	for _, object := range data.Objects {
		objectMap[object.ID] = object
	}

	// Заполнение мапы результатами для каждого класса
	for _, result := range data.Results {
		student := studentMap[result.StudentID]
		object := objectMap[result.ObjectID]

		subjectResults[object.ID][student.Grade] = append(subjectResults[object.ID][student.Grade], result.Result)
	}

	// Вывод таблиц для каждого предмета
	for objectID, grades := range subjectResults {
		object := objectMap[objectID]
		fmt.Println("________________")
		fmt.Printf("%-9s | Mean\n", object.Name)
		fmt.Println("________________")

		var totalResults []int
		var gradesList []int
		for grade := range grades {
			gradesList = append(gradesList, grade)
		}
		sort.Ints(gradesList)

		for _, grade := range gradesList {
			gradeResult := grades[grade]

			mean := calculateMean(gradeResult)
			totalResults = append(totalResults, gradeResult...)
			fmt.Printf("%-9s | %.2f\n", fmt.Sprintf("%d grade", grade), mean)
		}

		overallMean := calculateMean(totalResults)
		fmt.Println("________________")
		fmt.Printf("%-9s | %.2f\n", "mean", overallMean)
		fmt.Println("________________")
	}
}

func calculateMean(values []int) float64 {
	if len(values) == 0 {
		return 0
	}

	sum := 0
	for _, result := range values {
		sum += result
	}

	return float64(sum) / float64(len(values))
}
