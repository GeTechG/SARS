package routes

import (
	"git.it-college.ru/i21s617/SARS/auth_service/internal/grpc_service_client"
	"git.it-college.ru/i21s617/SARS/service_utilities/pkg/proto/ldap_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUser(c *gin.Context) {
	user, err := grpc_service_client.GetUserService().GetUser(c, &ldap_service.GetUserRequest{
		Uid: c.Param("uid"),
	})
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, user.User)
}
