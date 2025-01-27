package v1

import (
	"net/http"

	authv1 "github.com/Homyakadze14/ApiGatewateForOrbitOfSuccess/proto/gen/auth"
	"github.com/gin-gonic/gin"
)

type authRoutes struct {
	s authv1.AuthClient
}

func NewAuthRoutes(handler *gin.RouterGroup, s authv1.AuthClient) {
	r := &authRoutes{s}

	g := handler.Group("")
	{
		g.POST("/register", r.register)
	}
}

// @Summary     Register
// @Description Register
// @ID          Auth
// @Tags  	    register
// @Produce     json
// @Success     200
// @Router      /register [post]
func (r *authRoutes) register(c *gin.Context) {
	//resp, _ := r.s.Register(c.Request.Context())
	c.JSON(http.StatusOK, r.s)
}
