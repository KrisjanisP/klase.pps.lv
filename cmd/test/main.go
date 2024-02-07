package main

import (
	"log"

	"github.com/KrisjanisP/klase.pps.lv/internal/services/courses"
	"github.com/KrisjanisP/klase.pps.lv/internal/services/users"
)

func main() {
	var x = courses.ListCourses()
	log.Println(x)

	var y = users.ListUsers()
	log.Println(y)
}
