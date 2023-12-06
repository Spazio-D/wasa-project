package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// LOGIN AND REGISTER
	rt.router.POST("/session", rt.wrap(rt.doLogin))

	// SET USERNAME
	rt.router.PUT("/users/:user_id/username", rt.wrap(rt.setMyUserName))

	// CREATE A POST
	rt.router.POST("/users/:user_id/posts", rt.wrap(rt.uploadPhoto))


	
	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}

//CIAO AMO TVB SEI  BELLISSIMO
