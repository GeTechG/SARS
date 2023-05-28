package routes

import (
	"context"
	"git.it-college.ru/i21s617/SARS/rest_service/internal/clients"
	"git.it-college.ru/i21s617/SARS/service_utilities/pkg/proto/class_schedule_service"
	"git.it-college.ru/i21s617/SARS/service_utilities/pkg/proto/ldap_service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"google.golang.org/protobuf/encoding/protojson"
	"net/http"
	"strconv"
)

func GetAttendances(c *gin.Context) {
	classId, err := strconv.ParseInt(c.Param("class_id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := clients.GetAttendanceServiceClient().GetAttendance(c, &class_schedule_service.GetAttendanceRequest{ClassId: classId})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	m := protojson.MarshalOptions{EmitUnpopulated: true}
	resp, err := m.Marshal(response)

	c.Data(http.StatusOK, "application/json", resp)
}

func SetAttendancesValidation(ctx context.Context, sl validator.StructLevel) {
	attendance, ok := sl.Current().Interface().(class_schedule_service.Attendance)
	if ok {
		class, err := clients.GetClassScheduleClient().GetClass(ctx, &class_schedule_service.GetClassRequest{ClassId: attendance.GetClassId()})
		if err != nil {
			sl.ReportError(attendance.ClassId, "ClassId", "classId", "Invalid class_id", "")
			return
		}

		membersResponse, err := clients.GetGroupServiceClient().GetGroupMembers(ctx, &ldap_service.GetGroupMembersRequest{GroupName: class.GetGroup()})
		if err != nil {
			sl.ReportError(class.Group, "GroupId", "groupId", "Invalid group_id", "")
			return
		}

		var userIdContain = false
		for _, member := range membersResponse.Members {
			if member == attendance.GetUserUid() {
				userIdContain = true
			}
		}

		if !userIdContain {
			sl.ReportError(attendance.UserUid, "UserUid", "userUid", "Invalid user_uid", "")
		}
	}
}

func SetAttendances(c *gin.Context) {
	var attendances []*class_schedule_service.Attendance
	err := c.ShouldBindJSON(&attendances)
	if err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			var fieldErrors = make([]string, 0, len(errors))
			for _, fieldError := range errors {
				fieldErrors = append(fieldErrors, fieldError.Error())
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors:": fieldErrors})
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
		}
		return
	}

	response, err := clients.GetAttendanceServiceClient().SetAttendance(c, &class_schedule_service.SetAttendanceRequest{Attendances: attendances})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(response.Errors) > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	c.JSON(http.StatusOK, response)
}
