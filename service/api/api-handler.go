package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// LOGIN AND REGISTER
	rt.router.POST("/session", rt.getHelloWorld)

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
//CIAO AMO TVB SEI  BELLISSIMO