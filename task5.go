package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

/*
Перепишите задачу #3 с использованием структуры-дженерик Cache, изученной на семинаре.
Храните в кеше таблицы студентов и предметов.
*/

type Cache[K comparable, V any] struct {
	m map[K]V
}

func (c *Cache[K, V]) Init() {
	c.m = make(map[K]V)
}

func (c *Cache[K, V]) Set(key K, value V) {
	c.m[key] = value
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
	k, ok := c.m[key]

	return k, ok
}

func buildRatingTableWithCacheGeneric() {
	file, err := os.ReadFile("dz3.json")
	if err != nil {
		log.Fatal(err)
	}

	var data Data
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	studentCache := Cache[int, Student]{}
	studentCache.Init()

	objectCache := Cache[int, Object]{}
	objectCache.Init()

	// Заполнение кэша данными
	for _, student := range data.Students {
		studentCache.Set(student.ID, student)
	}
	for _, object := range data.Objects {
		objectCache.Set(object.ID, object)
	}

	// Вывод данных в виде таблицы
	fmt.Println("_____________________________________")
	fmt.Printf("%-15s | %-6s | %-10s | %-7s\n", "Student name", "Grade", "Object", "Result")
	fmt.Println("_____________________________________")
	for _, result := range data.Results {
		student, studentFound := studentCache.Get(result.StudentID)
		object, objectFound := objectCache.Get(result.ObjectID)
		if studentFound && objectFound {
			fmt.Printf("%-15s | %-6d | %-10s | %-7d\n", student.Name, student.Grade, object.Name, result.Result)
		}
	}
	fmt.Println("_____________________________________")
}
