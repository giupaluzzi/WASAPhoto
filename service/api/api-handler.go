package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {

	// Login
	rt.router.POST("/session", rt.wrap(rt.doLogin))

	// User
	rt.router.GET("/users/:userid", rt.wrap(rt.getUserProfile))

	// Username
	rt.router.PUT("/users/:userid/username", rt.wrap(rt.setUsername))

	// Stream
	rt.router.GET("/users/:userid/stream", rt.wrap(rt.getMyStream))

	// Follow
	rt.router.PUT("/users/:userid/following/:followinguid", rt.wrap(rt.followUser))
	rt.router.DELETE("/users/:userid/following/:followinguid", rt.wrap(rt.unfollowUser))

	// Photo
	rt.router.POST("/users/:userid/photos/", rt.wrap(rt.uploadPhoto))
	rt.router.DELETE("/users/:userid/photos/:photoid", rt.wrap(rt.deletePhoto))

	// Likes
	rt.router.PUT("/users/:userid/photos/:photoid/likes/:likeid", rt.wrap(rt.likePhoto))
	rt.router.DELETE("/users/:userid/photos/:photoid/likes/:likeid", rt.wrap(rt.unlikePhoto))

	// Ban
	rt.router.PUT("/users/:userid/banned/:banneduid", rt.wrap(rt.banUser))
	rt.router.DELETE("/users/:userid/banned/:banneduid", rt.wrap(rt.unbanUser))

	// Comment
	rt.router.POST("/users/:userid/photos/:photoid/comments/", rt.wrap(rt.commentPhoto))
	rt.router.DELETE("/users/:userid/photos/:photoid/comments/:commentid", rt.wrap(rt.deleteComment))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
