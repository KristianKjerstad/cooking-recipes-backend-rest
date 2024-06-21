package main

import (
	"fmt"
	"net/http"
	"rest/rest"
)

func main() {
	app := rest.New()
	fmt.Println("starting server")
	http.ListenAndServe(":4000", app.Router)
}
