package routes

import (
	"net/http"
	"time"

	"github.com/KrisjanisP/klase.pps.lv/internal/config"
	"github.com/KrisjanisP/klase.pps.lv/internal/models"
	"github.com/dgrijalva/jwt-go"
)

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
