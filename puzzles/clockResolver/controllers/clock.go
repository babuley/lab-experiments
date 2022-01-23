package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/babuley/puzzles/clockResolver/models"
)

type clockControler struct {
	clockPattern *regexp.Regexp
}

func (clockCtrl clockControler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Clock Puzzle"))
}

func newClockController() *clockControler {
	return &clockControler{
		clockPattern: regexp.MustCompile(`^/clock/(\d+)/(\d+)/?`),
	}
}

func (clockControler *clockControler) solve(hour int, minute int, w http.ResponseWriter) {
	c := models.Clock{
		Hour:   hour,
		Minute: minute,
	}
	angle, _ := c.SolveAngle()
	fmt.Println(angle)
}

func (clockCtrl *clockControler) parseRequest(r *http.Request) (models.Clock, error) {
	dec := json.NewDecoder(r.Body)
	var c models.Clock
	err := dec.Decode(&c)
	if err != nil {
		return models.Clock{}, err
	}
	return c, nil
}
