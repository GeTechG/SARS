package routes

import (
	"git.it-college.ru/i21s617/SARS/auth_service/internal/grpc"
	"git.it-college.ru/i21s617/SARS/auth_service/internal/proto/ldap"
	"git.it-college.ru/i21s617/SARS/auth_service/internal/sessions"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
	"net/http"
)

func Auth(c *gin.Context) {
	userRequest := &ldap.AuthUserRequest{}
	err := c.ShouldBindJSON(userRequest)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
	}
	auth, err := grpc.GetUserService().Auth(c, userRequest)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
	}

	m := protojson.MarshalOptions{EmitUnpopulated: true}
	resp, err := m.Marshal(auth)

	if auth.GetValid() {
		sessions.GetSessions().Put(c.Request.Context(), "uid", auth.GetUser().Uid)
		sessions.GetSessions().Put(c.Request.Context(), "userType", auth.GetUser().GetUserType())
		c.Data(http.StatusOK, "application/json", resp)
		return
	}

	c.Data(http.StatusNotAcceptable, "application/json", resp)
}

func IsAuth(c *gin.Context) {
	uid := sessions.GetSessions().Get(c.Request.Context(), "uid").(string)
	userType := sessions.GetSessions().Get(c.Request.Context(), "userType").(string)

	c.JSON(http.StatusOK, map[string]string{
		"uid":      uid,
		"userType": userType,
	})
}
