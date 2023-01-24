package app

import (
	"github.com/gin-gonic/gin"
	"github.com/thimovez/cinema/pkg/util"
)

type Server struct {
	router *gin.Engine
	config *util.Config
}

func NewServer(router *gin.Engine, config *util.Config) (*Server, error) {
	return &Server{
		router: router,
		config: config,
	}, nil
}

func (s *Server) Run() error {
	r := s.Routes()

	err := r.Run(s.config.Host)
	if err != nil {
		return err
	}
	return nil
}
