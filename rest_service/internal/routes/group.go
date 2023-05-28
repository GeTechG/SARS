package routes

import (
	"git.it-college.ru/i21s617/SARS/rest_service/internal/clients"
	"git.it-college.ru/i21s617/SARS/service_utilities/pkg/proto/ldap_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMembers(c *gin.Context) {
	response, err := clients.GetGroupServiceClient().GetGroupMembers(c, &ldap_service.GetGroupMembersRequest{
		GroupName: c.Param("name"),
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	members := response.GetMembers()

	usersResponse, err := clients.GetUserService().GetUsers(c, &ldap_service.GetUsersRequest{Uids: members})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, usersResponse.Users)
}
