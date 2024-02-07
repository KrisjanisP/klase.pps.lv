package routes

import (
	"net/http"

	"github.com/KrisjanisP/klase.pps.lv/internal/services/courses"
	"github.com/KrisjanisP/klase.pps.lv/internal/templates/pages"
	"github.com/a-h/templ"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func NewCompletedRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./internal/assets"))))

	r.Handle("/login", templ.Handler(pages.Login(courses.ListCourses())))
	r.Post("/login", loginPostHandlerFunc)
	r.Post("/logout", logoutPostHandlerFunc)

	r.Group(func(r chi.Router) {
		r.Use(tokenAuthMiddleware)
		r.Handle("/", http.RedirectHandler("/home", http.StatusSeeOther))
		r.Handle("/home", templ.Handler(pages.Home()))
	})

	return r
}
