package main

import (
	"log"

	"github.com/KrisjanisP/klase.pps.lv/internal/services/courses"
)

func main() {
	var x = courses.ListCourses()
	log.Println(x)
}
