package main

import "fmt"

type Student struct {
	name        string
	scholarship int
	course      int
}

func (s *Student) sum() int {
	total := 0
	if s.course == 2 {
		total += s.scholarship
	}
	return total
}
// array holida kiritish
func main() {
Student:=[]Student{
	Student{
		"Zafar",
		1000,
		3,
	}
	Student{
		"Zafar",
		1000,
		3,
	}
	Student{
		"Zafar",
		1000,
		3,
	}
}


	student1:=Student{
		name : "student1",
		scholarship : 100,
		course:2,
	}
	student2:=Student{
		name : "student2",
		scholarship : 115,
		course:3,
	}
	student3:=Student{
		name : "student3",
		scholarship : 150,
		course:1,
	}
	student4:=Student{
		name : "student4",
		scholarship : 120,
		course:2,
	}
	student5:=Student{
		name : "student5",
		scholarship : 160,
		course:2,
	}
	student6:=Student{
		name : "student6",
		scholarship : 190,
		course:4,
	}
	student7:=Student{
		name : "student7",
		scholarship : 140,
		course:2,
	}
	student8:=Student{
		name : "student8",
		scholarship : 100,
		course:1,
	}
	student9:=Student{
		name : "student9",
		scholarship : 100,
		course:3,
	}
	student10:=Student{
		name : "student10",
		scholarship : 120,
		course:2,
	}
	total := student1.sum() + student2.sum()+student3.sum()+student4.sum()+student5.sum()+student6.sum()+student7.sum()+student8.sum()+student9.sum()+student10.sum()

	fmt.Printf("Umumiy: %v\n", total)
}