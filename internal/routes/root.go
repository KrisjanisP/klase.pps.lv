package routes

import (
	"net/http"

	"github.com/KrisjanisP/klase.pps.lv/internal/routes/middleware"
	"github.com/go-chi/chi/v5"
)

func getRootRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.TokenAuthMiddleware)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})
	return r
}
