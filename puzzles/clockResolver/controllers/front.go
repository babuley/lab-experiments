package controllers

import "net/http"

func RegisterControllers() {
	clockCtrl := newClockController()

	http.Handle("/clock", *clockCtrl)

	http.Handle("/clock/", *clockCtrl)
}
