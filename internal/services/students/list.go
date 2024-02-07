package users

import "github.com/KrisjanisP/klase.pps.lv/internal/models"

func ListUsers() []models.Student {
	return []models.Student{
		{
			Name: "Krisjanis",
			Age:  20,
		},
		{
			Name: "Janis",
			Age:  20,
		},
	}
}
