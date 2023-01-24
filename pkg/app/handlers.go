package app

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thimovez/cinema/pkg/api"
)

func (s *Server) CreateUser(c *gin.Context) {
	// Set header to json type
	c.Header("Content-Type", "application/json")

	// Declare a new User struct.
	var user api.User

	// Decode the request body into the struct.
	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Printf("API error: %v", err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	// Transfer our user in CreateUser

	response := map[string]string{
		"userName": user.UserName,
	}

	// Encode and return new json response
	c.JSON(http.StatusOK, response)
}
