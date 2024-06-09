package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
)

/*
Перепишите задачу #4 с использованием функций высшего порядка, изученных на лекции.
Желательно реализуйте эти функции самостоятельно.
*/

func buildRatingBySubjectsTableWithHOF() {
	file, err := os.ReadFile("dz3.json")
	if err != nil {
		log.Fatal(err)
	}

	var data Data
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	// Создание мапы для хранения результатов каждого предмета
	subjectResults := make(map[int]map[int][]int)
	for _, object := range data.Objects {
		subjectResults[object.ID] = make(map[int][]int)
	}

	objectMap := make(map[int]Object)
	for _, object := range data.Objects {
		objectMap[object.ID] = object
	}

	// Заполнение мапы результатами
	for _, result := range data.Results {
		object := findObjectByIDWithHOF(result.ObjectID, data.Objects)
		student := findStudentByIDWithHOF(result.StudentID, data.Students)

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

		resultsStrings := Map(gradesList, func(grade int) string {
			results := grades[grade]
			mean := calculateMeanWithHOF(results)
			totalResults = append(totalResults, results...)

			return fmt.Sprintf("%-9s | %.2f", fmt.Sprintf("%d grade", grade), mean)
		})

		for _, resultString := range resultsStrings {
			fmt.Println(resultString)
		}

		overallMean := calculateMeanWithHOF(totalResults)
		fmt.Println("________________")
		fmt.Printf("%-9s | %.2f\n", "mean", overallMean)
		fmt.Println("________________")
	}
}

func findStudentByIDWithHOF(id int, students []Student) Student {
	filteredStudents := Filter(students, func(s Student) bool {
		return s.ID == id
	})

	if len(filteredStudents) > 0 {
		return filteredStudents[0]
	}

	return Student{}
}

func findObjectByIDWithHOF(id int, objects []Object) Object {
	filteredObjects := Filter(objects, func(o Object) bool {
		return o.ID == id
	})
	if len(filteredObjects) > 0 {
		return filteredObjects[0]
	}
	return Object{}
}

func calculateMeanWithHOF(values []int) float64 {
	if len(values) == 0 {
		return 0
	}

	sum := Reduce(values, 0, func(r int, acc int) int {
		return r + acc
	})

	return float64(sum) / float64(len(values))
}

func Filter[T any](s []T, f func(T) bool) []T {
	var r []T
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}

	return r
}

func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	r := make([]T2, len(s))
	for i, v := range s {
		r[i] = f(v)
	}

	return r
}

func Reduce[T1, T2 any](s []T1, init T2, f func(T1, T2) T2) T2 {
	r := init
	for _, v := range s {
		r = f(v, r)
	}

	return r
}
