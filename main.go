package main

import (
	"fmt"
	"net/http"
	_ "rest/docs"

	"rest/rest"
)

// @title REST API for recipes backend.
// @version 1.0
// @description REST API.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:4000
// @BasePath /
func main() {
	app := rest.New()
	fmt.Println("starting server")
	http.ListenAndServe(":4000", app.Router)
}
