package main

import "fmt"

type Student1 struct {
	name        string
	scholarship int
	course      int
}

func (s *Student1) Name5() string {
	if len(s.name) < 5 {
		return s.name
	}
	return ""
}

func main() {
	student1:=Student1{
		name : "ent1",
		scholarship : 100,
		course:2,
	}
	student2:=Student1{
		name : "student2",
		scholarship : 115,
		course:3,
	}
	student3:=Student1{
		name : "ent3",
		scholarship : 150,
		course:1,
	}
	student4:=Student1{
		name : "student4",
		scholarship : 120,
		course:2,
	}
	student5:=Student1{
		name : "ent5",
		scholarship : 160,
		course:2,
	}
	student6:=Student1{
		name : "student6",
		scholarship : 190,
		course:4,
	}
	student7:=Student1{
		name : "ent7",
		scholarship : 140,
		course:2,
	}
	student8:=Student1{
		name : "student8",
		scholarship : 100,
		course:1,
	}
	student9:=Student1{
		name : "ent9",
		scholarship : 100,
		course:3,
	}
	student10:=Student1{
		name : "student10",
		scholarship : 120,
		course:2,
	}
	//for i:=1;i<=10;i++{
	//	name:=" "
	//	name=student(i).Name5
	//	fmt.Println(name)
	//}

	name := student1.Name5()+ " "+student2.Name5()+" "+student3.Name5()+" "+student4.Name5()+" "+student5.Name5()+" "+student6.Name5()+" "+student7.Name5()+" "+student8.Name5()+" "+student9.Name5()+" "+student10.Name5()
	fmt.Println("Student name:",name)
}
