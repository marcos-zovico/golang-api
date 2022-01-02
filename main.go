package main

import (
	"fmt"
	"golang-api/src/router"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Executando api")

	r := router.Generate()

	log.Fatal(http.ListenAndServe(":5000", r))
}
