package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/KrisjanisP/klase.pps.lv/internal/routes"
)

func main() {
	r := routes.NewCompletedRouter()

	fmt.Println("Listening on :3069")
	err := http.ListenAndServe(":3069", r)
	log.Println(err)
}
