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

	// FOLLOW A USER
	rt.router.PUT("/users/:follower_id/follows/followed_id", rt.wrap(rt.followUser))

	// UNFOLLOW A USER
	rt.router.DELETE("/users/:follower_id/follows/followed_id", rt.wrap(rt.unfollowUser))

	// BAN A USER
	rt.router.PUT("/users/:user_id/banned/:target_user_id", rt.wrap(rt.banUser))

	// UNBAN A USER
	rt.router.DELETE("/users/:user_id/banned/:target_user_id", rt.wrap(rt.unbanUser))

	// GET USER PROFILE
	rt.router.GET("/users/:user_id/", rt.wrap(rt.getUserProfile))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}

//CIAO AMO TVB SEI  BELLISSIMO
