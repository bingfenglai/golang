package main

import (
	"fmt"
)

func main() {
	var students []map[string]string

	student1 := map[string]string{
		"name": "εε",
		"gender": "male",
	}

	student2 := map[string]string{
		"name": "εδΈ",
		"gender": "male",
	}
	students = append(students, student1)
	students = append(students, student2)

	for _, student := range students {

		fmt.Println("name: ",student["name"],"gender: ",student["gender"])
	}

}
