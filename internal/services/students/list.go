package students

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/KrisjanisP/klase.pps.lv/internal/models"
)

type tempStudent struct {
	Names    []string `json:"names"`
	Password string   `json:"password"`
}

func ListStudents() []models.Student {
	// Open the JSON file
	file, err := os.Open("./data/students.json")
	if err != nil {
		log.Fatalf("Error opening courses.json file: %v", err)
	}
	defer file.Close()

	// Read the file content
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading courses.json file: %v", err)
	}

	// Unmarshal the JSON data into the tempStudents slice
	var tempStudents []tempStudent
	err = json.Unmarshal(data, &tempStudents)
	if err != nil {
		log.Fatalf("Error unmarshalling courses data: %v", err)
	}

	// Map the tempStudents to models.Student
	var students []models.Student
	for _, tStudent := range tempStudents {
		student := models.Student{
			Names:    tStudent.Names,
			Password: tStudent.Password,
		}
		students = append(students, student)
	}

	return students
}
