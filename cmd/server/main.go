package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/KrisjanisP/klase.pps.lv/internal/routes"
)

const port = 3070

func main() {
	r := routes.NewCompletedRouter()

	log.Printf("Listening on :%d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), r)
	log.Println(err)
}
