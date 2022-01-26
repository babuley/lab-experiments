package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/babuley/puzzles/clockResolver/models"
)

type clockControler struct {
	clockPattern *regexp.Regexp
}

func (clockCtrl clockControler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/clocks" {
		if r.Method == http.MethodGet {
			clockCtrl.getAll(w, r)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		if r.Method == http.MethodGet {
			clockCtrl.getClock(w, r)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

func (clockCtrl *clockControler) getAll(w http.ResponseWriter, r *http.Request) {
	clocks := models.GetAllClocks()
	encodeResponseAsJSON(clocks, w)
}

func (clockCtrl *clockControler) getClock(w http.ResponseWriter, r *http.Request) {
	matches := clockCtrl.clockPattern.FindStringSubmatch(r.URL.Path)
	if len(matches) == 0 {
		w.WriteHeader(http.StatusBadRequest)
	}
	hour, err := strconv.Atoi(matches[1])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	minute, err := strconv.Atoi(matches[2])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	clock := models.NewClock(hour, minute)
	encodeResponseAsJSON(clock, w)
}

func newClockController() *clockControler {
	return &clockControler{
		clockPattern: regexp.MustCompile(`/clock/(\d+)/(\d+)`),
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
