package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/KrisjanisP/klase.pps.lv/internal/templates"
	"github.com/a-h/templ"
)

func main() {
	component := templates.Login()

	http.Handle("/", templ.Handler(component))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("internal/assets"))))

	fmt.Println("Listening on :3069")
	err := http.ListenAndServe(":3069", nil)
	log.Println(err)
}
