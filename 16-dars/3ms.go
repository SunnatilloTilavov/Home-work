package main

import "fmt"

type Course struct {
	Name  string
	Course int
}

func main() {
	courses := []Course{
		{"Course 3", 3},
		{"Course 1", 1},
		{"Course 2", 2},
		{"Course 5", 5},
		{"Course 4", 4},
	}
	
	for i := 0; i < len(courses); i++ {
		for j := i + 1; j < len(courses); j++ {
			if courses[i].Course > courses[j].Course {
				// Almashtirish
				courses[i], courses[j] = courses[j], courses[i]
			}
		}
	}

	for _, course := range courses {
		fmt.Printf("%s: %d\n", course.Name, course.Course)
	}
}
