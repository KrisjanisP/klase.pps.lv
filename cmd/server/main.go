package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/KrisjanisP/klase.pps.lv/internal/templates/pages"
	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	fs := http.StripPrefix("/static/", http.FileServer(http.Dir("./internal/assets")))
	r.Handle("/static/*", fs) // Use /* to match all files under /static/

	r.Handle("/", templ.Handler(pages.Login()))
	r.Handle("/home", templ.Handler(pages.Home()))

	fmt.Println("Listening on :3069")
	err := http.ListenAndServe(":3069", r)
	log.Println(err)

}
