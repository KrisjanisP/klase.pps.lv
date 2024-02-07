package main

import (
	"log"

	"github.com/KrisjanisP/klase.pps.lv/internal/services/courses"
	users "github.com/KrisjanisP/klase.pps.lv/internal/services/students"
)

func main() {
	var x = courses.ListCourses()
	log.Println(x)

	var y = users.ListStudents()
	log.Println(y)
}
