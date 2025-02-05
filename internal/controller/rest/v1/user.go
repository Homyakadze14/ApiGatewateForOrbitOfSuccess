package v1

import (
	"log/slog"
	"net/http"

	"github.com/Homyakadze14/ApiGatewateForOrbitOfSuccess/internal/common"
	"github.com/Homyakadze14/ApiGatewateForOrbitOfSuccess/internal/entities"
	userv1 "github.com/Homyakadze14/ApiGatewateForOrbitOfSuccess/proto/gen/user"
	"github.com/gin-gonic/gin"
)

type userRoutes struct {
	s   userv1.UserClient
	log *slog.Logger
}

func NewUserRoutes(log *slog.Logger, handler *gin.RouterGroup, s userv1.UserClient) {
	r := &userRoutes{
		log: log,
		s:   s}

	g := handler.Group("/user")
	{
		g.PUT("", r.update)
	}
}

// @Summary     Update user info
// @Description Update user info
// @ID          Update user info
// @Tags  	    User
// @Accept      json
// @Param 		info body entities.UserUpdateRequest false "info"
// @Produce     json
// @Success     200 {object} userv1.UpdateInfoResponse
// @Failure     400
// @Failure     404
// @Failure     500
// @Failure     503
// @Router      /user [put]
func (r *userRoutes) update(c *gin.Context) {
	const op = "userRoutes.update"

	log := r.log.With(
		slog.String("op", op),
	)

	var req *entities.UserUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": common.GetErrMessages(err).Error()})
		return
	}

	resp, err := r.s.UpdateInfo(c.Request.Context(), req.ToGRPC())
	if err != nil {
		code, err := common.GetProtoErrWithStatusCode(err)
		log.Error(err.Error())
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
