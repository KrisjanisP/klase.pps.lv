package routes

import (
	"net/http"
	"time"

	"github.com/KrisjanisP/klase.pps.lv/internal/config"
	"github.com/KrisjanisP/klase.pps.lv/internal/models"
	"github.com/KrisjanisP/klase.pps.lv/internal/templates/pages"
	"github.com/a-h/templ"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
)

func getLoginRouter() chi.Router {
	r := chi.NewRouter()

	r.Handle("/login", templ.Handler(pages.Login(make([]models.Course, 0))))
	r.Post("/login", loginPostHandlerFunc)

	return r
}

func loginPostHandlerFunc(w http.ResponseWriter, r *http.Request) {
	// Example: User authentication logic goes here.
	// For simplicity, we're skipping the actual authentication.
	// var creds UserCredentials
	// In a real application, you'd decode these credentials from the request body.

	// After validating user credentials, create the JWT token
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &models.JWTClaims{
		FirstName: "John",
		LastName:  "Doe",
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
