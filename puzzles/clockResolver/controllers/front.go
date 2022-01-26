package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func RegisterControllers() {
	clockCtrl := newClockController()

	http.Handle("/clock/", *clockCtrl)
	http.Handle("/clocks", *clockCtrl)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}
