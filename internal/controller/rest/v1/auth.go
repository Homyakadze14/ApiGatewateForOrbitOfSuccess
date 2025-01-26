package v1

import (
	"github.com/Homyakadze14/ApiGatewateForOrbitOfSuccess/internal/services"

	"github.com/gin-gonic/gin"
)

type authRoutes struct {
	s *services.AuthService
}

func NewAuthRoutes(handler *gin.RouterGroup, s *services.AuthService) {
	r := &authRoutes{s}

	_ = r
	g := handler.Group("")
	{
	}
	_ = g
}
