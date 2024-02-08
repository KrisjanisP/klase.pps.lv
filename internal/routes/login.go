package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/KrisjanisP/klase.pps.lv/internal/config"
	"github.com/KrisjanisP/klase.pps.lv/internal/models"
	"github.com/KrisjanisP/klase.pps.lv/internal/services/courses"
	"github.com/KrisjanisP/klase.pps.lv/internal/services/students"
	"github.com/KrisjanisP/klase.pps.lv/internal/templates/partials"
	"github.com/dgrijalva/jwt-go"
)

func loginPostHandlerFunc(w http.ResponseWriter, r *http.Request) {
	studentList := students.ListStudents()

	// Parse form data
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Retrieve the JSON-encoded name list and password from the form
	studentNameJSON := r.FormValue("student")
	password := r.FormValue("password")

	var names []string
	if err := json.Unmarshal([]byte(studentNameJSON), &names); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Println(studentNameJSON, password)

	// Authentication logic
	var authenticatedStudent *models.Student
	for _, student := range studentList {
		if reflect.DeepEqual(student.Names, names) {
			if student.Password == password {
				// If the names match and the password is correct, mark as authenticated
				authenticatedStudent = &student
				break
			} else {
				msg := "Invalid password"
				partials.LoginCard(courses.ListCourses(), &msg).Render(r.Context(), w)
				return
			}
		}
	}

	if authenticatedStudent == nil {
		msg := "Invalid student"
		partials.LoginCard(courses.ListCourses(), &msg).Render(r.Context(), w)
		return
	}

	// After validating user credentials, create the JWT token
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &models.JWTClaims{
		Names: authenticatedStudent.Names,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(config.GetJWTKey())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func logoutPostHandlerFunc(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",         // The name of the token cookie set during login
		Value:    "",              // Clear the token value
		Expires:  time.Unix(0, 0), // Set the expiration time to the past
		MaxAge:   -1,              // Immediately expire the cookie
		HttpOnly: true,            // Make the cookie inaccessible to JavaScript running in the browser
		Path:     "/",             // Ensure the cookie is deleted for the whole domain
	})

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
