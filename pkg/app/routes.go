package app

import (
	"github.com/gin-gonic/gin"
)

func (s *Server) Routes() *gin.Engine {
	r := s.router

	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	r.Use(gin.Logger())
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	r.POST("/user", s.CreateUser)

	return r
}
