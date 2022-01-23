package main

import (
	"fmt"
	"net/http"

	"github.com/babuley/puzzles/clockResolver/controllers"
	"github.com/babuley/puzzles/clockResolver/models"
)

func main() {
	clocks := clocks()
	for _, clock := range clocks {
		angle, _ := clock.SolveAngle()
		fmt.Println("Given hour and minute hands, calculate angle", clock, angle)
	}

	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}

func clocks() []models.Clock {
	slice := []models.Clock{}
	h := 0
	for h < 12 {
		m := 0
		for m < 60 {
			slice = append(slice, models.Clock{Hour: h, Minute: m})
			m++
		}
		h++
	}
	return slice
}
