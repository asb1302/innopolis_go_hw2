package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

/*
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
*/

type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Grade int    `json:"grade"`
}

type Object struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Result struct {
	ObjectID  int `json:"object_id"`
	StudentID int `json:"student_id"`
	Result    int `json:"result"`
}

type Data struct {
	Students []Student `json:"students"`
	Objects  []Object  `json:"objects"`
	Results  []Result  `json:"results"`
}

func buildRatingTable() {
	file, err := os.ReadFile("dz3.json")
	if err != nil {
		log.Fatal(err)
	}

	var data Data
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("_____________________________________")
	fmt.Printf("%-15s | %-6s | %-10s | %-7s\n", "Student name", "Grade", "Object", "Result")
	fmt.Println("_____________________________________")
	for _, result := range data.Results {
		student := findStudentByID(result.StudentID, data.Students)
		object := findObjectByID(result.ObjectID, data.Objects)

		fmt.Printf("%-15s | %-6d | %-10s | %-7d\n", student.Name, student.Grade, object.Name, result.Result)
	}
	fmt.Println("_____________________________________")
}

func findStudentByID(id int, students []Student) Student {
	for _, student := range students {
		if student.ID == id {
			return student
		}
	}

	return Student{}
}

func findObjectByID(id int, objects []Object) Object {
	for _, object := range objects {
		if object.ID == id {
			return object
		}
	}

	return Object{}
}
