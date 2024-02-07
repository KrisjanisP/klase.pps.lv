package routes

import (
	"net/http"

	"github.com/KrisjanisP/klase.pps.lv/internal/templates/pages"
	"github.com/a-h/templ"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func NewCompletedRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Handle("/", getRootRouter())
	r.Handle("/home", templ.Handler(pages.Home()))
	r.Mount("/login", getLoginRouter())

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./internal/assets"))))

	return r
}
