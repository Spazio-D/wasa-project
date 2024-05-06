package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {

	// LOGIN AND REGISTER
	rt.router.POST("/session", rt.wrap(rt.doLogin, false))

	// SET USERNAME
	rt.router.PUT("/users/:user_id/username", rt.wrap(rt.setMyUserName, true))

	// CREATE A POST
	rt.router.POST("/users/:user_id/posts", rt.wrap(rt.uploadPhoto, true))

	// GET USER POSTS
	rt.router.GET("/users/:user_id/posts", rt.wrap(rt.getPosts, true))

	// FOLLOW A USER
	rt.router.PUT("/users/:user_id/follows/:followed_id", rt.wrap(rt.followUser, true))

	// UNFOLLOW A USER
	rt.router.DELETE("/users/:user_id/follows/:followed_id", rt.wrap(rt.unfollowUser, true))

	// BAN A USER
	rt.router.PUT("/users/:user_id/banned/:target_user_id", rt.wrap(rt.banUser, true))

	// UNBAN A USER
	rt.router.DELETE("/users/:user_id/banned/:target_user_id", rt.wrap(rt.unbanUser, true))

	// GET USER PROFILE
	rt.router.GET("/users/:user_id", rt.wrap(rt.getUserProfile, true))

	// GET USER STREAM
	rt.router.GET("/users/:user_id/stream", rt.wrap(rt.getMyStream, true))

	// LIKE A POST
	rt.router.PUT("/users/:user_id/posts/:post_id/likes/:liker_id", rt.wrap(rt.likePhoto, true))

	// UNLIKE A POST
	rt.router.DELETE("/users/:user_id/posts/:post_id/likes/:liker_id", rt.wrap(rt.unlikePhoto, true))

	// COMMENT A POST
	rt.router.POST("/users/:user_id/posts/:post_id/comments", rt.wrap(rt.commentPhoto, true))

	// UNCOMMENT A POST
	rt.router.DELETE("/users/:user_id/posts/:post_id/comments/:comment_id", rt.wrap(rt.uncommentPhoto, true))

	// DELETE A POST
	rt.router.DELETE("/users/:user_id/posts/:post_id", rt.wrap(rt.deletePhoto, true))

	// SEARCH USERS
	rt.router.GET("/users", rt.wrap(rt.searchUsers, true))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
