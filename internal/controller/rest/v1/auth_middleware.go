package v1

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/Homyakadze14/ApiGatewateForOrbitOfSuccess/internal/common"
	authv1 "github.com/Homyakadze14/ApiGatewateForOrbitOfSuccess/proto/gen/auth"
	"github.com/gin-gonic/gin"
)

func authMiddleware(log *slog.Logger, s authv1.AuthClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		_, err := s.Verify(ctx, &authv1.VerifyRequest{AccessToken: token})
		if err != nil {
			_, err := common.GetProtoErrWithStatusCode(err)
			log.Error(err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.Next()
	}
}
