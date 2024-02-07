package routes

import (
	"fmt"
	"net/http"

	"github.com/KrisjanisP/klase.pps.lv/internal/config"
	"github.com/KrisjanisP/klase.pps.lv/internal/models"
	"github.com/dgrijalva/jwt-go"
)

// Middleware to validate token
func tokenAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				fmt.Fprintln(w, "Access denied. Redirected. No token provided.")
				return
			}
			// For any other type of error, return a bad request status
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		tokenStr := cookie.Value
		claims := &models.JWTClaims{}

		// Parse the JWT string and store the result in `claims`.
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return config.GetJWTKey(), nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// If everything is good with the token, pass the request to the next handler.
		next.ServeHTTP(w, r)
	})
}
