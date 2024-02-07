package students

import "github.com/KrisjanisP/klase.pps.lv/internal/models"

func ListStudents() []models.Student {
	return []models.Student{
		{
			Names:    []string{},
			Password: "asdf",
		},
	}
}
