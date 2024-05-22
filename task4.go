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
	subjectResults := make(map[string]map[int][]int)
	for _, object := range data.Objects {
		subjectResults[object.Name] = make(map[int][]int)
	}

	// Заполнение мапы результатами для каждого класса
	for _, result := range data.Results {
		object := findObjectByID(result.ObjectID, data.Objects)
		student := findStudentByID(result.StudentID, data.Students)

		subjectResults[object.Name][student.Grade] = append(subjectResults[object.Name][student.Grade], result.Result)
	}

	// Вывод таблиц для каждого предмета
	for subject, grades := range subjectResults {
		fmt.Println("________________")
		fmt.Printf("%-9s | Mean\n", subject)
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

func calculateMean(results []int) float64 {
	sum := 0
	for _, result := range results {
		sum += result
	}

	return float64(sum) / float64(len(results))
}
