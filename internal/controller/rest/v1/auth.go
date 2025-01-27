package v1

import (
	"log/slog"
	"net/http"

	"github.com/Homyakadze14/ApiGatewateForOrbitOfSuccess/internal/common"
	"github.com/Homyakadze14/ApiGatewateForOrbitOfSuccess/internal/entities"
	authv1 "github.com/Homyakadze14/ApiGatewateForOrbitOfSuccess/proto/gen/auth"
	"github.com/gin-gonic/gin"
)

type authRoutes struct {
	s   authv1.AuthClient
	log *slog.Logger
}

func NewAuthRoutes(log *slog.Logger, handler *gin.RouterGroup, s authv1.AuthClient) {
	r := &authRoutes{
		log: log,
		s:   s}

	g := handler.Group("/auth")
	{
		g.POST("/register", r.register)
		g.POST("/login", r.login)
	}
}

// @Summary     Register
// @Description Register
// @ID          Register
// @Tags  	    Auth
// @Accept      json
// @Param 		register body entities.RegisterRequest false "register"
// @Produce     json
// @Success     200 {object} authv1.RegisterResponse
// @Failure     400
// @Failure     404
// @Failure     500
// @Failure     503
// @Router      /auth/register [post]
func (r *authRoutes) register(c *gin.Context) {
	const op = "authRoutes.register"

	log := r.log.With(
		slog.String("op", op),
	)

	var req *entities.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": common.GetErrMessages(err).Error()})
		return
	}

	resp, err := r.s.Register(c.Request.Context(), req.ToGRPC())
	if err != nil {
		code, err := common.GetProtoErrWithStatusCode(err)
		log.Error(err.Error())
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Summary     Login
// @Description Login
// @ID          Login
// @Tags  	    Auth
// @Accept      json
// @Param 		login body entities.LoginRequest false "login"
// @Produce     json
// @Success     200 {object} authv1.LoginResponse
// @Failure     400
// @Failure     401
// @Failure     404
// @Failure     500
// @Failure     503
// @Router      /auth/login [post]
func (r *authRoutes) login(c *gin.Context) {
	const op = "authRoutes.login"

	log := r.log.With(
		slog.String("op", op),
	)

	var req *entities.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": common.GetErrMessages(err).Error()})
		return
	}

	resp, err := r.s.Login(c.Request.Context(), req.ToGRPC())
	if err != nil {
		code, err := common.GetProtoErrWithStatusCode(err)
		log.Error(err.Error())
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
