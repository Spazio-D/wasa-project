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

	// GET USER STREAM
	rt.router.GET("/users/:user_id/stream", rt.wrap(rt.getMyStream))

	// LIKE A POST
	rt.router.PUT("/users/:liked_id/posts/:post_id/likes/:user_id", rt.wrap(rt.likePhoto))

	// UNLIKE A POST
	rt.router.DELETE("/users/:liked_id/posts/:post_id/likes/:user_id", rt.wrap(rt.unlikePhoto))

	// COMMENT A POST
	rt.router.POST("/users/:commented_id/posts/:post_id/comments", rt.wrap(rt.commentPhoto))

	// UNCOMMENT A POST
	rt.router.DELETE("/users/:commented_id/posts/:post_id/comments/:comment_id", rt.wrap(rt.uncommentPhoto))

	// DELETE A POST
	rt.router.DELETE("/users/:user_id/posts/:post_id", rt.wrap(rt.deletePhoto))

	// GET POSTS
	rt.router.GET("/users/:user_id/posts", rt.wrap(rt.getPosts))

	// SEARCH USERS
	rt.router.GET("/users", rt.wrap(rt.searchUsers))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}

//CIAO AMO TVB SEI  BELLISSIMO
