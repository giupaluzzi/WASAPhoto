package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {

	//User
	rt.router.GET("/users/:id", rt.getUserProfile)
	rt.router.POST("/users/", rt.createUser)

	//Follow
	rt.router.PUT("users/:id/following/:id", rt.followUser)

	//Stream
	rt.router.GET("/users/:id/stream", rt.getMyStream)

	//Username
	rt.router.PUT("/users/:id/username", rt.setUsername)

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
