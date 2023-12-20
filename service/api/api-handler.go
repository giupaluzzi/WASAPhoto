package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {

	//Login
	rt.router.POST("/session", rt.doLogin)

	//User
	rt.router.GET("/users/:userid", rt.getUserProfile)

	//Username
	rt.router.PUT("/users/:userid/username", rt.setUsername)

	//Stream
	rt.router.GET("/users/:userid/stream", rt.getMyStream)

	//Follow
	rt.router.PUT("/users/:userid/following/:followinguid", rt.followUser)
	rt.router.DELETE("/users/:userid/following/:followinguid", rt.unfollowUser)

	//Photo
	rt.router.POST("/users/:userid/photos/", rt.uploadPhoto)
	rt.router.DELETE("/users/:userid/photos/:photoid", rt.deletePhoto)

	//Likes
	rt.router.PUT("/users/:userid/photos/:photoid/likes/:likeid", rt.likePhoto)
	rt.router.DELETE("/users/:userid/photos/:photoid/likes/:likeid", rt.unlikePhoto)

	//Ban
	rt.router.PUT("/users/:userid/banned/:banneduid", rt.banUser)
	rt.router.DELETE("/users/:userid/banned/:banneduid", rt.unbanUser)

	//Comment
	rt.router.POST("/users/:userid/photos/:photoid/comments/", rt.commentPhoto)
	rt.router.DELETE("/users/:userid/photos/:photoid/comments/:commentid", rt.deleteComment)

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
