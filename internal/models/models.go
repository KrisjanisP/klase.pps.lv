package models

type Course struct {
	ID           string
	TeacherNames [][]string
	StudentNames [][]string
}

type Student struct {
	Names    []string
	Password string
}
