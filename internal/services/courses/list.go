package courses

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/KrisjanisP/klase.pps.lv/internal/models"
)

// Define a temporary struct that matches the JSON structure
type tempCourse struct {
	ID       string     `json:"id"`
	Teachers [][]string `json:"teachers"`
	Students [][]string `json:"students"`
}

func ListCourses() []models.Course {
	// Open the JSON file
	file, err := os.Open("./data/courses.json")
	if err != nil {
		log.Fatalf("Error opening courses.json file: %v", err)
	}
	defer file.Close()

	// Read the file content
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading courses.json file: %v", err)
	}

	// Unmarshal the JSON data into the tempCourses slice
	var tempCourses []tempCourse
	err = json.Unmarshal(data, &tempCourses)
	if err != nil {
		log.Fatalf("Error unmarshalling courses data: %v", err)
	}

	// Map the tempCourses to models.Course
	var courses []models.Course
	for _, tCourse := range tempCourses {
		course := models.Course{
			ID:           tCourse.ID,
			TeacherNames: tCourse.Teachers,
			StudentNames: tCourse.Students,
		}
		courses = append(courses, course)
	}

	// Return the mapped courses
	return courses
}
