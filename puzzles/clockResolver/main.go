package main

import (
	"net/http"

	"github.com/babuley/puzzles/clockResolver/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
