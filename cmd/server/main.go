package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/KrisjanisP/klase.pps.lv/internal/templates/pages"
	"github.com/a-h/templ"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var jwtKey = []byte("your_secret_key")

type UserCredentials struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Password  string `json:"password"`
}

type Claims struct {
	UserID string `json:"userid"`
	jwt.StandardClaims
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	fs := http.StripPrefix("/static/", http.FileServer(http.Dir("./internal/assets")))
	r.Handle("/static/*", fs) // Use /* to match all files under /static/

	r.Handle("/", templ.Handler(pages.Login()))
	r.Handle("/home", templ.Handler(pages.Home()))

	r.Post("/login", LoginHandler)

	fmt.Println("Listening on :3069")
	err := http.ListenAndServe(":3069", r)
	log.Println(err)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Example: User authentication logic goes here.
	// For simplicity, we're skipping the actual authentication.
	// var creds UserCredentials
	// In a real application, you'd decode these credentials from the request body.

	// Example user ID, in a real scenario, this would be retrieved from your user management system upon successful authentication.
	userID := "unique_user_id"

	// After validating user credentials, create the JWT token
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
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
