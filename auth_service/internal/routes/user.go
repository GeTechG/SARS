package routes

import (
	"git.it-college.ru/i21s617/SARS/auth_service/internal/grpc"
	"git.it-college.ru/i21s617/SARS/auth_service/internal/proto/ldap"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUser(c *gin.Context) {
	user, err := grpc.GetUserService().GetUser(c, &ldap.GetUserRequest{
		Uid: c.Param("uid"),
	})
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, user.User)
}
